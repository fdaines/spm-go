package model

type PackageInfo struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Files []string `json:"files"`
	FilesCount int `json:"files_count"`
}

type PackagesSummary struct {
	Packages []*PackageInfo `json:"packages"`
}
