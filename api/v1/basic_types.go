package v1

type Image struct {
	Repository string `json:"repository"`
	Tag        string `json:"tag,omitempty"`
}

type ResourceItem struct {
	Cpu    string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
}

type Resource struct {
	Request ResourceItem `json:"request,omitempty"`
	Limit   ResourceItem `json:"Limit,omitempty"`
}

type Container struct {
	Name     string   `json:"name"`
	Image    Image    `json:"image"`
	Command  []string `json:"command,omitempty"`
	Args     []string `json:"args,omitempty"`
	Resource Resource `json:"resource,omitempty"`
}
