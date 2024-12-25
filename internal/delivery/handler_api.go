package delivery

import (
	"author-service/internal/domain/author"
	servicesAuthor "author-service/internal/domain/author/services"
	"author-service/pkg/exception"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type handler struct {
	serviceAuthor servicesAuthor.AuthorService
}

func NewHandler(serviceAuthor servicesAuthor.AuthorService) *handler {
	return &handler{
		serviceAuthor: serviceAuthor,
	}
}

const idleTimeout = 60 * time.Second

func (handler *handler) Init() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.FiberErrorHandler,
		IdleTimeout:  idleTimeout,
		JSONEncoder:  sonic.Marshal,
		JSONDecoder:  sonic.Unmarshal,
	})

	// Middleware
	app.Use(logger.New())
	app.Use(etag.New())
	app.Use(requestid.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
	author.NewHandlerRESTAuthor(handler.serviceAuthor, app)

	return app
}
