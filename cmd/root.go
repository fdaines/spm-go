package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
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
