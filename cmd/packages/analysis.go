package packages

import (
	"github.com/fdaines/spm-go/model"
	"go/build"
)

func AnalyzePackages(packagesInfo []*model.PackageInfo) []*model.PackageInfo {
	var context = build.Default
	for index, pkgInfo := range packagesInfo {
		pkg, err := context.Import(pkgInfo.Path, "", 0)
		if err == nil {
			packagesInfo[index] = fillFiles(pkgInfo, pkg)
		}
	}
	return packagesInfo
}

func fillFiles(packageInfo *model.PackageInfo, pkg *build.Package) *model.PackageInfo {
	packageInfo.Files = pkg.GoFiles
	packageInfo.FilesCount = len(pkg.GoFiles)

	return packageInfo
}
