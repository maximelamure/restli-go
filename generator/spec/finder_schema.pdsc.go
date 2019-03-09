package spec

type FinderSchema struct {
	Name            string                                    `json:"name,omitempty"`
	Doc             string                                    `json:"doc,omitempty"`
	Parameters      []*ParameterSchema                        `json:"parameters,omitempty"`
	Metadata        *MetadataSchema                           `json:"metadata,omitempty"`
	AssocKey        string                                    `json:"assocKey,omitempty"`
	AssocKeys       []string                                  `json:"assocKeys,omitempty"`
	PagingSupported bool                                      `json:"pagingSupported,omitempty"`
	Annotations     map[string]*CustomAnnotationContentSchema `json:"annotations,omitempty"`
}
