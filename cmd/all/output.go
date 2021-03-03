package all

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
		index := 1
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{
			"#",
			"Package",
			"Files",
			"Afferent",
			"Efferent",
			"Abstractions",
			"Implementations",
			"Instability",
			"Abstractness",
			"Distance",
		})

		for _, pkg := range packages {
			t.AppendRows([]table.Row{
				{
					index,
					pkg.Path,
					pkg.FilesCount,
					pkg.AfferentCoupling,
					pkg.EfferentCoupling,
					pkg.AbstractionsCount,
					pkg.ImplementationsCount,
					pkg.Instability,
					pkg.Abstractness,
					pkg.Distance,
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
