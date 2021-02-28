package utils

func FilterDependencies(dependencies []string, internalPackages []string) ([]string, []string) {
	var internals []string
	var externals []string

	for _,dep := range dependencies {
		if SliceContains(internalPackages, dep) {
			internals = append(internals, dep)
		} else {
			externals = append(externals, dep)
		}
	}

	return internals, externals
}
