package abstractness

import (
	"fmt"
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/utils"
	"github.com/fdaines/spm-go/utils/output"
)

func PrintPackages(packages []*model.PackageInfo, format string) {
	if format == "csv" {
		utils.PrintMessage("Package;Abstractions;Implementations;Abstractness")
		for _, p := range packages {
			utils.PrintMessage(fmt.Sprintf("%s;%d;%d;%.2f",
				p.Path,
				p.AbstractionsCount,
				p.ImplementationsCount,
				p.Abstractness,
			))
		}
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
