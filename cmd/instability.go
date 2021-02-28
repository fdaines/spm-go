package cmd

import (
	"github.com/fdaines/spm-go/cmd/instability"
	"github.com/spf13/cobra"
)

var (
	instabilityOutputFormat string
	instabilityCmd = &cobra.Command{
		Use:   "instability",
		Short: "Analyzes instability of packages",
		Args:  validateInstabilityArgs,
		Run:   analyzeInstability,
	}
)

func init() {
	rootCmd.AddCommand(instabilityCmd)
	instabilityCmd.Flags().StringVarP(&outputFormat, "format", "f", "console", "Output format")
}

func analyzeInstability(cmd *cobra.Command, args []string) {
	pkgsInfo := instability.AnalyzePackages()
	instability.PrintPackages(pkgsInfo, outputFormat)
}

func validateInstabilityArgs(cmd *cobra.Command, args []string) error {
	err := instability.ValidateOutputFormat(outputFormat)
	return err
}
