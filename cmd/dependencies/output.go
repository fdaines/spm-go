package dependencies

import (
	"encoding/json"
	"fmt"
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/utils"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
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
		index := 1
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "Package", "Internals", "Externals", "Standard", "Total"})

		for _, pkg := range packages {
			t.AppendRows([]table.Row{
				{
					index,
					pkg.Path,
					pkg.Dependencies.InternalsCount,
					pkg.Dependencies.ExternalsCount,
					pkg.Dependencies.StandardCount,
					pkg.Dependencies.TotalCount,
				},
			})
			index = index + 1
		}
		t.Render()
	} else if format == "json" {
		summary := &model.PackagesSummary{
			Packages: packages,
		}
		b, err := json.MarshalIndent(summary, "", "    ")
		if err != nil {
			utils.PrintMessage(fmt.Sprintf("%v\n", err))
		}
		utils.PrintMessage(fmt.Sprintf("%s\n", string(b)))
	}
}
