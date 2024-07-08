package controller

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type IExampleController interface {
	GetExample(c *fiber.Ctx) error
}

type ExampleController struct {
	DB *gorm.DB
}

func NewExampleController(db *gorm.DB) *ExampleController {
	return &ExampleController{DB: db}
}

func (r *ExampleController) GetExample(c *fiber.Ctx) error {
	return c.JSON("Get Example")
}
