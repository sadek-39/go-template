package cmd

import (
	"fmt"
	"os"

	"github.com/sadek-39/basic-go-template/internal/config"
	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "root",
	}
)

func init() {
	RootCmd.AddCommand(serveCmd)
}

func Execute() {
	config.LoadConfig()
	
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}