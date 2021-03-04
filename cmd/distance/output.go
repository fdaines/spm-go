package distance

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
		utils.PrintMessage("Package;Instability;Abstractness;Distance")
		for _, p := range packages {
			utils.PrintMessage(fmt.Sprintf("%s;%.2f;%.2f;%.2f", p.Path, p.Instability, p.Abstractness, p.Distance))
		}
	} else if format == "console" {
		index := 1
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "Package", "Instability", "Abstractness", "Distance"})

		for _, pkg := range packages {
			t.AppendRows([]table.Row{
				{
					index,
					pkg.Path,
					pkg.Instability,
					pkg.Abstractness,
					pkg.Distance,
				},
			})
			index = index + 1
		}
		t.Render()
	} else if format == "json" {
		output.JsonFormatter(packages)
	}
}
