package utils

type Definition struct {
	Path              string `json:"path"`
	Name              string `json:"name"`
	Lang              string `json:"lang"`
	DependencyManager string `json:"dependencyManager"`
}
