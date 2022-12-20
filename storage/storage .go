package storage

import "lms/lms_book_service/protogen/book_service"

// Interfaces ...
type Interfaces interface {
	// For Author service
	AddAuthor(id string, entity *book_service.CreateAuthorRequest) error
	GetAuthorByID(id string) (*book_service.Author, error)
	GetAuthorList(offset, limit int, search string) (resp *book_service.GetAuthorListResponse, err error)
	DeleteAuthor(idStr string) error

	// For Book service
	AddBook(id string, entity *book_service.CreateBookRequest) error
	GetBookByID(id string) (*book_service.Book, error)
	GetBookList(offset, limit int, search string) (resp *book_service.GetBookListResponse, err error)
	DeleteBook(idStr string) error
	UpdateBook(book *book_service.UpdateBookRequest) error

	// For Category service
	AddCategory(id string, entity *book_service.CreateCategoryRequest) error
	GetCategoryByID(id string) (*book_service.Category, error)
	GetCategoryList(offset, limit int, search string) (resp *book_service.GetCategoryListResponse, err error)
	DeleteCategory(idStr string) error
	UpdateCategory(category *book_service.UpdateCategoryRequest) error

	// For Location service
	AddLocation(id string, entity *book_service.CreateLocationRequest) error
	GetLocationByID(id string) (*book_service.Location, error)
	GetLocationList(offset, limit int, search string) (resp *book_service.GetLocationListResponse, err error)
	DeleteLocation(idStr string) error
	UpdateLocation(location *book_service.UpdateLocationRequest) error
}
