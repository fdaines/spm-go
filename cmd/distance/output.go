package distance

import (
	"fmt"
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/utils"
	"github.com/fdaines/spm-go/utils/output"
)

func PrintPackages(packages []*model.PackageInfo, format string) {
	if format == "csv" {
		utils.PrintMessage("Package;Instability;Abstractness;Distance")
		for _, p := range packages {
			utils.PrintMessage(fmt.Sprintf("%s;%.2f;%.2f;%.2f", p.Path, p.Instability, p.Abstractness, p.Distance))
		}
	} else if format == "console" {
		outputColumns := []output.MetricOutput{
			output.Instability,
			output.Abstractness,
			output.Distance,
		}
		output.ConsoleFormatter(packages, outputColumns)
	} else if format == "json" {
		output.JsonFormatter(packages)
	}
}
