package model

type PackageInfo struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Files []string `json:"files"`
	FilesCount int `json:"files_count"`
	Dependencies *DependenciesInfo `json:"dependencies"`
}

type PackagesSummary struct {
	Packages []*PackageInfo `json:"packages"`
}

type DependenciesInfo struct {
	Internals []string `json:"internals"`
	Externals []string `json:"externals"`
	InternalsCount int `json:"internals_count"`
	ExternalsCount int `json:"externals_count"`
	TotalCount int `json:"count"`
}
