package packages

import (
	"fmt"
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/utils"
	"go/build"
)

func AnalyzePackages() []*model.PackageInfo {
	var packagesInfo []*model.PackageInfo
	var context = build.Default

	pkgs, err := utils.GetPackages()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		for _, packageName := range pkgs {
			pkg, err := context.Import(packageName, "", 0)
			if err == nil {
				packagesInfo = append(packagesInfo,
					&model.PackageInfo{
						Name: pkg.Name,
						Path: pkg.ImportPath,
						Files: pkg.GoFiles,
						FilesCount: len(pkg.GoFiles),
					})
			}
		}
	}

	return packagesInfo
}
