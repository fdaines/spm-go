package cmd

import (
	"fmt"
	"github.com/fdaines/spm-go/cmd/packages"
	"github.com/fdaines/spm-go/utils"
	"github.com/fdaines/spm-go/utils/output"
	pkg "github.com/fdaines/spm-go/utils/packages"
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
	utils.ExecuteWithTimer(func() {
		utils.PrintMessage("Packages analysis started.")
		mainPackage, err := pkg.GetMainPackage()
		if err != nil {
			fmt.Printf("Error: %+v\n", err)
			return
		}
		pkgsInfo := pkg.GetBasicPackagesInfo()
		utils.PrintMessage("Gathering package metrics, please wait until the command is finished running.")
		for index, pkgInfo := range pkgsInfo {
			utils.PrintStep()
			pkg, err := context.Import(pkgInfo.Path, "", 0)
			if err == nil {
				pkgsInfo[index] = packages.FillFiles(pkgInfo, pkg)
			}
		}
		utils.PrintVerboseMessage("Done.")
		printPackages(pkgsInfo)
		output.GenerateHtmlOutput(pkgsInfo, mainPackage, "packages")
	})
}
