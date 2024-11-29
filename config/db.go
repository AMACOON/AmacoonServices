package config

import (
	"fmt"
	"log"

	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB(logger *logrus.Logger) (*gorm.DB, error) {
    dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=require",
        viper.GetString("DB_HOST"),
        viper.GetString("DB_PORT"),
        viper.GetString("DB_USERNAME"),
        viper.GetString("DB_NAME"),
        viper.GetString("DB_PASSWORD"),
    )


	var db *gorm.DB
	var err error

	for retries := 5; retries >= 0; retries-- {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}

		logger.Warnf("Failed to connect to database: %v", err)

		if retries > 0 {
			logger.Infof("Retrying in 5 seconds...")
			time.Sleep(time.Second * 5)
		}
	}

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Set pool configurations (you can adjust these numbers based on your application's needs)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func SetupDBOld() (*gorm.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"amacoon001_add1",
		"armin013",
		"mysql.catclubsystem.com",
		"3306",
		"amacoon01",
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)

	}

	return db, nil
}
