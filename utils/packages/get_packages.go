package packages

import (
	"errors"
	"fmt"
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/utils"
	"go/build"
	"golang.org/x/tools/go/packages"
)

func GetBasicPackagesInfo() []*model.PackageInfo {
	var packagesInfo []*model.PackageInfo
	var context = build.Default

	pkgs, err := getPackages()
	if err != nil {
		utils.PrintMessage(fmt.Sprintf(
			"Error: %v\n",
			err,
		))
	} else {
		for index, packageName := range pkgs {
			utils.PrintVerboseMessage(fmt.Sprintf("Loading package (%d/%d): %s", index+1, len(pkgs), packageName))
			pkg, err := context.Import(packageName, "", 0)
			if err == nil {
				packagesInfo = append(packagesInfo, &model.PackageInfo{
					Name: pkg.Name,
					Path: pkg.ImportPath,
				})
			}
		}
	}

	return packagesInfo
}

func getPackages() ([]string, error) {
	utils.PrintVerboseMessage(fmt.Sprintf("Looking for packages."))
	cfg := &packages.Config{}
	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot load module packages: %+v", err))
	}
	var packages []string
	for _, pkg := range pkgs {
		packages = append(packages, pkg.PkgPath)
	}
	utils.PrintMessage(fmt.Sprintf("%d packages found...", len(packages)))
	return packages, nil
}
