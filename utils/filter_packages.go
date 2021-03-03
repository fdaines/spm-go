package utils

import "strings"

func FilterDependencies(dependencies []string, internalPackages []string) ([]string, []string, []string) {
	internals := []string{}
	externals := []string{}
	standard := []string{}

	for _, dep := range dependencies {
		if SliceContains(internalPackages, dep) {
			internals = append(internals, dep)
		} else {
			if strings.Contains(dep, ".") {
				externals = append(externals, dep)
			} else {
				standard = append(standard, dep)
			}
		}
	}

	return internals, externals, standard
}
