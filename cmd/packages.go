package cmd

import (
	"errors"
	"fmt"
	"github.com/fdaines/spm-go/utils"
	"github.com/spf13/cobra"
)

var packagesCmd = &cobra.Command{
	Use:   "packages",
	Short: "Lists packages",
	Args: validateArgs,
	Run: listPackages,
}

func listPackages(cmd *cobra.Command, args []string) {
	packages, err := utils.GetPackages()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	printPackages(packages, outputFormat)
}

var outputFormat string
func init() {
	rootCmd.AddCommand(packagesCmd)
	packagesCmd.Flags().StringVarP(&outputFormat, "format", "f", "console", "Output format")
}

func validateArgs(cmd *cobra.Command, args []string) error {
	supportedOutputFormats := map[string]bool{"csv": true, "console": true, "json": true}
	if !supportedOutputFormats[outputFormat] {
		return errors.New("output format should be one of 'plain', 'console' or 'json'")
	}
	return nil
}

func printPackages(packages []string, format string) {
	if format == "csv" {
		fmt.Printf("Packages\n")
		for _, p := range packages {
			fmt.Printf("%s\n", p)
		}
	} else if format == "console" {
		fmt.Println("Output in 'console' format is not implemented.")
	} else if format == "json" {
		fmt.Println("Output in 'json' format is not implemented.")
	}
}
