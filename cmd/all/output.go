package all

import (
	"fmt"
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/utils"
	"github.com/fdaines/spm-go/utils/output"
)

func PrintPackages(packages []*model.PackageInfo, format string) {
	if format == "csv" {
		utils.PrintMessage("Package;Files;Afferent;Efferent;Abstractions;Implementations;Instability;Abstractness;Distance")
		for _, p := range packages {
			utils.PrintMessage(fmt.Sprintf(
				"%s;%d;%d;%d;%.2f;%.2f;%.2f\n",
				p.Path,
				p.FilesCount,
				p.AfferentCoupling,
				p.EfferentCoupling,
				p.Instability,
				p.Abstractness,
				p.Distance,
			))
		}
	} else if format == "console" {
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
		output.ConsoleFormatter(packages, outputColumns)
	} else if format == "json" {
		output.JsonFormatter(packages)
	}
}
