package spec

type CustomAnnotationSchema struct {
	Annotations map[string]*CustomAnnotationContentSchema `json:"annotations,omitempty"`
}

