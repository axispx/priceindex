package app

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/antitokens/priceindex/api"
	"github.com/antitokens/priceindex/config"
	"github.com/antitokens/priceindex/model"
	"github.com/antitokens/priceindex/source"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	Source source.Source
	Router *fiber.App
	DB     *gorm.DB
	Config *config.Config
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

	return &App{Source: source, Router: router, DB: db, Config: cfg}
}

func (a *App) Start() {
	go a.indexTokenPrice()

	a.Router.Get("/price/:token", api.GetPriceHandler(a.Source, a.DB))
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
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		priceAnti, err := a.Source.GetPrice("anti")
		if err != nil {
			panic(err)
		}

		pricePro, err := a.Source.GetPrice("pro")
		if err != nil {
			panic(err)
		}

		prices := []model.Price{priceAnti, pricePro}

		if err := a.DB.Create(&prices).Error; err != nil {
			panic(err)
		}
	}
}
