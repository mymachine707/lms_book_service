package category

import (
	"context"
	"fmt"
	"lms/lms_book_service/protogen/book_service"
	"lms/lms_book_service/storage"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type categoryService struct {
	stg storage.Interfaces
	book_service.UnimplementedCategoryServiceServer
}

func NewCategoryService(stg storage.Interfaces) *categoryService {
	return &categoryService{
		stg: stg,
	}
}

func (s *categoryService) CreateCategory(ctx context.Context, req *book_service.CreateCategoryRequest) (*book_service.Category, error) {
	fmt.Println("<<< ---- CreateCategory ---->>>")

	id := uuid.New()

	err := s.stg.AddCategory(id.String(), req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.AddCategory: %s", err)
	}

	category, err := s.stg.GetCategoryByID(id.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetCategoryByID: %s", err)
	}

	return &book_service.Category{
		Id:        category.Id,
		Title:     category.Title,
		Status:    category.Status,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}, nil
}

func (s *categoryService) UpdateCategory(ctx context.Context, req *book_service.UpdateCategoryRequest) (*book_service.Category, error) {
	fmt.Println("<<< ---- UpdateCategory ---->>>")

	err := s.stg.UpdateCategory(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.UpdateCategory: %s", err)
	}

	category, err := s.stg.GetCategoryByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetCategoryByID: %s", err)
	}

	return &book_service.Category{
		Id:        category.Id,
		Title:     category.Title,
		Status:    category.Status,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}, nil
}

func (s *categoryService) DeleteCategory(ctx context.Context, req *book_service.DeleteCategoryRequest) (*book_service.DeleteCategoryResponse, error) {
	fmt.Println("<<< ---- DeleteCategory ---->>>")

	category, err := s.stg.GetCategoryByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetCategoryByID: %s", err)
	}

	err = s.stg.DeleteCategory(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteCategory: %s", err)
	}

	return &book_service.DeleteCategoryResponse{
		Status: category.Status,
	}, nil
}

func (s *categoryService) EnabledCategory(ctx context.Context, req *book_service.EnabledCategoryRequest) (*book_service.EnabledCategoryResponse, error) {
	fmt.Println("<<< ---- EnabledCategory ---->>>")

	category, err := s.stg.GetCategoryByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetCategoryByID: %s", err)
	}

	err = s.stg.EnabledCategory(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.EnabledCategory: %s", err)
	}

	return &book_service.EnabledCategoryResponse{
		Status: category.Status,
	}, nil
}

func (s *categoryService) GetCategoryList(ctx context.Context, req *book_service.GetCategoryListRequest) (*book_service.GetCategoryListResponse, error) {
	fmt.Println("<<< ---- GetCategoryList ---->>>")

	res, err := s.stg.GetCategoryList(int(req.Offset), int(req.Limit), req.Search)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetCategoryList: %s", err)
	}

	return res, nil
}

func (s *categoryService) GetCategoryById(ctx context.Context, req *book_service.GetCategoryByIdRequest) (*book_service.Category, error) {
	fmt.Println("<<< ---- GetCategoryById ---->>>")

	category, err := s.stg.GetCategoryByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetCategoryByID: %s", err)
	}

	return &book_service.Category{
		Id:        category.Id,
		Title:     category.Title,
		Status:    category.Status,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}, nil
}
