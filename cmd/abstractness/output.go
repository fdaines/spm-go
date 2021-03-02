package abstractness

import (
	"encoding/json"
	"fmt"
	"github.com/fdaines/spm-go/model"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

func PrintPackages(packages []*model.PackageInfo, format string) {
	if format == "csv" {
		fmt.Printf("Package;Abstractions;Implementations;Abstractness\n")
		for _, p := range packages {
			fmt.Printf("%s;%d;%d;%.2f\n", p.Path, p.AbstractionsCount, p.ImplementationsCount, p.Abstractness)
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
		summary := &model.PackagesSummary{
			Packages: packages,
		}
		b, err := json.MarshalIndent(summary, "", "    ")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s\n", string(b))
	}
}
