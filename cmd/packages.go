package cmd

import (
	"github.com/fdaines/spm-go/cmd/packages"
	"github.com/spf13/cobra"
)

var (
	packagesCmd = &cobra.Command{
		Use:   "packages",
		Short: "Lists packages",
		Args:  ValidateArgs,
		Run:   listPackages,
	}
)

func init() {
	rootCmd.AddCommand(packagesCmd)
}

func listPackages(cmd *cobra.Command, args []string) {
	pkgsInfo := getBasicPackagesInfo()
	pkgsInfo = packages.AnalyzePackages(pkgsInfo)
	packages.PrintPackages(pkgsInfo, OutputFormat)
}
