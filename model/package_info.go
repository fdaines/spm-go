package model

type PackageInfo struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Files []string `json:"files,omitempty"`
	FilesCount int `json:"files_count,omitempty"`
	Dependencies *DependenciesInfo `json:"dependencies,omitempty"`
	Dependants []string `json:"dependants,omitempty"`
	AfferentCoupling int `json:"afferent_coupling,omitempty"`
	EfferentCoupling int `json:"efferent_coupling,omitempty"`
	Instability float64 `json:"instability,omitempty"`
}

type PackagesSummary struct {
	Packages []*PackageInfo `json:"packages"`
}

type DependenciesInfo struct {
	Standard []string `json:"standard,omitempty"`
	Internals []string `json:"internals,omitempty"`
	Externals []string `json:"externals,omitempty"`
	StandardCount int `json:"standard_count,omitempty"`
	InternalsCount int `json:"internals_count,omitempty"`
	ExternalsCount int `json:"externals_count,omitempty"`
	TotalCount int `json:"count,omitempty"`
}
