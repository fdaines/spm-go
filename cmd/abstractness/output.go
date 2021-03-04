package abstractness

import (
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/utils/output"
)

func PrintPackages(packages []*model.PackageInfo, format string) {
	outputColumns := []output.MetricOutput{
		output.Abstractions,
		output.Implementations,
		output.Abstractness,
	}
	if format == "csv" {
		output.ConsoleFormatter(packages, outputColumns)
	} else if format == "console" {
		outputColumns := []output.MetricOutput{
			output.Abstractions,
			output.Implementations,
			output.Abstractness,
		}
		output.ConsoleFormatter(packages, outputColumns)
	} else if format == "json" {
		output.JsonFormatter(packages)
	}
}
