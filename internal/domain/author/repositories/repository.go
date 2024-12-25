package repositories

import (
	"author-service/internal/domain/author/dtos"
	"author-service/internal/domain/author/models"
	"author-service/pkg/helper"
	"context"
	"gorm.io/gorm"
)

type repositoryAuthor struct {
	db *gorm.DB
}

type AuthorRepository interface {
	Transaction(ctx context.Context, fn func(repo AuthorRepository) error) error
	Create(ctx context.Context, user models.Author) (*models.Author, error)
	GetAll(ctx context.Context, filter dtos.AuthorFilter) ([]models.Author, *helper.PaginationMeta, error)
	FindById(ctx context.Context, id string) (*models.Author, error)
	Update(ctx context.Context, input *models.Author) error
	Delete(ctx context.Context, id string) error
}
