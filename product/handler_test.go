package product

import (
	"SmartBasket/db"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestGetHandler(t *testing.T){
	database, err := db.Connect()
	assert.Nil(t, err)
	repository := NewRepository(database)
	service := NewService(repository)
	handler := NewHandler(service)

	app := fiber.New()
	app.Get("/products/:id",handler.Get)

	id,err := repository.Create(Product{Name: "Test Product",Price: "25",Type: "Test Type"})
	assert.Nil(t, err)

	req := httptest.NewRequest("GET",fmt.Sprintf("/products/%d",id),nil)
	res,err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 200,res.StatusCode)

}