package cmd

import (
	"errors"
	"fmt"
	"github.com/fdaines/spm-go/common"
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/utils"
	"github.com/spf13/cobra"
	"go/build"
	"os"
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		utils.PrintMessage(err.Error())
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "spm",
	Version: "0.9.0",
	Short: "Software Package Metrics for Go",
	Long: "Software Package Metrics for Go",
}
var context = build.Default

func init() {
	rootCmd.PersistentFlags().StringVarP(&common.OutputFormat, "format", "f", "console", "Output format")
	rootCmd.PersistentFlags().BoolVarP(&common.Verbose, "verbose", "v", false, "Verbose Output")
}

func ValidateArgs(cmd *cobra.Command, args []string) error {
	err := validateOutputFormat(common.OutputFormat)
	return err
}

func validateOutputFormat(outputFormat string) error {
	supportedOutputFormats := map[string]bool{"csv": true, "console": true, "json": true}
	if !supportedOutputFormats[outputFormat] {
		return errors.New("output format should be one of 'plain', 'console' or 'json'")
	}
	return nil
}

func getBasicPackagesInfo() []*model.PackageInfo {
	var packagesInfo []*model.PackageInfo
	var context = build.Default

	pkgs, err := utils.GetPackages()
	if err != nil {
		utils.PrintMessage(fmt.Sprintf(
			"Error: %v\n",
			err,
		))
	} else {
		for index, packageName := range pkgs {
			utils.PrintVerboseMessage(fmt.Sprintf("Loading package (%d/%d): %s", index+1, len(pkgs), packageName))
			pkg, err := context.Import(packageName, "", 0)
			if err == nil {
				packagesInfo = append(packagesInfo, &model.PackageInfo{
					Name:         pkg.Name,
					Path:         pkg.ImportPath,
				})
			}
		}
	}

	return packagesInfo
}