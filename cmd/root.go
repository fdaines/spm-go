package cmd

import (
	"errors"
	"fmt"
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/utils"
	"github.com/spf13/cobra"
	"go/build"
	"os"
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "spm",
	Version: "1.0.0",
	Short: "Software Package Metrics for Go",
	Long: "Software Package Metrics for Go",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&OutputFormat, "format", "f", "console", "Output format")
}

func ValidateArgs(cmd *cobra.Command, args []string) error {
	err := validateOutputFormat(OutputFormat)
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
		fmt.Printf("Error: %v\n", err)
	} else {
		for _, packageName := range pkgs {
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