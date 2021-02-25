package model

type PackageInfo struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Files []string `json:"files"`
	FilesNumber int `json:"files_number"`
}

type PackagesSummary struct {
	Packages []PackageInfo `json:packages`
}
