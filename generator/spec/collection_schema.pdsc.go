package spec

type CollectionSchema struct {
	Identifier      *IdentifierSchema       `json:"identifier,omitempty"`
	AlternativeKeys []*AlternativeKeySchema `json:"alternativeKeys,omitempty"`
	Supports        []string                `json:"supports,omitempty"`
	Methods         []*RestMethodSchema     `json:"methods,omitempty"`
	Finders         []*FinderSchema         `json:"finders,omitempty"`
	Actions         []*ActionSchema         `json:"actions,omitempty"`
	Entity          *EntitySchema           `json:"entity,omitempty"`
}

