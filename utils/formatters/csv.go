package formatters

import (
	"fmt"
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/utils"
	"strings"
)

func csvFormatter(packages []*model.PackageInfo, columns []MetricOutput) {
	header := []string{"Package"}
	for _, c := range columns {
		header = append(header, c.Title)
	}
	utils.PrintMessage(strings.Join(header, ";"))

	for _, p := range packages {
		values := []string{p.Path}
		for _, c := range columns {
			values = append(values, fmt.Sprintf("%v", c.Value(p)))
		}
		utils.PrintMessage(strings.Join(values, ";"))
	}
}
