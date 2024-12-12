package app

import (
	"github.com/antitokens/priceindex/api"
	"github.com/antitokens/priceindex/model"
	"github.com/antitokens/priceindex/source"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	Source source.Source
	Router *fiber.App
	DB     *gorm.DB
}

func New() *App {
	source := source.NewRaydium()
	router := fiber.New()

	envPath := ".env"
	env, err := godotenv.Read(envPath)
	if err != nil {
		panic(err)
	}

	dbUrl := env["PRICEINDEX_DB_URL"]

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &App{Source: source, Router: router, DB: db}
}

func (a *App) Start() {
	a.Router.Get("/price/:token", api.GetPriceHandler(a.Source, a.DB))
	a.Router.Listen(":3000")
}

func (a *App) Migrate() {
	a.DB.AutoMigrate(&model.Price{})
}
