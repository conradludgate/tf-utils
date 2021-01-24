package tfutils

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Data represents a read only data source
type Data interface {
	Schema
	SetClient(m interface{})

	Read() error
}

// CRUD represents a CRUD resource type
type CRUD interface {
	Schema
	SetClient(m interface{})

	Read() error
	Create() error
	Update(old map[string]interface{}) error
	Delete() error
}

func save(d *schema.ResourceData, old, new map[string]interface{}) {
	for key, v := range new {
		if !reflect.DeepEqual(old[key], v) {
			d.Set(key, v)
		}
	}
}
func load(d *schema.ResourceData, sm map[string]*schema.Schema) map[string]interface{} {
	r := make(map[string]interface{}, len(sm))
	for key := range sm {
		r[key] = d.Get(key)
	}
	r["id"] = d.Id()
	return r
}
func loadOld(d *schema.ResourceData, sm map[string]*schema.Schema) (old, new map[string]interface{}) {
	old = make(map[string]interface{}, len(sm))
	old = make(map[string]interface{}, len(sm))
	for key := range sm {
		old[key], new[key] = d.GetChange(key)
	}
	old["id"] = d.Id()
	new["id"] = d.Id()
	return
}

// BuildData creates a new data *schema.Resource type
func BuildData(data Data) *schema.Resource {
	sm := data.IntoSchemaMap()
	return &schema.Resource{
		Schema: sm,

		Read: func(d *schema.ResourceData, m interface{}) error {
			data.SetClient(m)

			r := load(d, sm)

			data.UnmarshalResource(r)

			if err := data.Read(); err != nil {
				return err
			}

			save(d, r, data.MarshalResource())

			return nil
		},
	}
}

// BuildCRUD creates a new CRUD *schema.Resource type
func BuildCRUD(crud CRUD) *schema.Resource {
	sm := crud.IntoSchemaMap()
	return &schema.Resource{
		Schema: sm,

		Create: func(d *schema.ResourceData, m interface{}) error {
			crud.SetClient(m)

			r := load(d, sm)

			crud.UnmarshalResource(r)

			if err := crud.Create(); err != nil {
				return err
			}

			save(d, r, crud.MarshalResource())

			return nil
		},

		Read: func(d *schema.ResourceData, m interface{}) error {
			crud.SetClient(m)

			r := load(d, sm)

			crud.UnmarshalResource(r)

			if err := crud.Read(); err != nil {
				return err
			}

			save(d, r, crud.MarshalResource())

			return nil
		},

		Update: func(d *schema.ResourceData, m interface{}) error {
			crud.SetClient(m)

			old, new := loadOld(d, sm)

			crud.UnmarshalResource(new)

			if err := crud.Update(old); err != nil {
				return err
			}

			save(d, new, crud.MarshalResource())

			return nil
		},

		Delete: func(d *schema.ResourceData, m interface{}) error {
			crud.SetClient(m)

			r := load(d, sm)

			crud.UnmarshalResource(r)
			return crud.Delete()
		},
	}
}

type Schema interface {
	UnmarshalResource(map[string]interface{})
	MarshalResource() map[string]interface{}
	IntoSchemaMap() map[string]*schema.Schema
}

type Set interface {
	Schema
	Hash() int
}

func BuildHashFunc(s Set) schema.SchemaSetFunc {
	return func(d interface{}) int {
		s.UnmarshalResource(d.(map[string]interface{}))
		return s.Hash()
	}
}

func GetHashFunc(s Schema) schema.SchemaSetFunc {
	var i interface{} = s
	if set, ok := i.(Set); ok {
		return BuildHashFunc(set)
	}
	return schema.HashResource(&schema.Resource{Schema: s.IntoSchemaMap()})
}
