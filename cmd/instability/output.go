package instability

import (
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/utils/output"
)

func PrintPackages(packages []*model.PackageInfo, format string) {
	outputColumns := []output.MetricOutput{
		output.AfferentCoupling,
		output.EfferentCoupling,
		output.Instability,
	}
	if format == "csv" {
		output.CsvFormatter(packages, outputColumns)
	} else if format == "console" {
		output.ConsoleFormatter(packages, outputColumns)
	} else if format == "json" {
		output.JsonFormatter(packages)
	}
}
