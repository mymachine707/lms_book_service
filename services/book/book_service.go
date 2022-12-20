package book

import (
	"context"
	"fmt"
	"lms/lms_book_service/protogen/book_service"
	"lms/lms_book_service/storage"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type bookService struct {
	stg storage.Interfaces
	book_service.UnimplementedBookServiceServer
}

func NewBookService(stg storage.Interfaces) *bookService {
	return &bookService{
		stg: stg,
	}
}

func (s *bookService) CreateBook(ctx context.Context, req *book_service.CreateBookRequest) (*book_service.Book, error) {
	fmt.Println("<<< ---- CreateBook ---->>>")

	id := uuid.New()

	err := s.stg.AddBook(id.String(), req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.AddBook: %s", err)
	}

	book, err := s.stg.GetBookByID(id.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetBookByID: %s", err)
	}

	return &book_service.Book{
		Id:         book.Id,
		Name:       book.Name,
		AuthorId:   book.AuthorId,
		CategoryId: book.CategoryId,
		LocationId: book.LocationId,
		ISBN:       book.ISBN,
		Quantity:   book.Quantity,
		Status:     book.Status,
		CreatedAt:  book.CreatedAt,
		UpdatedAt:  book.UpdatedAt,
	}, nil
}

func (s *bookService) UpdateBook(ctx context.Context, req *book_service.UpdateBookRequest) (*book_service.UpdateBookResponse, error) {
	fmt.Println("<<< ---- UpdateBook ---->>>")

	err := s.stg.UpdateBook(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.UpdateBook: %s", err)
	}

	book, err := s.stg.GetBookByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetBookByID: %s", err)
	}

	return &book_service.UpdateBookResponse{
		Id:         book.Id,
		LocationId: book.LocationId,
		Quantity:   book.Quantity,
		UpdatedAt:  book.UpdatedAt,
	}, nil
}

func (s *bookService) DeleteBook(ctx context.Context, req *book_service.DeleteBookRequest) (*book_service.DeleteBookResponse, error) {
	fmt.Println("<<< ---- DeleteBook ---->>>")

	book, err := s.stg.GetBookByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetBookByID: %s", err)
	}

	err = s.stg.DeleteBook(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteBook: %s", err)
	}

	return &book_service.DeleteBookResponse{
		Status: book.Status,
	}, nil
}

func (s *bookService) EnabledBook(ctx context.Context, req *book_service.EnabledBookRequest) (*book_service.EnabledBookResponse, error) {
	fmt.Println("<<< ---- EnabledBook ---->>>")

	book, err := s.stg.GetBookByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetBookByID: %s", err)
	}

	err = s.stg.EnabledBook(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.EnabledBook: %s", err)
	}

	return &book_service.EnabledBookResponse{
		Status: book.Status,
	}, nil
}

func (s *bookService) GetBookList(ctx context.Context, req *book_service.GetBookListRequest) (*book_service.GetBookListResponse, error) {
	fmt.Println("<<< ---- GetBookList ---->>>")

	res, err := s.stg.GetBookList(int(req.Offset), int(req.Limit), req.Search)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetBookList: %s", err)
	}

	return res, nil
}

func (s *bookService) GetBookById(ctx context.Context, req *book_service.GetBookByIDRequest) (*book_service.Book, error) {
	fmt.Println("<<< ---- GetBookById ---->>>")

	book, err := s.stg.GetBookByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetBookByID: %s", err)
	}

	return &book_service.Book{
		Id:         book.Id,
		Name:       book.Name,
		AuthorId:   book.AuthorId,
		CategoryId: book.CategoryId,
		LocationId: book.LocationId,
		ISBN:       book.ISBN,
		Quantity:   book.Quantity,
		Status:     book.Status,
		CreatedAt:  book.CreatedAt,
		UpdatedAt:  book.UpdatedAt,
	}, nil
}
