package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/scuba13/AmacoonServices/cmd/server/initialize"
	"github.com/scuba13/AmacoonServices/cmd/server/migrate"
	"github.com/scuba13/AmacoonServices/cmd/server/setup"
	"github.com/scuba13/AmacoonServices/config"
	"github.com/spf13/viper"
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
	config.LoadConfig(logger)

	//Initialize S3
	s3 := setup.SetupS3(logger)

	// Initialize DB
	dbOld := setup.SetupDatabaseOld(logger)
	db, err := setup.SetupDatabase(logger)
	if err != nil {
		// Lidar com o erro aqui, por exemplo, registrar e encerrar o programa
		logger.Fatalf("Database setup failed: %v", err)
	}

	// Migrate data
	MigrateService := migrate.NewMigrateService(db, dbOld, logger)
	migrate.SetupRouter(MigrateService, logger, e)

	// Initialize repositories, handlers, and routes
	initialize.InitializeApp(e, logger, db, s3)

	// Start server
	logger.Info("Starting Server")
	if err := e.Start(":" + viper.GetString("server.port")); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
