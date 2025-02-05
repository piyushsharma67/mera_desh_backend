package db

import (
	"context"
	"fmt"
	"log"
	configPkg "social_web_server/config"

	"github.com/jackc/pgx/v5/pgxpool"
)
var DB *pgxpool.Pool

// ConnectDB initializes the database connection
func ConnectDB(env string) {
	 // Load database config

	dsn,err := configPkg.Loadconfig(env)
	if err!=nil{
		log.Fatal(err)
	}
	dbPool, err := pgxpool.New(context.Background(), dsn.GetDSN())
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	DB = dbPool
	fmt.Println("✅ Successfully connected to the database!")
}