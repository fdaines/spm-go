package packages

import (
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/utils/output"
)

func PrintPackages(packages []*model.PackageInfo, format string) {
	if format == "csv" {
		output.CsvFormatter(packages, []output.MetricOutput{output.FilesCount})
	} else if format == "console" {
		output.ConsoleFormatter(packages, []output.MetricOutput{output.FilesCount})
	} else if format == "json" {
		output.JsonFormatter(packages)
	}
}
