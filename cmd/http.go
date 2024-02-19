package cmd

import (
	"github.com/spf13/cobra"
	_ "notify/docs"
	"notify/internal/app"
)

// @BasePath /api/
// @title Notifier API
// @version 0.0.1
// @Description Service for sending notifications to users

// @contact.name Dima Kalinin
// @contact.url https://dimdimi4-and-fixiki.github.io/dima-portal/
// @contact.email kalinindima123@yandex.ru

func RunHTTP() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "http",
		Short: "Run http server",
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		a, err := app.GetGlobalApp()
		if err != nil {
			return err
		}

		a.StartHTTPServer()
		return nil
	}

	return cmd
}
