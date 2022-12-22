package postgres

import (
	"errors"
	"lms/lms_book_service/protogen/book_service"
	"time"
)

// AddLocation ...
func (stg Postgres) AddLocation(id string, entity *book_service.CreateLocationRequest) error {

	_, err := stg.db.Exec(`INSERT INTO "location" (
		"id",
		"name"
		) VALUES(
		$1,
		$2
	)`,
		id,
		entity.Name,
	)

	if err != nil {
		return err
	}

	return nil
}

// GetLocationByID ...
func (stg Postgres) GetLocationByID(id string) (*book_service.Location, error) {
	result := &book_service.Location{}

	var updatedAt *time.Time
	err := stg.db.QueryRow(`SELECT
		"id",
		"name",
		"status",
		"created_at",
		"updated_at"
	FROM "location" WHERE id=$1`, id).Scan(
		&result.Id,
		&result.Name,
		&result.Status,
		&result.CreatedAt, // !
		&updatedAt,
	)

	if err != nil {
		return result, err
	}

	if updatedAt != nil {
		result.UpdatedAt = updatedAt.String()
	}

	return result, nil
}

// GetLocationList ...
func (stg Postgres) GetLocationList(offset, limit int, search string) (resp *book_service.GetLocationListResponse, err error) {

	resp = &book_service.GetLocationListResponse{
		Locations: make([]*book_service.Location, 0),
	}

	rows, err := stg.db.Queryx(`
	Select 
		"id",
		"name",
		"status",
		"created_at",
		"updated_at"
 		from "location" WHERE 
 		("name" ILIKE '%' || $1 || '%')
 		LIMIT $2 
		OFFSET $3`, search, limit, offset)

	if err != nil {
		return resp, err
	}

	for rows.Next() {
		a := &book_service.Location{}
		var updatedAt *string

		err = rows.Scan(
			&a.Id,
			&a.Name,
			&a.Status,
			&a.CreatedAt,
			&updatedAt,
		)

		if updatedAt != nil {
			a.UpdatedAt = *updatedAt
		}

		if err != nil {
			return resp, err
		}

		resp.Locations = append(resp.Locations, a)

	}

	return resp, nil
}

// DeleteLocation ...
func (stg Postgres) DeleteLocation(idStr string) error {
	rows, err := stg.db.Exec(`UPDATE "location" SET "updated_at"=now(),"status"='disabled' Where id=$1 and "status"='enabled'`, idStr)

	if err != nil {
		return err
	}

	n, err := rows.RowsAffected()

	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("cannot delete Location becouse 'status'='disabled'")
}

// EnabledLocation ...
func (stg Postgres) EnabledLocation(idStr string) error {
	rows, err := stg.db.Exec(`UPDATE "location" SET "updated_at"=now(),"status"='enabled' Where id=$1 and "status"='disabled'`, idStr)

	if err != nil {
		return err
	}

	n, err := rows.RowsAffected()

	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("cannot enabled Location becouse 'status'='enabled'")
}

// UpdateLocation ...
func (stg Postgres) UpdateLocation(location *book_service.UpdateLocationRequest) error {

	rows, err := stg.db.NamedExec(`Update "location" set "name"=:n, "updated_at"=now() Where "id"=:id and "status"='enabled'`, map[string]interface{}{
		"id": location.Id,
		"n":  location.Name,
	})

	if err != nil {
		return err
	}

	n, err := rows.RowsAffected()

	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("cannot update Location becouse 'status'='disabled'")
}
