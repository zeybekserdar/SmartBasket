package product

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Get(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
}

type handler struct {
	service Service
}

func (h handler) Get(ctx *fiber.Ctx) error {
	id,err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).JSON(Response{Error: err.Error()})
	}

	model,err := h.service.Get(uint(id))
	if err != nil {
		return ctx.Status(404).JSON(Response{Error: err.Error()})
	}

	return ctx.Status(200).JSON(Response{Data: model})

}

func (h handler) Create(ctx *fiber.Ctx) error {
	model := Product{}
	err := ctx.BodyParser(&model)
	if err != nil {
		return ctx.Status(400).JSON(Response{Error: err.Error()})
	}
	_,err = h.service.Create(model)
	if err != nil {
		ctx.Status(400).JSON(Response{Error: err.Error()})
	}
	return ctx.SendStatus(201)
}

var _ Handler = handler{}

func NewHandler(service Service) Handler {
	return handler{service: service}
}

type Response struct {
	Error string 		`json:"error"`
	Data interface{}	`json:"data"`
}