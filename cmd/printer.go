package cmd

import (
	"github.com/fdaines/spm-go/common"
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/utils/formatters"
)

func printPackages(pkgs []*model.PackageInfo) {
	outputColumns := []formatters.MetricOutput{formatters.FilesCount}
	formatters.FormatOutput(pkgs, common.OutputFormat, outputColumns)
}

func printDependencies(pkgs []*model.PackageInfo) {
	outputColumns := []formatters.MetricOutput{
		formatters.InternalDependencies,
		formatters.ExternalDependencies,
		formatters.StandardDependencies,
		formatters.TotalDependencies,
	}
	formatters.FormatOutput(pkgs, common.OutputFormat, outputColumns)
}

func printInstability(pkgs []*model.PackageInfo) {
	outputColumns := []formatters.MetricOutput{
		formatters.AfferentCoupling,
		formatters.EfferentCoupling,
		formatters.Instability,
	}
	formatters.FormatOutput(pkgs, common.OutputFormat, outputColumns)
}

func printAbstractness(pkgs []*model.PackageInfo) {
	outputColumns := []formatters.MetricOutput{
		formatters.Abstractions,
		formatters.Implementations,
		formatters.Abstractness,
	}
	formatters.FormatOutput(pkgs, common.OutputFormat, outputColumns)
}

func printDistance(pkgs []*model.PackageInfo) {
	outputColumns := []formatters.MetricOutput{
		formatters.Instability,
		formatters.Abstractness,
		formatters.Distance,
	}
	formatters.FormatOutput(pkgs, common.OutputFormat, outputColumns)
}

func printAll(pkgs []*model.PackageInfo) {
	outputColumns := []formatters.MetricOutput{
		formatters.FilesCount,
		formatters.AfferentCoupling,
		formatters.EfferentCoupling,
		formatters.Abstractions,
		formatters.Implementations,
		formatters.Instability,
		formatters.Abstractness,
		formatters.Distance,
	}
	formatters.FormatOutput(pkgs, common.OutputFormat, outputColumns)
}
