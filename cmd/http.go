package cmd

import (
	"github.com/spf13/cobra"
	"notify/internal/app"
)

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
