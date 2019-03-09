package spec

type EntitySchema struct {
	Path         string            `json:"path,omitempty"`
	Actions      []*ActionSchema   `json:"actions,omitempty"`
	Subresources []*ResourceSchema `json:"subresources,omitempty"`
}

