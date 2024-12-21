package api

import (
	"net/http"
	"time"

	"github.com/antitokens/priceindex/model"
	"github.com/antitokens/priceindex/source"
	"github.com/antitokens/priceindex/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ApiHandler struct {
	db *gorm.DB
}

func NewApiHandler(db *gorm.DB) *ApiHandler {
	return &ApiHandler{db: db}
}

func (ah *ApiHandler) GetPrice(c *fiber.Ctx) error {
	token := c.Params("token")
	address := utils.GetTokenAddress(token)

	var price model.Price
	if err := ah.db.Where("address = ?", address).Order("timestamp DESC").First(&price).Error; err != nil {
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

func (ah *ApiHandler) GetHourlyPrice(c *fiber.Ctx) error {
	start := c.Query("start")
	end := c.Query("end")

	if _, err := time.Parse("2006-01-02", start); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid start date"})
	}

	if _, err := time.Parse("2006-01-02", end); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid end date"})
	}

	token := c.Params("token")
	address := utils.GetTokenAddress(token)

	var prices []model.HourlyPrice
	if err := ah.db.Table("hourly_prices").Where("address = ? AND date(hour) >= ? AND date(hour) <= ?", address, start, end).Find(&prices).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "No price history found"})
	}

	var pricesResponse []model.HourlyPriceResponse
	for _, price := range prices {
		pricesResponse = append(pricesResponse, model.HourlyPriceResponse{
			price.Hour.Unix(), price.Address, price.AvgPrice.String(),
		})
	}

	return c.JSON(pricesResponse)
}

func (ah *ApiHandler) GetDailyPrice(c *fiber.Ctx) error {
	start := c.Query("start")
	end := c.Query("end")

	if _, err := time.Parse("2006-01-02", start); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid start date"})
	}

	if _, err := time.Parse("2006-01-02", end); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid end date"})
	}

	token := c.Params("token")
	address := utils.GetTokenAddress(token)

	var prices []model.DailyPrice
	if err := ah.db.Table("daily_prices").Where("address = ? AND date(day) >= ? AND date(day) <= ?", address, start, end).Find(&prices).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "No price history found"})
	}

	var pricesResponse []model.DailyPriceResponse
	for _, price := range prices {
		pricesResponse = append(pricesResponse, model.DailyPriceResponse{
			price.Day.Unix(), price.Address, price.AvgPrice.String(),
		})
	}

	return c.JSON(pricesResponse)
}

func GetHistoryHandler(source source.Source, db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"error": "Not implemented"})
	}
}
