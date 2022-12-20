package author

import (
	"context"
	"fmt"
	"lms/lms_book_service/protogen/book_service"
	"lms/lms_book_service/storage"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authorService struct {
	stg storage.Interfaces
	book_service.UnimplementedAuthorServiceServer
}

func NewAuthorService(stg storage.Interfaces) *authorService {
	return &authorService{
		stg: stg,
	}
}

func (s *authorService) CreateAuthor(ctx context.Context, req *book_service.CreateAuthorRequest) (*book_service.Author, error) {
	fmt.Println("<<< ---- CreateAuthor ---->>>")

	id := uuid.New()

	err := s.stg.AddAuthor(id.String(), req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.AddAuthor: %s", err)
	}

	author, err := s.stg.GetAuthorByID(id.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetAuthorByID: %s", err)
	}

	return &book_service.Author{
		Id:        author.Id,
		Name:      author.Name,
		Status:    author.Status,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
	}, nil
}

func (s *authorService) DeleteAuthor(ctx context.Context, req *book_service.DeleteAuthorRequest) (*book_service.DeleteAuthorResponse, error) {
	fmt.Println("<<< ---- DeleteAuthor ---->>>")

	author, err := s.stg.GetAuthorByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetAuthorByID: %s", err)
	}

	err = s.stg.DeleteAuthor(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteAuthor: %s", err)
	}

	return &book_service.DeleteAuthorResponse{
		Status: author.Status,
	}, nil
}

func (s *authorService) EnabledAuthor(ctx context.Context, req *book_service.EnabledAuthorRequest) (*book_service.EnabledAuthorResponse, error) {
	fmt.Println("<<< ---- EnabledAuthor ---->>>")

	author, err := s.stg.GetAuthorByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetAuthorByID: %s", err)
	}

	err = s.stg.EnabledAuthor(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.EnabledAuthor: %s", err)
	}

	return &book_service.EnabledAuthorResponse{
		Status: author.Status,
	}, nil
}

func (s *authorService) GetAuthorList(ctx context.Context, req *book_service.GetAuthorListRequest) (*book_service.GetAuthorListResponse, error) {
	fmt.Println("<<< ---- GetAuthorList ---->>>")

	res, err := s.stg.GetAuthorList(int(req.Offset), int(req.Limit), req.Search)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetAuthorList: %s", err)
	}

	return res, nil
}

func (s *authorService) GetAuthorById(ctx context.Context, req *book_service.GetAuthorByIDRequest) (*book_service.Author, error) {
	fmt.Println("<<< ---- GetAuthorById ---->>>")

	author, err := s.stg.GetAuthorByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetAuthorByID: %s", err)
	}

	return &book_service.Author{
		Id:        author.Id,
		Name:      author.Name,
		Status:    author.Status,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
	}, nil
}
