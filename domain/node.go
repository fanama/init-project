package domain

type NodePackage struct {
	Name            string      `json:"name"`
	Version         string      `json:"version"`
	Script          NodeScript  `json:"script"`
	Dependencies    interface{} `json:"dependencies"`
	DevDependencies interface{} `json:"devDependencies"`
}

type NodeScript struct {
	Start   string `json:"start"`
	Dev     string `json:"dev"`
	Build   string `json:"build"`
	Lint    string `json:"lint"`
	Prepare string `json:"prepare"`
}
