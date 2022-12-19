package main

import (
	"fmt"
	"lms/lms_book_service/config"
	"lms/lms_book_service/storage"
	"lms/lms_book_service/storage/postgres"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/lib/pq"
	"github.com/swaggo/swag/example/basic/docs"
)

func initGRPC(cfg config.Config, stg storage.Interfaces) {

}

// @license.name	Apache 2.0
func main() {

	cfg := config.Load()

	psqlConfigString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	docs.SwaggerInfo.Title = cfg.App
	docs.SwaggerInfo.Version = cfg.AppVersion

	var err error
	var stg storage.Interfaces

	stg, err = postgres.InitDB(psqlConfigString)
	if err != nil {
		panic(err)
	}

	fmt.Printf("gRPC server tutorial in Go in GRPCPort: %s", cfg.GRPCPort)

	// gRPC port connection ...
	listener, err := net.Listen("tcp", cfg.GRPCPort)

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	CategoryService := category.NewCategoryService(stg)
	eCommerce.RegisterCategoryServiceServer(s, CategoryService)

	ProductService := products.NewProductService(stg)
	eCommerce.RegisterProductServiceServer(s, ProductService)

	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
