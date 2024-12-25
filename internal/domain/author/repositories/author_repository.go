package repositories

import (
	"author-service/internal/domain/author/dtos"
	"author-service/internal/domain/author/models"
	"author-service/pkg/helper"
	"author-service/pkg/logger"
	"context"
	"gorm.io/gorm"
)

func NewRepositoryAuthor(db *gorm.DB) *repositoryAuthor {
	return &repositoryAuthor{db}
}

func (repository repositoryAuthor) beginTransaction() *gorm.DB { return repository.db.Begin() }

func (repository repositoryAuthor) withTx(ctx context.Context, tx *gorm.DB) *repositoryAuthor {
	repository.db = tx.WithContext(ctx)
	return &repository
}

func (repository repositoryAuthor) Transaction(ctx context.Context, fn func(repo AuthorRepository) error) error {
	tx := repository.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	repo := repository.withTx(ctx, tx)
	err := fn(repo)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (repository repositoryAuthor) Create(ctx context.Context, user models.Author) (*models.Author, error) {
	if err := repository.db.WithContext(ctx).Create(&user).Error; err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return &user, nil
}

func (repository repositoryAuthor) GetAll(ctx context.Context, filter dtos.AuthorFilter) ([]models.Author, *helper.PaginationMeta, error) {
	records := make([]models.Author, 0)
	paginateMeta := helper.PaginationMeta{
		Page:  filter.Pagination.Page,
		Limit: filter.Pagination.Limit,
	}

	query := repository.db.WithContext(ctx).Model(models.Author{}).Scopes(helper.PaginateScope(&filter.Pagination))
	query.Count(&paginateMeta.Total).Order("created_date DESC").Find(&records)

	paginateMeta.TotalPage = helper.GetTotalPage(paginateMeta.Total, paginateMeta.Limit)

	return records, &paginateMeta, nil
}

func (repository repositoryAuthor) FindById(ctx context.Context, id string) (*models.Author, error) {
	record := &models.Author{}
	err := repository.db.WithContext(ctx).Take(record, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return record, nil
}

func (repository repositoryAuthor) Update(ctx context.Context, input *models.Author) error {
	if err := repository.db.
		WithContext(ctx).
		Updates(input).Error; err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (repository repositoryAuthor) Delete(ctx context.Context, id string) error {
	err := repository.db.WithContext(ctx).Unscoped().Where("id = ?", id).Delete(&models.Author{}).Error
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
