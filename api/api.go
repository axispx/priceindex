package api

import (
	"net/http"

	"github.com/antitokens/priceindex/model"
	"github.com/antitokens/priceindex/source"
	"github.com/antitokens/priceindex/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetPriceHandler(source source.Source, db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Params("token")
		address := utils.GetTokenAddress(token)

		var price model.Price
		if err := db.Where("address = ?", address).Order("timestamp DESC").First(&price).Error; err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Price not found"})
		}

		priceResponse := model.PriceResponse{
			Price:     price.Price.String(),
			Timestamp: price.Timestamp.Unix(),
			Source:    price.Source,
			Address:   price.Address,
		}

		return c.JSON(priceResponse)
	}
}

func GetHistoryHandler(source source.Source, db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"error": "Not implemented"})
	}
}
