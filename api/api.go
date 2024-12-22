package api

import (
	"net/http"
	"time"

	"github.com/antitokens/priceindex/model"
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
	address := utils.GetTokenAddresses(token)

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

func (ah *ApiHandler) GetPriceHistory(c *fiber.Ctx) error {
	start := c.Query("start")
	end := c.Query("end")

	token := c.Params("token")
	address := utils.GetTokenAddress(token)

	tx := ah.db.Table("prices").Where("address = ?", address)

	// If start or end isn't provided, we assume the dates to be the oldest date
	// and the latest date repectively. We check for validity if they are not empty.
	if start != "" {
		if _, err := time.Parse("2006-01-02", start); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid start date"})
		}

		tx = tx.Where("date(timestamp) >= ?", start)
	}

	if end != "" {
		if _, err := time.Parse("2006-01-02", end); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid end date"})
		}

		tx = tx.Where("date(timestamp) <= ?", end)
	}

	var prices []model.Price
	if err := tx.Order("timestamp").Find(&prices).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "No price history found"})
	}

	pricesResponse := []model.PriceResponse{}
	for _, price := range prices {
		pricesResponse = append(pricesResponse, model.PriceResponse{
			price.Price.String(), price.Timestamp.Unix(), price.Source, price.Address,
		})
	}

	return c.JSON(pricesResponse)
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
	err := ah.db.Table("hourly_prices").
		Where("address = ? AND date(hour) >= ? AND date(hour) <= ?", address, start, end).
		Order("hour").
		Find(&prices).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "No price history found"})
	}

	pricesResponse := []model.HourlyPriceResponse{}
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
	err := ah.db.Table("daily_prices").
		Where("address = ? AND date(day) >= ? AND date(day) <= ?", address, start, end).
		Order("day").
		Find(&prices).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "No price history found"})
	}

	pricesResponse := []model.DailyPriceResponse{}
	for _, price := range prices {
		pricesResponse = append(pricesResponse, model.DailyPriceResponse{
			price.Day.Unix(), price.Address, price.AvgPrice.String(),
		})
	}

	return c.JSON(pricesResponse)
}
