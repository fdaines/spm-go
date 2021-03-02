package cmd

import (
	"github.com/fdaines/spm-go/cmd/dependencies"
	"github.com/spf13/cobra"
)

var (
	dependenciesCmd = &cobra.Command{
		Use:   "dependencies",
		Short: "Lists dependencies of each package",
		Args:  validateDependenciesArgs,
		Run:   listPackagesDependencies,
	}
)

func init() {
	rootCmd.AddCommand(dependenciesCmd)
}

func listPackagesDependencies(cmd *cobra.Command, args []string) {
	pkgsInfo := dependencies.AnalyzePackages()
	dependencies.PrintPackages(pkgsInfo, OutputFormat)
}

func validateDependenciesArgs(cmd *cobra.Command, args []string) error {
	err := dependencies.ValidateOutputFormat(OutputFormat)
	return err
}
