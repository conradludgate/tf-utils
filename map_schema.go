package tfutils

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type MapSchema struct {
	s *schema.Schema
}

func newMapSchema(s *schema.Schema) MapSchema {
	s.Elem = &schema.Schema{Type: s.Type}
	s.Type = schema.TypeMap
	return MapSchema{s}
}

func (s MapSchema) Build() *schema.Schema {
	return s.s
}
