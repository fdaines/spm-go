package cmd

import (
	"github.com/fdaines/spm-go/cmd/packages"
	"github.com/fdaines/spm-go/common"
	"github.com/fdaines/spm-go/utils"
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
		pkgsInfo := getBasicPackagesInfo()
		utils.PrintMessage("Gathering package metrics, please wait until the command is finished running.")
		for index, pkgInfo := range pkgsInfo {
			pkg, err := context.Import(pkgInfo.Path, "", 0)
			if err == nil {
				pkgsInfo[index] = packages.FillFiles(pkgInfo, pkg)
			}
		}
		packages.PrintPackages(pkgsInfo, common.OutputFormat)
	})
}
