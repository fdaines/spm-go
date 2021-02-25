package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/utils"
	"github.com/spf13/cobra"
	"go/build"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

var packagesCmd = &cobra.Command{
	Use:   "packages",
	Short: "Lists packages",
	Args:  validateArgs,
	Run:   listPackages,
}

var outputFormat string
func init() {
	rootCmd.AddCommand(packagesCmd)
	packagesCmd.Flags().StringVarP(&outputFormat, "format", "f", "console", "Output format")
}

func listPackages(cmd *cobra.Command, args []string) {
	packages, err := utils.GetPackages()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	packagesInfo := analyzePackages(packages)
	printPackages(packagesInfo, outputFormat)
}

func validateArgs(cmd *cobra.Command, args []string) error {
	supportedOutputFormats := map[string]bool{"csv": true, "console": true, "json": true}
	if !supportedOutputFormats[outputFormat] {
		return errors.New("output format should be one of 'plain', 'console' or 'json'")
	}
	return nil
}

func analyzePackages(packages []string) []*model.PackageInfo {
	var packagesInfo []*model.PackageInfo
	var context = build.Default

	for _, packageName := range packages {
		pkg, err := context.Import(packageName, "", 0)
		if err == nil {
			packagesInfo = append(packagesInfo,
				&model.PackageInfo{
					Name: pkg.Name,
					Path: pkg.ImportPath,
					Files: pkg.GoFiles,
					FilesCount: len(pkg.GoFiles),
				})
		}
	}
	return packagesInfo
}

func printPackages(packages []*model.PackageInfo, format string) {
	if format == "csv" {
		fmt.Printf("Package;Files\n")
		for _, p := range packages {
			fmt.Printf("%s;%d\n", p.Path, p.FilesCount)
		}
	} else if format == "console" {
		index := 1
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "Package", "Files Count"})

		for _, pkg := range packages {
			t.AppendRows([]table.Row{
				{index, pkg.Path, pkg.FilesCount},
			})
			index = index + 1
		}
		t.Render()
	} else if format == "json" {
		summary := &model.PackagesSummary{
			Packages: packages,
		}
		b, err := json.MarshalIndent(summary, "", "    ")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s\n", string(b))
	}
}