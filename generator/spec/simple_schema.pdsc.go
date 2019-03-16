package spec

type SimpleSchema struct {
	Supports []string            `json:"supports,omitempty"`
	Methods  []*RestMethodSchema `json:"methods,omitempty"`
	Actions  []*ActionSchema     `json:"actions,omitempty"`
	Entity   *EntitySchema       `json:"entity,omitempty"`
}

