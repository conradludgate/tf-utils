package example

import (
	"github.com/conradludgate/tfutils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (s *Foo) IntoSchemaMap() map[string]*schema.Schema {
	m := make(map[string]*schema.Schema)

	m["a"] = &schema.Schema{
		Type:        schema.TypeString,
		Default:     "foo",
		Optional:    true,
		Description: "Some value",
	}

	m["b"] = &schema.Schema{
		Type:        schema.TypeInt,
		Required:    true,
		Description: "Some other value",
	}

	m["list"] = &schema.Schema{
		Type:        schema.TypeList,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Optional:    true,
		Description: "Simple lists can be implemented",
	}

	anotherListElem := Baz{}
	m["another_list"] = &schema.Schema{
		Type:        schema.TypeList,
		Elem:        &schema.Resource{Schema: anotherListElem.IntoSchemaMap()},
		Optional:    true,
		Description: "Complex lists can too, as long as Baz\nimplements the 'Schema' interface",
	}

	m["map"] = &schema.Schema{
		Type:        schema.TypeMap,
		Elem:        &schema.Schema{Type: schema.TypeInt},
		Optional:    true,
		Description: "Maps can only be over simple types (terraform limitation)",
	}

	setElem := Bar{}
	m["set"] = &schema.Schema{
		Type:        schema.TypeSet,
		Elem:        &schema.Resource{Schema: setElem.IntoSchemaMap()},
		Set:         tfutils.GetHashFunc(&setElem),
		Optional:    true,
		Description: "map[int]... represents a Set.\nIf Bar implements the `Set` interface,\nthen that will be the Set function",
	}

	return m
}

func (s *Foo) UnmarshalResource(d map[string]interface{}) {
	s.A = d["a"].(string)

	s.B = d["b"].(int)

	listElems := d["list"].([]interface{})
	s.List = make([]string, len(listElems))
	for i, v := range listElems {
		s.List[i] = v.(string)
	}

	anotherListElems := d["another_list"].([]interface{})
	s.AnotherList = make([]Baz, len(anotherListElems))
	for i, v := range anotherListElems {
		s.AnotherList[i].UnmarshalResource(v.(map[string]interface{}))
	}

	mapElems := d["map"].(map[string]interface{})
	s.Map = make(map[string]int, len(mapElems))
	for k, v := range mapElems {
		s.Map[k] = v.(int)
	}

	setElems := d["set"].(*schema.Set)
	s.Set = make(map[int]Bar, setElems.Len())
	for _, v := range setElems.List() {
		t := Bar{}
		t.UnmarshalResource(v.(map[string]interface{}))
		s.Set[setElems.F(v)] = t
	}
}

func (s *Foo) MarshalResource() map[string]interface{} {
	d := make(map[string]interface{})

	d["a"] = s.A

	d["b"] = s.B

	d["list"] = s.List

	anotherList := make([]map[string]interface{}, len(s.AnotherList))
	for i, v := range s.AnotherList {
		anotherList[i] = v.MarshalResource()
	}
	d["another_list"] = anotherList

	d["map"] = s.Map

	set := &schema.Set{F: tfutils.GetHashFunc(&Bar{})}
	for _, v := range s.Set {
		set.Add(v.MarshalResource())
	}
	d["set"] = set

	return d
}
