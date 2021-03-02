package cmd

import (
	"github.com/fdaines/spm-go/cmd/instability"
	"github.com/spf13/cobra"
)

var (
	instabilityCmd = &cobra.Command{
		Use:   "instability",
		Short: "Analyzes instability of packages",
		Args:  ValidateArgs,
		Run:   analyzeInstability,
	}
)

func init() {
	rootCmd.AddCommand(instabilityCmd)
}

func analyzeInstability(cmd *cobra.Command, args []string) {
	pkgsInfo := instability.AnalyzePackages()
	instability.PrintPackages(pkgsInfo, OutputFormat)
}
