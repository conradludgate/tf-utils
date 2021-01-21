package tfutils

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type SetSchema struct {
	s *schema.Schema
}

func newSetSchema(s *schema.Schema) SetSchema {
	s.Elem = &schema.Schema{Type: s.Type}
	s.Type = schema.TypeSet
	return SetSchema{s}
}

func newComplexSetSchema(r *schema.Resource) SetSchema {
	return SetSchema{&schema.Schema{
		Type: schema.TypeSet,
		Elem: r,
	}}
}

func (s SetSchema) Build() *schema.Schema {
	return s.s
}

func SetOf(sm Structure) SetSchema {
	return newComplexSetSchema(sm.Schema().BuildResource())
}

func (s SetSchema) MinItems(min int) SetSchema {
	s.s.MinItems = min
	return s
}

func (s SetSchema) MaxItems(max int) SetSchema {
	s.s.MaxItems = max
	return s
}

func (s SetSchema) SetFunc(f schema.SchemaSetFunc) SetSchema {
	s.s.Set = f
	return s
}
