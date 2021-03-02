package cmd

import (
	"github.com/fdaines/spm-go/cmd/dependencies"
	"github.com/spf13/cobra"
)

var (
	dependenciesCmd = &cobra.Command{
		Use:   "dependencies",
		Short: "Lists dependencies of each package",
		Args:  ValidateArgs,
		Run:   listPackagesDependencies,
	}
)

func init() {
	rootCmd.AddCommand(dependenciesCmd)
}

func listPackagesDependencies(cmd *cobra.Command, args []string) {
	pkgsInfo := getBasicPackagesInfo()
	for index, pkgInfo := range pkgsInfo {
		pkg, err := context.Import(pkgInfo.Path, "", 0)
		if err == nil {
			pkgsInfo[index] = dependencies.FillDependencies(pkgsInfo[index], pkg, pkgsInfo)
		}
	}
	dependencies.PrintPackages(pkgsInfo, OutputFormat)
}
