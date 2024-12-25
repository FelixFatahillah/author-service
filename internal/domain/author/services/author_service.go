package services

import (
	"author-service/internal/domain/author/dtos"
	"author-service/internal/domain/author/models"
	"author-service/internal/domain/author/repositories"
	"author-service/internal/shared"
	"author-service/pkg/exception"
	"author-service/pkg/helper"
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

func NewServiceAuthor(repository repositories.AuthorRepository, redis redis.Client) *serviceAuthor {
	return &serviceAuthor{
		Repository: repository,
		Redis:      redis,
	}
}

func (service serviceAuthor) Create(ctx context.Context, input dtos.CreateAuthorDto) (*dtos.CreateAuthorResponseDto, error) {
	record, err := service.Repository.Create(ctx, models.Author{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		PhoneNumber: input.PhoneNumber,
		Email:       input.Email,
		Biography:   input.Biography,
		Education:   input.Education,
	})
	if err != nil {
		return nil, err
	}
	return &dtos.CreateAuthorResponseDto{
		FirstName:   record.FirstName,
		LastName:    record.LastName,
		PhoneNumber: record.PhoneNumber,
		Email:       record.Email,
		Biography:   record.Biography,
		Education:   record.Education,
	}, nil
}
func (service serviceAuthor) GetAll(ctx context.Context, filter dtos.AuthorFilter) ([]models.Author, *helper.PaginationMeta, error) {
	return service.Repository.GetAll(ctx, filter)
}

func (service serviceAuthor) FindById(ctx context.Context, id string) (*models.Author, error) {
	return service.Repository.FindById(ctx, id)
}

func (service serviceAuthor) Update(ctx context.Context, input dtos.UpdateAuthorDto) error {
	err := service.Repository.Transaction(ctx, func(repo repositories.AuthorRepository) error {
		record, _ := service.Repository.FindById(ctx, input.ID)
		if record == nil {
			return &exception.ErrWithCode{
				Code: http.StatusNotFound,
				Err:  errors.New("category not found"),
			}
		}
		_ = service.Repository.Update(ctx,
			&models.Author{
				BaseModel: shared.BaseModel{
					UpdatedDate: time.Now().Local(),
				},
				FirstName:   input.FirstName,
				LastName:    input.LastName,
				PhoneNumber: input.PhoneNumber,
				Email:       input.Email,
				Biography:   input.Biography,
				Education:   input.Education,
			},
		)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (service serviceAuthor) Delete(ctx context.Context, id string) error {
	category, _ := service.Repository.FindById(ctx, id)
	if category == nil {
		return &exception.ErrWithCode{
			Code: http.StatusNotFound,
			Err:  errors.New("author not found"),
		}
	}
	return service.Repository.Delete(ctx, id)
}
