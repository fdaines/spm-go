package utils

import (
	"errors"
	"fmt"
	"golang.org/x/tools/go/packages"
)

func GetPackages() ([]string, error) {
	cfg := &packages.Config{}
	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot load module packages: %+v", err))
	}
	var packages []string
	for _, pkg := range pkgs {
		packages = append(packages, pkg.PkgPath)
	}
	return packages, nil
}
