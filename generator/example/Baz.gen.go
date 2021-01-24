package example

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (s *Baz) IntoSchemaMap() map[string]*schema.Schema {
	m := make(map[string]*schema.Schema)

	m["d"] = &schema.Schema{
		Type:     schema.TypeFloat,
		Optional: true,
	}

	return m
}

func (s *Baz) UnmarshalResource(d map[string]interface{}) {
	s.D = d["d"].(float32)
}

func (s *Baz) MarshalResource() map[string]interface{} {
	d := make(map[string]interface{})

	d["d"] = s.D

	return d
}
