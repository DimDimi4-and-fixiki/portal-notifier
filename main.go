package main

import (
	"log"

	"github.com/spf13/cobra"

	"notify/cmd"
	"notify/internal/app"
)

func initApp() {
	a, err := app.NewApp()
	if err != nil {
		log.Fatal("Fail to create app: ", err)
	}

	app.SetGlobalApp(a)
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "Run HTTP Server with notifier",
		Short: "Main entry-point command for the application",
	}

	cobra.OnInitialize(initApp)

	rootCmd.AddCommand(
		cmd.RunHTTP(),
	)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("failed to execute root cmd: %v", err)

		return
	}
}
