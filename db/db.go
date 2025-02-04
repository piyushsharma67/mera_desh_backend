package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	DBHost     string `json:"db_host"`
	DBPort     string `json:"db_port"`
	DBUser     string `json:"db_user"`
	DBPassword string `json:"db_password"`
	DBName     string `json:"db_name"`
	DBSSLMode  string `json:"db_sslmode"`
}

var DB *pgxpool.Pool
var config Config

// LoadConfig reads database config from a JSON file
func LoadConfig() {
	file, err := os.Open("config.json") // Modify this to your config file path
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatalf("Error decoding config file: %v", err)
	}
}

// GetDSN constructs the PostgreSQL DSN
func GetDSN() string {
	LoadConfig()
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName, config.DBSSLMode,
	)
}

// ConnectDB initializes the database connection
func ConnectDB() {
	 // Load database config

	dsn := GetDSN()
	dbPool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	DB = dbPool
	fmt.Println("✅ Successfully connected to the database!")
}