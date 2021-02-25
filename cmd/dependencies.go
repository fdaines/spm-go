package cmd

import (
	"github.com/fdaines/spm-go/cmd/dependencies"
	"github.com/spf13/cobra"
)

var (
	dependenciesOutputFormat string
	dependenciesCmd = &cobra.Command{
		Use:   "dependencies",
		Short: "Lists dependencies of each package",
		Args:  validateDependenciesArgs,
		Run:   listPackagesDependencies,
	}
)

func init() {
	rootCmd.AddCommand(dependenciesCmd)
	dependenciesCmd.Flags().StringVarP(&dependenciesOutputFormat, "format", "f", "console", "Output format")
}

func listPackagesDependencies(cmd *cobra.Command, args []string) {
	pkgsInfo := dependencies.AnalyzePackages()
	dependencies.PrintPackages(pkgsInfo, dependenciesOutputFormat)
}

func validateDependenciesArgs(cmd *cobra.Command, args []string) error {
	err := dependencies.ValidateOutputFormat(dependenciesOutputFormat)
	return err
}
