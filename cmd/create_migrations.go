package cmd

import (
	"fmt"
	"github.com/pressly/goose"
	"github.com/spf13/cobra"
)

var migrateCreateCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a new migration file",
	Long:  "Create a new migration file with the specified name inside the db/migrations folder",
	Args:  cobra.ExactArgs(1), // Ensure one argument is passed
	Run: func(cmd *cobra.Command, args []string) {
		// Get the name of the migration file
		name := args[0]
		dir := "database/migrations" // Migration directory

		// Create the migration file using Goose
		err := goose.Create(nil, dir, name, "sql")
		if err != nil {
			fmt.Printf("Error creating migration: %v\n", err)
			return
		}

		fmt.Printf("✅ Migration file created successfully: %s\n", name)
	},
}

func init() {
	// Add the create command to the root command
	AddCommand(migrateCreateCmd)
}
