package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/sadek-39/basic-go-template/internal/config"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Run: serve,
}

func serve(cmd *cobra.Command, args []string) {
	var e = echo.New()

	port := config.App().Port

	e.Logger.Fatal(e.Start(":" + port))
}