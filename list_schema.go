package tfutils

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type ListSchema struct {
	s *schema.Schema
}

func newListSchema(s *schema.Schema) ListSchema {
	s.Elem = &schema.Schema{Type: s.Type}
	s.Type = schema.TypeList
	return ListSchema{s}
}

func newComplexListSchema(r *schema.Resource) ListSchema {
	return ListSchema{&schema.Schema{
		Type: schema.TypeList,
		Elem: r,
	}}
}

func (s ListSchema) Build() *schema.Schema {
	return s.s
}

func ListOf(sm Structure) ListSchema {
	return newComplexListSchema(sm.Schema().BuildResource())
}
