package cmd

import (
	"fmt"
	"github.com/fdaines/spm-go/cmd/dependencies"
	"github.com/fdaines/spm-go/utils"
	"github.com/fdaines/spm-go/utils/output"
	pkg "github.com/fdaines/spm-go/utils/packages"
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
	utils.ExecuteWithTimer(func() {
		utils.PrintMessage("Dependencies analysis started.")
		mainPackage, err := pkg.GetMainPackage()
		if err != nil {
			fmt.Printf("Error: %+v\n", err)
			return
		}
		pkgsInfo := pkg.GetBasicPackagesInfo()
		utils.PrintMessage("Gathering package metrics, please wait until the command is finished running.")
		for index, pkgInfo := range pkgsInfo {
			utils.PrintStep()
			pkgsInfo[index] = dependencies.FillDependencies(pkgInfo, pkgsInfo)
		}
		utils.PrintVerboseMessage("Done.")
		printDependencies(pkgsInfo)
		output.GenerateHtmlOutput(pkgsInfo, mainPackage, "dependencies")
	})
}
