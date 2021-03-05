package output

import (
	"fmt"
	"github.com/fdaines/spm-go/common"
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/templates"
	"html/template"
	"os"
	"time"
)

func GenerateHtmlOutput(packages []*model.PackageInfo, module string) {
	summary := &htmlData{
		Version: common.Version,
		Module: module,
		TimeStamp: time.Now().Format("Mon Jan 02 2006 at 15:04:05"),
		Packages: packages,
	}
	t, err := template.New("output").Parse(templates.HtmlTemplate)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	f, err := os.Create("output.html")
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	err = t.Execute(f, summary)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}
	f.Close()
}

type htmlData struct {
	Version string
	Module string
	TimeStamp string
	Packages []*model.PackageInfo
}
