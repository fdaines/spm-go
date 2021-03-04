package abstractness

import (
	"fmt"
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/utils"
	"github.com/fdaines/spm-go/utils/output"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
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
		index := 1
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "Package", "Abstractions", "Implementations", "Abstractness"})

		for _, pkg := range packages {
			t.AppendRows([]table.Row{
				{
					index,
					pkg.Path,
					pkg.AbstractionsCount,
					pkg.ImplementationsCount,
					pkg.Abstractness,
				},
			})
			index = index + 1
		}
		t.Render()
	} else if format == "json" {
		output.JsonFormatter(packages)
	}
}
