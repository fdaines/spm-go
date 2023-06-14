package packages

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fdaines/spm-go/common"
	"golang.org/x/mod/modfile"
)

const goModFile = "go.mod"

func GetMainPackage() (string, error) {
	if common.MainPackage != "main" {
		return common.MainPackage, nil
	}
	if _, err := os.Stat(goModFile); err == nil {
		content, _ := ioutil.ReadFile(goModFile)
		modulePath := modfile.ModulePath(content)
		fmt.Printf("Module: %s\n", modulePath)
		return modulePath, nil
	} else if os.IsNotExist(err) {
		fmt.Printf("Could not load %s file.\n", goModFile)
		return "", err
	} else {
		fmt.Printf("Could not load %s file.\n", goModFile)
		return "", err
	}
}
