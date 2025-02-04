package cmd

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq" 
	"social_web_server/db"
	"github.com/pressly/goose"
	"github.com/spf13/cobra"
)

var migrateCMD = &cobra.Command{
	Use:   "migrate",
	Short: "run the migrtations",
	Long:  "Run the migration files",
	Run: func(cmd *cobra.Command, args []string) {
		runMigrations()
	},
}

func init() {
	AddCommand(migrateCMD)
}

func runMigrations() {
	dsn := db.GetDSN() // Get DSN from config
	// Connect to PostgreSQL using the standard sql package
	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	err = goose.SetDialect("postgres")

	if err != nil {
		log.Fatalf("Failed to set Goose dialect: %v", err)
	}

	// Apply migrations from the "db/migrations" directory
	migrationsDir := "db/migrations"
	err = goose.Up(dbConn, migrationsDir)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("✅ Database migration applied successfully!")
}
