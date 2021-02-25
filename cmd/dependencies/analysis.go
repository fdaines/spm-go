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
				internals, externals := filterDependencies(pkg.Imports, pkgs)
				depInfo := &model.DependenciesInfo{
					Externals: externals,
					Internals: internals,
					ExternalsCount: len(externals),
					InternalsCount: len(internals),
					TotalCount: len(pkg.Imports),
				}
				packagesInfo = append(packagesInfo,
					&model.PackageInfo{
						Name: pkg.Name,
						Path: pkg.ImportPath,
						Dependencies: depInfo,
					})
			}
		}
	}

	return packagesInfo
}

func filterDependencies(dependencies []string, internalPackages []string) ([]string, []string) {
	var internals []string
	var externals []string

	for _,dep := range dependencies {
		if utils.SliceContains(internalPackages, dep) {
			internals = append(internals, dep)
		} else {
			externals = append(externals, dep)
		}
	}

	return internals, externals
}