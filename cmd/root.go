package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "social_web_app",
	Short: "This is my Go server",
	Long:  "This is my Go server",
}

// Execute runs the root command
func Execute() error{
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}

// AddCommand allows other command files to register their commands
func AddCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}
