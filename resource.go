package tfutils

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

// Data represents a read only data source
type Data interface {
	Read(d *schema.ResourceData, m interface{}) error
}

// CRUD represents a CRUD resource type
type CRUD interface {
	Read(d *schema.ResourceData, m interface{}) error
	Create(d *schema.ResourceData, m interface{}) error
	Update(d *schema.ResourceData, m interface{}) error
	Delete(d *schema.ResourceData, m interface{}) error
}

// SchemaMap is a builder for a *schema.Resource type
type SchemaMap map[string]Schema

// BuildCRUD creates a new CRUD *schema.Resource type
func (sm SchemaMap) BuildCRUD(crud CRUD) *schema.Resource {
	r := sm.BuildResource()
	r.Create = crud.Create
	r.Delete = crud.Delete
	r.Update = crud.Update
	r.Read = crud.Read
	return r
}

// BuildSchemaMap converts a SchemaMap into a map[string]*schema.Schema
func (sm SchemaMap) BuildSchemaMap() map[string]*schema.Schema {
	m := (map[string]Schema)(sm)
	s := make(map[string]*schema.Schema, len(m))
	for k, v := range m {
		s[k] = v.Build()
	}
	return s
}

// BuildResource converts a SchemaMap into a *schema.Resource
func (sm SchemaMap) BuildResource() *schema.Resource {
	return &schema.Resource{
		Schema: sm.BuildSchemaMap(),
	}
}

// BuildDataSource creates a new data source *schema.Resource type
func (sm SchemaMap) BuildDataSource(data Data) *schema.Resource {
	r := sm.BuildResource()
	r.Read = data.Read
	return r
}

// IntoSet converts the SchemaMap into a Set Schema over this structure
func (sm SchemaMap) IntoSet() SetSchema {
	return newComplexSetSchema(sm.BuildResource())
}

// IntoList converts the SchemaMap into a List Schema over this structure
func (sm SchemaMap) IntoList() ListSchema {
	return newComplexListSchema(sm.BuildResource())
}
