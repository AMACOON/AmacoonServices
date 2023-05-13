package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"github.com/sirupsen/logrus"
)

func SetupDB(config *Config, logger *logrus.Logger) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	var db *gorm.DB
	var err error

	for retries := 5; retries >= 0; retries-- {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
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



func SetupDBOld(config *Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"amacoon001_add1",
		"armin013",
		"mysql.catclubsystem.com",
		config.DBPort,
		"amacoon01",
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		
	}

	return db, nil
}
