package location

import (
	"context"
	"fmt"
	"lms/lms_book_service/protogen/book_service"
	"lms/lms_book_service/storage"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type locationService struct {
	stg storage.Interfaces
	book_service.UnimplementedLocationServiceServer
}

func NewLocationService(stg storage.Interfaces) *locationService {
	return &locationService{
		stg: stg,
	}
}

func (s *locationService) CreateLocation(ctx context.Context, req *book_service.CreateLocationRequest) (*book_service.Location, error) {
	fmt.Println("<<< ---- CreateLocation ---->>>")

	id := uuid.New()

	err := s.stg.AddLocation(id.String(), req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.AddLocation: %s", err)
	}

	location, err := s.stg.GetLocationByID(id.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetLocationByID: %s", err)
	}

	return &book_service.Location{
		Id:        location.Id,
		Name:      location.Name,
		Status:    location.Status,
		CreatedAt: location.CreatedAt,
		UpdatedAt: location.UpdatedAt,
	}, nil
}

func (s *locationService) UpdateLocation(ctx context.Context, req *book_service.UpdateLocationRequest) (*book_service.Location, error) {
	fmt.Println("<<< ---- UpdateLocation ---->>>")

	err := s.stg.UpdateLocation(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.UpdateLocation: %s", err)
	}

	location, err := s.stg.GetLocationByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetLocationByID: %s", err)
	}

	return &book_service.Location{
		Id:        location.Id,
		Name:      location.Name,
		Status:    location.Status,
		CreatedAt: location.CreatedAt,
		UpdatedAt: location.UpdatedAt,
	}, nil
}

func (s *locationService) DeleteLocation(ctx context.Context, req *book_service.DeleteLocationRequest) (*book_service.DeleteLocationResponse, error) {
	fmt.Println("<<< ---- DeleteLocation ---->>>")

	location, err := s.stg.GetLocationByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetLocationByID: %s", err)
	}

	err = s.stg.DeleteLocation(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteLocation: %s", err)
	}

	return &book_service.DeleteLocationResponse{
		Status: location.Status,
	}, nil
}

func (s *locationService) EnabledLocation(ctx context.Context, req *book_service.EnabledLocationRequest) (*book_service.EnabledLocationResponse, error) {
	fmt.Println("<<< ---- EnabledLocation ---->>>")

	err := s.stg.EnabledLocation(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.EnabledLocation: %s", err)
	}

	location, err := s.stg.GetLocationByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetLocationByID: %s", err)
	}

	return &book_service.EnabledLocationResponse{
		Status: location.Status,
	}, nil
}

func (s *locationService) GetLocationList(ctx context.Context, req *book_service.GetLocationListRequest) (*book_service.GetLocationListResponse, error) {
	fmt.Println("<<< ---- GetLocationList ---->>>")

	res, err := s.stg.GetLocationList(int(req.Offset), int(req.Limit), req.Search)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetLocationList: %s", err)
	}

	return res, nil
}

func (s *locationService) GetLocationById(ctx context.Context, req *book_service.GetLocationByIdRequest) (*book_service.Location, error) {
	fmt.Println("<<< ---- GetLocationById ---->>>")

	location, err := s.stg.GetLocationByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetLocationByID: %s", err)
	}

	return &book_service.Location{
		Id:        location.Id,
		Name:      location.Name,
		Status:    location.Status,
		CreatedAt: location.CreatedAt,
		UpdatedAt: location.UpdatedAt,
	}, nil
}
