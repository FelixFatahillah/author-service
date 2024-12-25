package delivery

import (
	servicesAuthor "author-service/internal/domain/author/services"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	AuthorService servicesAuthor.AuthorService
}

type Deps struct {
	Repository *Repositories
	Redis      redis.Client
	//GRPC       *GRPC
}

func NewService(deps Deps) *Service {
	return &Service{
		AuthorService: servicesAuthor.NewServiceAuthor(deps.Repository.AuthorRepository, deps.Redis),
	}
}
