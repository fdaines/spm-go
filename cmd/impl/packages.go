package impl

import "github.com/fdaines/spm-go/model"

func FillFiles(packageInfo *model.PackageInfo) *model.PackageInfo {
	packageInfo.Files = packageInfo.PackageData.GoFiles
	packageInfo.FilesCount = len(packageInfo.PackageData.GoFiles)

	return packageInfo
}