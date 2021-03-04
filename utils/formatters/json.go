package formatters

import (
	"encoding/json"
	"fmt"
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/utils"
)

func jsonFormatter(packages []*model.PackageInfo) {
	summary := &model.PackagesSummary{
		Packages: packages,
	}
	b, err := json.MarshalIndent(summary, "", "    ")
	if err != nil {
		utils.PrintMessage(fmt.Sprintf("%v\n", err))
	}
	utils.PrintMessage(fmt.Sprintf("%s\n", string(b)))
}
