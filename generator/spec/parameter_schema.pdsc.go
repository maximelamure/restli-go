package spec

type ParameterSchema struct {
	Name        string                                    `json:"name,omitempty"`
	Type        string                                    `json:"type,omitempty"`
	Items       string                                    `json:"items,omitempty"`
	Optional    bool                                      `json:"optional,omitempty"`
	Default     string                                    `json:"default,omitempty"`
	Doc         string                                    `json:"doc,omitempty"`
	Annotations map[string]*CustomAnnotationContentSchema `json:"annotations,omitempty"`
}

