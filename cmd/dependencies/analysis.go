package dependencies

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
				packageInfo := &model.PackageInfo{
					Name:         pkg.Name,
					Path:         pkg.ImportPath,
				}
				packageInfo.Dependencies = resolveDependencies(pkg, pkgs)
				packagesInfo = append(packagesInfo, packageInfo)
			}
		}
	}

	return packagesInfo
}

func resolveDependencies(pkg *build.Package, pkgs []string) *model.DependenciesInfo {
	internals, externals, standard := utils.FilterDependencies(pkg.Imports, pkgs)
	depInfo := &model.DependenciesInfo{
		Standard:       standard,
		Externals:      externals,
		Internals:      internals,
		StandardCount:  len(standard),
		ExternalsCount: len(externals),
		InternalsCount: len(internals),
		TotalCount:     len(pkg.Imports),
	}
	return depInfo
}