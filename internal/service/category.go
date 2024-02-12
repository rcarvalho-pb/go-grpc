package service

import (
	"github.com/rcarvalho-pb/go-grpc/internal/database"
	"github.com/rcarvalho-pb/go-grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDb database.Category) *CategoryService {
	return &CategoryService{CategoryDB: categoryDb}
}
