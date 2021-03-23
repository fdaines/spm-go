package dependencies

import (
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/utils/packages"
	"go/build"
)

func FillDependencies(packageInfo *model.PackageInfo, packagesInfo []*model.PackageInfo) *model.PackageInfo {
	var pkgs []string
	for _, current := range packagesInfo {
		pkgs = append(pkgs, current.Path)
	}
	packageInfo.Dependencies = resolveDependencies(packageInfo.PackageData, pkgs)
	return packageInfo
}

func resolveDependencies(pkg *build.Package, pkgs []string) *model.DependenciesInfo {
	internals, externals, standard := packages.FilterDependencies(pkg.Imports, pkgs)
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
