package spec

type ActionSchema struct {
	Name        string                                    `json:"name,omitempty"`
	Doc         string                                    `json:"doc,omitempty"`
	Parameters  []*ParameterSchema                        `json:"parameters,omitempty"`
	Returns     string                                    `json:"returns,omitempty"`
	Throws      []string                                  `json:"throws,omitempty"`
	Annotations map[string]*CustomAnnotationContentSchema `json:"annotations,omitempty"`
}
