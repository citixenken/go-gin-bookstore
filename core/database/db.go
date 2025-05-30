package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

type DBClient interface {
	DBMigrate() error
}

type Client struct {
	db *gorm.DB
}

func NewClient() (DBClient, error) {
	databaseHost := os.Getenv("DB_HOST")
	databaseUserName := os.Getenv("DB_USERNAME")
	databasePassword := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	databasePort := os.Getenv("DB_PORT")
	dbPort, err := strconv.Atoi(databasePort)

	if err != nil {
		log.Fatal("Invalid DB Port")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		databaseHost, databaseUserName, databasePassword, databaseName, dbPort, "disable")

	dbInfo, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	client := Client{db: dbInfo}
	return client, nil
}

func (c Client) DBMigrate() error {
	return nil
}

func (c Client) CloseDBConnection() {
	db, err := c.db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	db.Close()
}
