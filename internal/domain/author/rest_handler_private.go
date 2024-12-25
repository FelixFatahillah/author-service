package author

import (
	"author-service/internal/domain/author/dtos"
	"author-service/pkg/helper"
	"author-service/pkg/response"
	"author-service/pkg/validation"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func (handler handlerRESTAuthor) handlerCreate(ctx *fiber.Ctx) error {
	var body dtos.CreateAuthorDto
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}
	if err := validation.Validate(body); err != nil {
		return err
	}

	data, err := handler.service.Create(ctx.Context(), body)
	if err != nil {
		return err
	}

	return response.Respond(ctx, fiber.StatusOK, "success", data, nil, nil)
}

func (handler handlerRESTAuthor) handlerGetAll(ctx *fiber.Ctx) error {
	paginate := helper.Pagination{
		Page:  1,
		Limit: 10,
	}
	err := ctx.QueryParser(&paginate)
	if paginate.Limit >= 100 {
		paginate.Limit = 100
	}

	filter := dtos.AuthorFilter{
		Pagination: paginate,
	}
	_ = ctx.QueryParser(&filter)

	data, meta, err := handler.service.GetAll(ctx.Context(), filter)
	if err != nil {
		return err
	}

	return response.Respond(ctx, fiber.StatusOK, "success", data, nil, meta)
}

func (handler handlerRESTAuthor) handlerFindById(ctx *fiber.Ctx) error {
	data, err := handler.service.FindById(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return err
	}

	return response.Respond(ctx, fiber.StatusOK, "success", data, nil, nil)
}

func (handler handlerRESTAuthor) handlerUpdate(ctx *fiber.Ctx) error {
	var body dtos.UpdateAuthorDto
	body.ID = ctx.Params("id")
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}
	if err := validation.Validate(body); err != nil {
		return err
	}

	err := handler.service.Update(ctx.Context(), body)
	if err != nil {
		return err
	}

	return response.Respond(ctx, fiber.StatusOK, fmt.Sprintf("success update %s", ctx.Params("id")), nil, nil, nil)
}

func (handler handlerRESTAuthor) handlerDelete(ctx *fiber.Ctx) error {
	err := handler.service.Delete(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return err
	}

	return response.Respond(ctx, fiber.StatusOK, fmt.Sprintf("success deleted %s", ctx.Params("id")), nil, nil, nil)
}
