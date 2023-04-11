package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"time"
	

)

func SetupDB(config *Config) (*gorm.DB, error) {
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	
	return db, nil
}

func SetupMongo(config *Config) (*mongo.Client, error) {
	// Define a string de conexão com o MongoDB
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		config.MongoDBUsername,
		config.MongoDBPassword,
		config.MongoDBHost,
		config.MongoDBPort,
		config.MongoDBName,
	)


	// Define as opções de conexão com o MongoDB
	opts := options.Client().ApplyURI(mongoURI)

	// Cria um contexto com timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Conecta ao MongoDB
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Testa a conexão com o MongoDB
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	return client, nil
}

