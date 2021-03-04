package output

import (
	"github.com/fdaines/spm-go/model"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

func ConsoleFormatter(packages []*model.PackageInfo, columns []MetricOutput) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	header := table.Row{"#", "Package"}
	for _,column := range columns {
		header = append(header, column.Title)
	}
	t.AppendHeader(header)

	for index, pkg := range packages {
		row := table.Row{index + 1, pkg.Path}
		for _,column := range columns {
			row = append(row, column.Value(pkg))
		}
		t.AppendRow(row)
	}
	t.Render()
}
