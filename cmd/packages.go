package cmd

import (
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/utils"
	"github.com/fdaines/spm-go/utils/output"
	pkg "github.com/fdaines/spm-go/utils/packages"
	"github.com/spf13/cobra"
	"go/build"
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
			utils.PrintError("Error loading main package", err)
			return
		}
		pkgsInfo := pkg.GetBasicPackagesInfo()
		utils.PrintMessage("Gathering package metrics, please wait until the command is finished running.")
		for index, pkgInfo := range pkgsInfo {
			utils.PrintStep()
			pkg, err := context.Import(pkgInfo.Path, "", 0)
			if err == nil {
				pkgsInfo[index] = fillFiles(pkgInfo, pkg)
			}
		}
		utils.PrintVerboseMessage("Done.")
		printPackages(pkgsInfo)
		output.GenerateHtmlOutput(pkgsInfo, mainPackage, "packages")
	})
}

func fillFiles(packageInfo *model.PackageInfo, pkg *build.Package) *model.PackageInfo {
	packageInfo.Files = pkg.GoFiles
	packageInfo.FilesCount = len(pkg.GoFiles)

	return packageInfo
}