package postgres

import (
	"errors"
	"fmt"
	"lms/lms_book_service/protogen/book_service"
	"time"
)

// AddCategory ...
func (stg Postgres) AddCategory(id string, entity *book_service.CreateCategoryRequest) error {

	_, err := stg.db.Exec(`INSERT INTO "category" (
		"id",
		"title"
		) VALUES(
		$1,
		$2
	)`,
		id,
		entity.Title,
	)
	if err != nil {
		return err
	}

	return nil
}

// GetCategoryByID ...
func (stg Postgres) GetCategoryByID(id string) (*book_service.Category, error) {
	result := &book_service.Category{}

	var updatedAt *time.Time
	err := stg.db.QueryRow(`SELECT
		"id",
		"title",
		"status",
		"created_at",
		"updated_at"
	FROM "category" WHERE id=$1`, id).Scan(
		&result.Id,
		&result.Title,
		&result.Status,
		&result.CreatedAt,
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

// GetCategoryList ...
func (stg Postgres) GetCategoryList(offset, limit int, search string) (resp *book_service.GetCategoryListResponse, err error) {

	resp = &book_service.GetCategoryListResponse{
		Categories: make([]*book_service.Category, 0),
	}

	rows, err := stg.db.Queryx(`
	Select 
		"id",
		"title",
		"status",
		"created_at",
		"updated_at"
 		from "category" WHERE 
 		("title" ILIKE '%' || $1 || '%')
 		LIMIT $2 
		OFFSET $3`, search, limit, offset)

	if err != nil {
		return resp, err
	}

	for rows.Next() {
		a := &book_service.Category{}
		var updatedAt *string

		err = rows.Scan(
			&a.Id,
			&a.Title,
			&a.Status,
			&a.CreatedAt,
			&updatedAt,
		)
		fmt.Println(a)
		if updatedAt != nil {
			a.UpdatedAt = *updatedAt
		}

		if err != nil {
			return resp, err
		}

		resp.Categories = append(resp.Categories, a)

	}

	return resp, nil
}

// DeleteCategory ...
func (stg Postgres) DeleteCategory(idStr string) error {
	rows, err := stg.db.Exec(`UPDATE "category" SET "updated_at"=now(),"status"='disabled' Where id=$1 and "status"='enabled'`, idStr)

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

	return errors.New("Cannot delete Category becouse 'status'='disabled'")
}

// EnabledCategory ...
func (stg Postgres) EnabledCategory(idStr string) error {
	rows, err := stg.db.Exec(`UPDATE "category" SET "updated_at"=now(),"status"='enabled' Where id=$1 and "status"='disabled'`, idStr)

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

	return errors.New("Cannot enabled Category becouse 'status'='enabled'")
}

// UpdateCategory ...
func (stg Postgres) UpdateCategory(category *book_service.UpdateCategoryRequest) error {

	rows, err := stg.db.NamedExec(`Update "category" set "title"=:t, "updated_at"=now() Where "id"=:id and "status"='enabled'`, map[string]interface{}{
		"id": category.Id,
		"t":  category.Title,
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

	return errors.New("Cannot update Category becouse 'status'='disabled'")
}
