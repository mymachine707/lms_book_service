package postgres

import (
	"errors"
	"lms/lms_book_service/protogen/book_service"
	"time"
)

var err error

// AddAuthor ...
func (stg Postgres) AddAuthor(id string, entity *book_service.CreateAuthorRequest) error {

	_, err = stg.db.Exec(`INSERT INTO "author" (
		"id",
		"name",
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

// GetAuthorByID ...
func (stg Postgres) GetAuthorByID(id string) (*book_service.Author, error) {
	result := &book_service.Author{}

	var updatedAt *time.Time
	err := stg.db.QueryRow(`SELECT
		"id",
		"name",
		"status",
		"created_at",
		"updated_at"
	FROM "author" WHERE id=$1`, id).Scan(
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

// GetAuthorList ...
func (stg Postgres) GetAuthorList(offset, limit int, search string) (resp *book_service.GetAuthorListResponse, err error) {

	resp = &book_service.GetAuthorListResponse{
		Authors: make([]*book_service.Author, 0),
	}

	rows, err := stg.db.Queryx(`
	Select 
		"id",
		"name",
		"status",
		"created_at",
		"updated_at"
 		from "author" WHERE 
 		(("name" ILIKE '%' || $1 || '%')
 		LIMIT $2 
		OFFSET $3`, search, limit, offset)

	if err != nil {
		return resp, err
	}

	for rows.Next() {
		a := &book_service.Author{}
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

		resp.Authors = append(resp.Authors, a)

	}

	return resp, nil
}

// DeleteAuthor ...
func (stg Postgres) DeleteAuthor(idStr string) error {
	rows, err := stg.db.Exec(`UPDATE "author" SET "updated_at"=now(),"status"='disabled' Where id=$1 and "status"='enabled'`, idStr)

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

	return errors.New("Cannot delete Author becouse 'status'='disabled'")
}

// EnabledAuthor ...
func (stg Postgres) EnabledAuthor(idStr string) error {
	rows, err := stg.db.Exec(`UPDATE "author" SET "updated_at"=now(),"status"='enabled' Where id=$1 and "status"='disabled'`, idStr)

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

	return errors.New("Cannot enabled Author becouse 'status'='enabled'")
}
