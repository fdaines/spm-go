package utils

import (
	"errors"
	"fmt"
	"golang.org/x/tools/go/packages"
)

func GetPackages() ([]string, error) {
	PrintVerboseMessage(fmt.Sprintf("Looking for packages."))
	cfg := &packages.Config{}
	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot load module packages: %+v", err))
	}
	var packages []string
	for _, pkg := range pkgs {
		packages = append(packages, pkg.PkgPath)
	}
	PrintVerboseMessage(fmt.Sprintf("Should load data from %d packages...", len(packages)))
	return packages, nil
}
