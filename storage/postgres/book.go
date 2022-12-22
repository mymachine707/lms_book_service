package postgres

import (
	"errors"
	"lms/lms_book_service/protogen/book_service"
	"time"
)

// AddBook ...
func (stg Postgres) AddBook(id string, entity *book_service.CreateBookRequest) error {

	_, err := stg.db.Exec(`INSERT INTO "book" (
		"id",
		"name",
		"author_id" ,
   		"category_id",
   		"location_id",
   		"ISBN",
   		"quantity"
		) VALUES(
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7
	)`,
		id,
		entity.Name,
		entity.AuthorId,
		entity.CategoryId,
		entity.LocationId,
		entity.ISBN,
		entity.Quantity,
	)

	if err != nil {
		return err
	}

	return nil
}

// GetBookByID ...
func (stg Postgres) GetBookByID(id string) (*book_service.Book, error) {
	result := &book_service.Book{}

	var updatedAt *time.Time
	err := stg.db.QueryRow(`SELECT
		"id",
		"name",
		"author_id" ,
   		"category_id",
   		"location_id",
   		"ISBN",
   		"quantity",
		"status",
		"created_at",
		"updated_at"
	FROM "book" WHERE id=$1`, id).Scan(
		&result.Id,
		&result.Name,
		&result.AuthorId,
		&result.CategoryId,
		&result.LocationId,
		&result.ISBN,
		&result.Quantity,
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

// GetBookList ...
func (stg Postgres) GetBookList(offset, limit int, search string) (resp *book_service.GetBookListResponse, err error) {

	resp = &book_service.GetBookListResponse{
		Books: make([]*book_service.Book, 0),
	}

	rows, err := stg.db.Queryx(`
	Select 
		"id",
		"name",
		"author_id" ,
   		"category_id",
   		"location_id",
   		"ISBN",
   		"quantity",
		"status",
		"created_at",
		"updated_at"
 		from "book" WHERE 
 		("name" ILIKE '%' || $1 || '%')
 		LIMIT $2 
		OFFSET $3`, search, limit, offset)

	if err != nil {
		return resp, err
	}

	for rows.Next() {
		a := &book_service.Book{}
		var updatedAt *string

		err = rows.Scan(
			&a.Id,
			&a.Name,
			&a.AuthorId,
			&a.CategoryId,
			&a.LocationId,
			&a.ISBN,
			&a.Quantity,
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

		resp.Books = append(resp.Books, a)

	}

	return resp, nil
}

// DeleteBook ...
func (stg Postgres) DeleteBook(idStr string) error {
	rows, err := stg.db.Exec(`UPDATE "book" SET "updated_at"=now(),"status"='disabled' Where id=$1 and "status"='enabled'`, idStr)

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

	return errors.New("cannot delete Book becouse 'status'='disabled'")
}

// EnabledBook ...
func (stg Postgres) EnabledBook(idStr string) error {
	rows, err := stg.db.Exec(`UPDATE "book" SET "updated_at"=now(),"status"='enabled' Where id=$1 and "status"='disabled'`, idStr)

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

	return errors.New("cannot enabled Book becouse 'status'='enabled'")
}

// UpdateBook ...
func (stg Postgres) UpdateBook(book *book_service.UpdateBookRequest) error {

	rows, err := stg.db.NamedExec(`Update "book" set "quantity"=:q, "location_id"=:l, "updated_at"=now() Where "id"=:id and "status"='enabled'`, map[string]interface{}{
		"id": book.Id,
		"q":  book.Quantity,
		"l":  book.LocationId,
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

	return errors.New("cannot update Book becouse 'status'='disabled'")
}
