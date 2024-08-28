package cmd

import (
	"fmt"
	"os"

	"github.com/diego-all/falco-rules-gen/generator"
	"github.com/spf13/cobra"
)

var outputFile string

var generateCmd = &cobra.Command{
	Use:   "generate [rule-name]",
	Short: "Generates a Falco rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ruleName := args[0]
		if outputFile == "" {
			fmt.Println("Output file must be specified using -o")
			return
		}

		// Genera la regla de Falco usando la función GenerateFalcoRule
		rule, err := generator.GenerateFalcoRule(ruleName)
		if err != nil {
			fmt.Printf("Error generating Falco rule: %v\n", err)
			return
		}

		// Guarda la regla en el archivo especificado
		file, err := os.Create(outputFile)
		if err != nil {
			fmt.Printf("Error creating output file: %v\n", err)
			return
		}
		defer file.Close()

		_, err = file.WriteString(rule)
		if err != nil {
			fmt.Printf("Error writing to output file: %v\n", err)
			return
		}

		fmt.Printf("Rule '%s' has been generated and saved to '%s'\n", ruleName, outputFile)
	},
}

func init() {
	generateCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file to save the generated rule")
	generateCmd.MarkFlagRequired("output")
}
