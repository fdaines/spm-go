package packages

import (
	"github.com/fdaines/spm-go/model"
	"go/build"
)

func FillFiles(packageInfo *model.PackageInfo, pkg *build.Package) *model.PackageInfo {
	packageInfo.Files = pkg.GoFiles
	packageInfo.FilesCount = len(pkg.GoFiles)

	return packageInfo
}
