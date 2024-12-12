package api

import (
	"net/http"

	"github.com/antitokens/priceindex/source"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetPriceHandler(source source.Source, db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Params("token")
		priceResponse, err := source.GetPrice(token)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(priceResponse)
	}
}

func GetHistoryHandler(source source.Source, db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"error": "Not implemented"})
	}
}
