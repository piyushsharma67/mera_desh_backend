package cmd

import (
	"log"
	"os"
	configPkg "social_web_server/config"

	"github.com/spf13/cobra"
)

var env string

var rootCmd = &cobra.Command{
	Use:   "social_web_app",
	Short: "This is my Go server",
	Long:  "This is my Go server",
	Run: func(cmd *cobra.Command, args []string) {
		if env == "" {
			log.Fatal("Environment flag (--env) must be specified")
			os.Exit(1)
		}
		_, err := configPkg.Loadconfig(env)
		if err != nil {
			log.Fatalf("Failed to load configuration: %v", err)
		}
	},
}

// Execute runs the root command
func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	return nil
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&env, "env", "e", "", "Specify the environment: local, dev, stage, prod")
}

// AddCommand allows other command files to register their commands
func AddCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}
