package spec

import "github.com/maximelamure/restli-go/common"

type ResourceSchema struct {
	Name        string                                    `json:"name,omitempty"`
	Namespace   string                                    `json:"namespace,omitempty"`
	Path        string                                    `json:"path,omitempty"`
	Schema      string                                    `json:"schema,omitempty"`
	Doc         string                                    `json:"doc,omitempty"`
	Collection  *CollectionSchema                         `json:"collection,omitempty"`
	Association *AssociationSchema                        `json:"association,omitempty"`
	ActionsSet  *ActionsSetSchema                         `json:"actionsSet,omitempty"`
	Simple      *SimpleSchema                             `json:"simple,omitempty"`
	Annotations map[string]*CustomAnnotationContentSchema `json:"annotations,omitempty"`
}

func (s *ResourceSchema) GetFileName(pre string) string {
	name := common.ToSnake(s.Name) + ".go"
	if len(pre) == 0 {
		return name
	}

	return pre + "_" + name
}
