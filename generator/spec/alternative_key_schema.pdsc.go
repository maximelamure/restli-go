package spec

type AlternativeKeySchema struct {
	Name       string `json:"name,omitempty"`
	Doc        string `json:"doc,omitempty"`
	Type       string `json:"type,omitempty"`
	KeyCoercer string `json:"keyCoercer,omitempty"`
}

