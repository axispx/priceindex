package cmd

import (
	"fmt"
	"os"

	"github.com/antitokens/priceindex/app"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "priceindex",
	Long: "priceindex is a tool for indexing historical prices for anti-tokens.",
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the database",
}

func Execute() {
	app := app.New()

	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		app.Start()
	}

	migrateCmd.Run = func(cmd *cobra.Command, args []string) {
		app.Migrate()
	}

	rootCmd.AddCommand(migrateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
