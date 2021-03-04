package all

import (
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/utils/output"
)

func PrintPackages(packages []*model.PackageInfo, format string) {
	outputColumns := []output.MetricOutput{
		output.FilesCount,
		output.AfferentCoupling,
		output.EfferentCoupling,
		output.Abstractions,
		output.Implementations,
		output.Instability,
		output.Abstractness,
		output.Distance,
	}
	if format == "csv" {
		output.CsvFormatter(packages, outputColumns)
	} else if format == "console" {
		output.ConsoleFormatter(packages, outputColumns)
	} else if format == "json" {
		output.JsonFormatter(packages)
	}
}
