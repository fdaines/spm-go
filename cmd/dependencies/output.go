package dependencies

import (
	"fmt"
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/utils"
	"github.com/fdaines/spm-go/utils/output"
)

func PrintPackages(packages []*model.PackageInfo, format string) {
	if format == "csv" {
		utils.PrintMessage("Package;Internals;Externals;Standard;Total")
		for _, pkg := range packages {
			utils.PrintMessage(fmt.Sprintf("%s;%d;%d;%d;%d",
				pkg.Path,
				pkg.Dependencies.InternalsCount,
				pkg.Dependencies.ExternalsCount,
				pkg.Dependencies.StandardCount,
				pkg.Dependencies.TotalCount))
		}
	} else if format == "console" {
		outputColumns := []output.MetricOutput{
			output.InternalDependencies,
			output.ExternalDependencies,
			output.StandardDependencies,
			output.TotalDependencies,
		}
		output.ConsoleFormatter(packages, outputColumns)
	} else if format == "json" {
		output.JsonFormatter(packages)
	}
}
