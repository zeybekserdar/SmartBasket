package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
)

type IProductHandler interface {
	Get(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type handler struct {
	service IProductService
}

func (h handler) Update(ctx *fiber.Ctx) error {
	var id = ctx.Params("id")
	model, err := h.service.Get(uuid.FromStringOrNil(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(Response{Error: err.Error()})
	}

	newProduct := Product{}
	err = ctx.BodyParser(&newProduct)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(Response{Error: err.Error()})
	}

	updatedId, err := h.service.Update(*model, newProduct)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(Response{Error: err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(Response{Data: updatedId})
}

func (h handler) Delete(ctx *fiber.Ctx) error {
	var id = ctx.Params("id")
	model, err := h.service.Get(uuid.FromStringOrNil(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(Response{Error: err.Error()})
	}
	deletedId, err := h.service.Delete(*model)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(Response{Error: err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(Response{Data: deletedId})
}

func (h handler) Get(ctx *fiber.Ctx) error {
	var id = ctx.Params("id")
	model, err := h.service.Get(uuid.FromStringOrNil(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(Response{Error: err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(Response{Data: model})

}

func (h handler) GetAll(ctx *fiber.Ctx) error {
	products, err := h.service.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(Response{Error: err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(Response{Data: products})

}

func (h handler) Create(ctx *fiber.Ctx) error {
	product := Product{}
	err := ctx.BodyParser(&product)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(Response{Error: err.Error()})
	}

	validationErr, err := ValidateProduct(product)
	if validationErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(validationErr)
	}
	newProduct, errSvc := h.service.Create(product)
	if errSvc != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(Response{Error: errSvc.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(Response{Data: ProductDto{Name: newProduct.Name, Price: newProduct.Price, Type: newProduct.Type}})
}

var _ IProductHandler = handler{}

func NewHandler(service IProductService) IProductHandler {
	return handler{service: service}
}

type Response struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

type ProductDto struct {
	Name  string `json:"name"`
	Price string `json:"price"`
	Type  string `json:"type"`
}
