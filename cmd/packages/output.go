package packages

import (
	"fmt"
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/utils"
	"github.com/fdaines/spm-go/utils/output"
)

func PrintPackages(packages []*model.PackageInfo, format string) {
	if format == "csv" {
		utils.PrintMessage("Package;Files")
		for _, p := range packages {
			utils.PrintMessage(fmt.Sprintf("%s;%d", p.Path, p.FilesCount))
		}
	} else if format == "console" {
		output.ConsoleFormatter(packages, []output.MetricOutput{output.FilesCount})
	} else if format == "json" {
		output.JsonFormatter(packages)
	}
}
