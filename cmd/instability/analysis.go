package instability

import (
	"fmt"
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/utils"
	"go/build"
)

func AnalyzePackages() []*model.PackageInfo {
	var packagesInfo []*model.PackageInfo
	var context = build.Default
	var afferentMap = make(map[string][]string)

	pkgs, err := utils.GetPackages()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		for _, packageName := range pkgs {
			pkg, err := context.Import(packageName, "", 0)
			if err == nil {
				internals, externals, _ := utils.FilterDependencies(pkg.Imports, pkgs)
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
				for _,item := range internals {
					afferentMap[item] = append(afferentMap[item], pkg.ImportPath)
				}
			}
		}

		for index, pkgInfo := range packagesInfo {
			packagesInfo[index].Dependants = afferentMap[pkgInfo.Path]
			packagesInfo[index].AfferentCoupling = len(packagesInfo[index].Dependants)
			packagesInfo[index].EfferentCoupling = packagesInfo[index].Dependencies.InternalsCount
			packagesInfo[index].Instability = utils.RoundValue(float64(packagesInfo[index].EfferentCoupling) / float64(packagesInfo[index].EfferentCoupling + packagesInfo[index].AfferentCoupling))
		}
	}

	return packagesInfo
}
