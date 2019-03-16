package spec

import "github.com/maximelamure/restli-go/common"

type Schema struct {
	Type      string   `json:"type"`
	Name      string   `json:"name"`
	Namespace string   `json:"namespace"`
	Package   string   `json:"package"`
	Doc       string   `json:"doc"`
	Fields    []*Field `json:"fields"`
}

type Field struct {
	Name     string `json:"name"`
	Doc      string `json:"doc"`
	Type     string `json:"type"`
	Optional bool   `json:"optional,omitempty"`
}

func (s *Schema) GetFileName() string {
	return common.ToSnake(s.Name) + ".go"
}
