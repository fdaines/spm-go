package output

import (
	"github.com/fdaines/spm-go/common"
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/templates"
	"github.com/fdaines/spm-go/utils"
	"html/template"
	"os"
	"time"
)

func GenerateHtmlOutput(packages []*model.PackageInfo, module string, analysis string) {
	if !common.HtmlOutput {
		return
	}
	utils.PrintMessage("Creating html report into 'spm-go/output.html'......")
	checkOutputDirectory()
	summary := &htmlData{
		Version:   common.Version,
		Module:    module,
		TimeStamp: time.Now().Format("Mon Jan 02 2006 at 15:04:05"),
		Packages:  packages,
	}
	t, err := template.New("output").Parse(getHtmlTemplate(analysis))
	if err != nil {
		utils.PrintError("Error parsing output html template", err)
		return
	}

	f, err := os.Create("spm-go/output.html")
	if err != nil {
		utils.PrintError("Error creating html report", err)
		return
	}

	err = t.Execute(f, summary)
	if err != nil {
		utils.PrintError("Error creating html report", err)
		return
	}
	f.Close()
}

func getHtmlTemplate(analysis string) string {
	switch analysis {
	case "all":
		return templates.HtmlFullTemplate
	case "packages":
		return templates.HtmlPackagesTemplate
	case "dependencies":
		return templates.HtmlDependenciesTemplate
	case "instability":
		return templates.HtmlInstabilityTemplate
	case "abstractness":
		return templates.HtmlAbstractnessTemplate
	case "distance":
		return templates.HtmlDistanceTemplate
	}
	return templates.HtmlPackagesTemplate
}

func checkOutputDirectory() {
	if _, err := os.Stat("spm-go"); os.IsNotExist(err) {
		os.Mkdir("spm-go", 0755)
	}
}

type htmlData struct {
	Version   string
	Module    string
	TimeStamp string
	Packages  []*model.PackageInfo
}
