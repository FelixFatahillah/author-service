package author

import (
	"author-service/internal/domain/author/services"
	"author-service/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

type handlerRESTAuthor struct {
	service services.AuthorService
}

func NewHandlerRESTAuthor(service services.AuthorService, router *fiber.App) {
	handler := handlerRESTAuthor{
		service,
	}

	routerV1 := router.Group("/api/v1")
	routerProtected := routerV1.Group("/private/authors", middleware.RoleMiddleware())

	routerProtected.Delete("/:id", handler.handlerDelete)
	routerProtected.Get("/:id", handler.handlerFindById)
	routerProtected.Put("/:id", handler.handlerUpdate)
	routerProtected.Post("/", handler.handlerCreate)
	routerProtected.Get("/", handler.handlerGetAll)
}
