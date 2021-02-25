package packages

import (
	"encoding/json"
	"fmt"
	"github.com/fdaines/spm-go/model"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

func PrintPackages(packages []*model.PackageInfo, format string) {
	if format == "csv" {
		fmt.Printf("Package;Files\n")
		for _, p := range packages {
			fmt.Printf("%s;%d\n", p.Path, p.FilesCount)
		}
	} else if format == "console" {
		index := 1
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "Package", "Files Count"})

		for _, pkg := range packages {
			t.AppendRows([]table.Row{
				{index, pkg.Path, pkg.FilesCount},
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
