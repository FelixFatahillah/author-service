package delivery

import (
	repositoryAuthor "author-service/internal/domain/author/repositories"
	"gorm.io/gorm"
)

type Repositories struct {
	AuthorRepository repositoryAuthor.AuthorRepository
}

func NewRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		AuthorRepository: repositoryAuthor.NewRepositoryAuthor(db),
	}
}
