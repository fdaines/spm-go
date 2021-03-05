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

func GenerateHtmlOutput(packages []*model.PackageInfo, module string, analysis string) {
	checkOutputDirectory()
	summary := &htmlData{
		Version: common.Version,
		Module: module,
		TimeStamp: time.Now().Format("Mon Jan 02 2006 at 15:04:05"),
		Packages: packages,
	}
	t, err := template.New("output").Parse(getHtmlTemplate(analysis))
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	f, err := os.Create("spm-go/output.html")
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	err = t.Execute(f, summary)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}
	f.Close()
}

func getHtmlTemplate(analysis string) string {
	switch analysis {
	case "all": return templates.HtmlFullTemplate
	case "packages": return templates.HtmlPackagesTemplate
	case "dependencies": return templates.HtmlDependenciesTemplate
	}
	return templates.HtmlFullTemplate
}

func checkOutputDirectory() {
	if _, err := os.Stat("spm-go"); os.IsNotExist(err) {
		os.Mkdir("spm-go", 0755)
	}
}

type htmlData struct {
	Version string
	Module string
	TimeStamp string
	Packages []*model.PackageInfo
}
