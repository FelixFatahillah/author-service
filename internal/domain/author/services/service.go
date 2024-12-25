package services

import (
	"author-service/internal/domain/author/dtos"
	"author-service/internal/domain/author/models"
	"author-service/internal/domain/author/repositories"
	"author-service/pkg/helper"
	"context"
	"github.com/redis/go-redis/v9"
)

type serviceAuthor struct {
	Repository repositories.AuthorRepository
	Redis      redis.Client
}

type AuthorService interface {
	Create(ctx context.Context, input dtos.CreateAuthorDto) (*dtos.CreateAuthorResponseDto, error)
	GetAll(ctx context.Context, filter dtos.AuthorFilter) ([]models.Author, *helper.PaginationMeta, error)
	FindById(ctx context.Context, id string) (*models.Author, error)
	Update(ctx context.Context, input dtos.UpdateAuthorDto) error
	Delete(ctx context.Context, id string) error
}
