package instability

import (
	"fmt"
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/utils"
	"github.com/fdaines/spm-go/utils/output"
)

func PrintPackages(packages []*model.PackageInfo, format string) {
	if format == "csv" {
		utils.PrintMessage("Package;Afferent;Efferent;Instability")
		for _, p := range packages {
			utils.PrintMessage(fmt.Sprintf("%s;%d;%d;%.2f", p.Path, p.AfferentCoupling, p.EfferentCoupling, p.Instability))
		}
	} else if format == "console" {
		outputColumns := []output.MetricOutput{
			output.AfferentCoupling,
			output.EfferentCoupling,
			output.Instability,
		}
		output.ConsoleFormatter(packages, outputColumns)
	} else if format == "json" {
		output.JsonFormatter(packages)
	}
}
