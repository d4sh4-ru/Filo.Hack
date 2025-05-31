package main

import (
	"log"
	"os"

	"Filo.Hack/config"
	"Filo.Hack/internal/app/router"
	"Filo.Hack/internal/lib/storage"
	"github.com/labstack/echo/v4"
)

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is empty")
	}

	cfg := config.MustLoad(configPath)

	dbClient, err := storage.NewDBClient(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize DB client")
	}
	defer func() {
		err := storage.CloseDBConnection(dbClient.Db)
		if err != nil {
			log.Fatalf("Failed to close DB connection")
		}
		log.Print("DB connection closed")
	}()

	log.Print("DB connection initialized")

	e := echo.New()

	router.RegisterRouters(e, dbClient, cfg)

	if err := e.Start(cfg.HTTPServer.Port); err != nil {
		log.Fatalf("HTTP server critical error")
	}
}
