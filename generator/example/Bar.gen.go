package example

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (s *Bar) IntoSchemaMap() map[string]*schema.Schema {
	m := make(map[string]*schema.Schema)

	m["c"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}

	return m
}

func (s *Bar) UnmarshalResource(d map[string]interface{}) {
	s.C = d["c"].(string)
}

func (s *Bar) MarshalResource() map[string]interface{} {
	d := make(map[string]interface{})

	d["c"] = s.C

	return d
}
