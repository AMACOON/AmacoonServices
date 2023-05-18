package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/scuba13/AmacoonServices/cmd/server/initialize"
	"github.com/scuba13/AmacoonServices/cmd/server/migrate"
	"github.com/scuba13/AmacoonServices/cmd/server/setup"
	"github.com/scuba13/AmacoonServices/config"
)

func main() {
	// Initialize Logger
	logger := setup.SetupLogger()

	// Initialize Echo
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${remote_ip} ${method} ${uri} ${status} ${error}\n",
		Output: logger.Out,
	}))

	// Load configuration data
	cfg := config.LoadConfig()

	//Initialize S3
	s3 := setup.SetupS3(cfg, logger)

	// Initialize DB
	db := setup.SetupDatabase(cfg, logger)
	dbOld := setup.SetupDatabaseOld(cfg, logger)

	// Migrate data
	MigrateService := migrate.NewMigrateService(db, dbOld, logger)
	migrate.SetupRouter(MigrateService, logger, e)

	// Initialize repositories, handlers, and routes
	initialize.InitializeApp(e, logger, db, s3, cfg)

	// Start server
	logger.Info("Starting Server")
	if err := e.Start(":" + cfg.ServerPort); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
