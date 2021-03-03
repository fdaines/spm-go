package cmd

import (
	"fmt"
	"github.com/fdaines/spm-go/cmd/dependencies"
	"github.com/fdaines/spm-go/common"
	"github.com/fdaines/spm-go/utils"
	"github.com/spf13/cobra"
	"time"
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
	start := time.Now()
	utils.PrintMessage("Dependencies analysis started.")
	pkgsInfo := getBasicPackagesInfo()
	for index, pkgInfo := range pkgsInfo {
		pkg, err := context.Import(pkgInfo.Path, "", 0)
		if err == nil {
			pkgsInfo[index] = dependencies.FillDependencies(pkgsInfo[index], pkg, pkgsInfo)
		}
	}
	dependencies.PrintPackages(pkgsInfo, common.OutputFormat)
	elapsed := time.Since(start)

	utils.PrintMessage("Dependencies analysis finished.")
	utils.PrintMessage(fmt.Sprintf("Time: %.3f seconds", elapsed.Seconds()))
}
