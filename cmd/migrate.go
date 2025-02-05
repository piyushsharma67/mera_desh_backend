package cmd

import (
	"database/sql"
	"fmt"
	"log"
	configPkg "social_web_server/config"

	_ "github.com/lib/pq"
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
	fmt.Println(env)
	// Connect to PostgreSQL using the standard sql package
	config,err:=configPkg.Loadconfig(env)
	if err!=nil{
		log.Fatal(err)
	}
	dbConn, err := sql.Open("postgres", config.GetDSN())
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
