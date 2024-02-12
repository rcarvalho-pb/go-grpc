package service

import (
	"context"

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

func (c *CategoryService) CreateCategory(ctx context.Context, request *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(request.Name, request.Description)
	if err != nil {
		return nil, err
	}

	return &pb.CategoryResponse{
		Category: &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		},
	}, nil

}
