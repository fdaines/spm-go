package formatters

import "github.com/fdaines/spm-go/model"

func FormatOutput(packages []*model.PackageInfo, format string, columns []MetricOutput) {
	switch format {
	case "console":
		consoleFormatter(packages, columns)
	case "csv":
		csvFormatter(packages, columns)
	case "json":
		jsonFormatter(packages)
	}
}
