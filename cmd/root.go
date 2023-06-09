package cmd

import (
	"errors"
	"github.com/fdaines/spm-go/common"
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
	Use:     "spm-go",
	Version: common.Version,
	Short:   "Software Package Metrics for Go",
	Long:    "Software Package Metrics for Go",
}
var context = build.Default

func init() {
	rootCmd.PersistentFlags().StringVarP(&common.OutputFormat, "format", "f", "console", "Output format")
	rootCmd.PersistentFlags().StringVarP(&common.MainPackage, "main-package", "m", "main", "choose which package to use as mainPackage")
	rootCmd.PersistentFlags().BoolVarP(&common.Verbose, "verbose", "v", false, "Verbose Output")
	rootCmd.PersistentFlags().BoolVar(&common.HtmlOutput, "html", false, "Generate HTML Output")
}

func ValidateArgs(cmd *cobra.Command, args []string) error {
	err := validateOutputFormat(common.OutputFormat)
	return err
}

func validateOutputFormat(outputFormat string) error {
	supportedOutputFormats := map[string]bool{"csv": true, "console": true, "json": true}
	if !supportedOutputFormats[outputFormat] {
		return errors.New("formatters format should be one of 'plain', 'console' or 'json'")
	}
	return nil
}
