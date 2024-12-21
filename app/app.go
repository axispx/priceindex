package app

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/antitokens/priceindex/api"
	"github.com/antitokens/priceindex/config"
	"github.com/antitokens/priceindex/source"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	Source     source.Source
	Router     *fiber.App
	DB         *gorm.DB
	Config     *config.Config
	ApiHandler *api.ApiHandler
}

func New() *App {
	source := source.NewRaydium()
	router := fiber.New()
	cfg := config.LoadConfig()

	connStr := cfg.DB.ConnectionString
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	apiHandler := api.NewApiHandler(db)

	return &App{Source: source, Router: router, DB: db, Config: cfg, ApiHandler: apiHandler}
}

func (a *App) Start() {
	go a.indexTokenPrice()

	apiHandler := a.ApiHandler

	a.Router.Get("/price/:token", apiHandler.GetPrice)
	a.Router.Get("/price/hourly/:token", apiHandler.GetHourlyPrice)
	a.Router.Get("/price/daily/:token", apiHandler.GetDailyPrice)
	a.Router.Listen(":3000")
}

func (a *App) Migrate() {
	migrationDir, err := filepath.Abs(os.Getenv("PRICEINDEX_DB_MIGRATIONS_DIR"))
	if err != nil {
		log.Fatalln(err)
	}

	migrator, err := migrate.New("file://"+migrationDir, a.Config.DB.ConnectionString)
	if err != nil {
		log.Fatalln(err)
	}

	if err := migrator.Up(); err != nil {
		log.Fatalln(err)
	}
}

func (a *App) indexTokenPrice() {
	ticker := time.NewTicker(a.Config.IndexInterval)
	defer ticker.Stop()

	for range ticker.C {
		prices, err := a.Source.GetPrice("anti", "pro")
		if err != nil {
			panic(err)
		}

		if err := a.DB.Create(&prices).Error; err != nil {
			panic(err)
		}
	}
}
