package cmd

import (
	"fmt"
	"github.com/fdaines/spm-go/cmd/dependencies"
	"github.com/fdaines/spm-go/utils"
	"github.com/fdaines/spm-go/utils/output"
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
		mainPackage, err := getMainPackage()
		if err != nil {
			fmt.Printf("Error: %+v\n", err)
			return
		}
		pkgsInfo := getBasicPackagesInfo()
		utils.PrintMessage("Gathering package metrics, please wait until the command is finished running.")
		for index, pkgInfo := range pkgsInfo {
			utils.PrintStep()
			pkg, err := context.Import(pkgInfo.Path, "", 0)
			if err == nil {
				pkgsInfo[index] = dependencies.FillDependencies(pkgsInfo[index], pkg, pkgsInfo)
			}
		}
		utils.PrintVerboseMessage("Done.")
		printDependencies(pkgsInfo)
		output.GenerateHtmlOutput(pkgsInfo, mainPackage)
	})
}
