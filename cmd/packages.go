package cmd

import (
	"github.com/fdaines/spm-go/cmd/packages"
	"github.com/spf13/cobra"
)

var packagesCmd = &cobra.Command{
	Use:   "packages",
	Short: "Lists packages",
	Args:  validateArgs,
	Run:   listPackages,
}

var outputFormat string
func init() {
	rootCmd.AddCommand(packagesCmd)
	packagesCmd.Flags().StringVarP(&outputFormat, "format", "f", "console", "Output format")
}

func listPackages(cmd *cobra.Command, args []string) {
	pkgsInfo := packages.AnalyzePackages()
	packages.PrintPackages(pkgsInfo, outputFormat)
}

func validateArgs(cmd *cobra.Command, args []string) error {
	err := packages.ValidateOutputFormat(outputFormat)
	return err
}
