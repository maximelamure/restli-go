package spec

type RestMethodSchema struct {
	Method          string                                    `json:"method,omitempty"`
	Doc             string                                    `json:"doc,omitempty"`
	Parameters      []*ParameterSchema                        `json:"parameters,omitempty"`
	PagingSupported bool                                      `json:"pagingSupported,omitempty"`
	Annotations     map[string]*CustomAnnotationContentSchema `json:"annotations,omitempty"`
}

