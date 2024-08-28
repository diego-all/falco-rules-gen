package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "falco-rules-gen",
	Short: "A CLI tool to generate Falco rules",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Aquí puedes inicializar flags o configuraciones globales
}
