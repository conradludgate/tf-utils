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

	baz := Baz{}
	m["another_list"] = &schema.Schema{
		Type:     schema.TypeList,
		Elem:     &schema.Resource{Schema: baz.IntoSchemaMap()},
		Optional: true,
		Description: `Complex lists can too, as long as Baz
implements the 'Schema' interface`,
	}

	m["map"] = &schema.Schema{
		Type:        schema.TypeMap,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Optional:    true,
		Description: "Maps can only be over simple types (terraform limitation)",
	}

	bar := Bar{}
	m["set"] = &schema.Schema{
		Type:     schema.TypeSet,
		Elem:     &schema.Resource{Schema: bar.IntoSchemaMap()},
		Optional: true,
		Description: `map[int]... represents a Set.
If Bar implements the 'Set' interface,
then that will be the Set function`,
	}
	var bar_i interface{} = bar
	if set, ok := bar_i.(tfutils.Set); ok {
		m["set"].Set = tfutils.BuildHashFunc(set)
	}

	return m
}

func (s *Foo) UnmarshalResource(d map[string]interface{}) {
	s.A = d["a"].(string)

	s.B = d["b"].(int)

	list := d["list"].([]interface{})
	s.List = make([]string, len(list))
	for i, v := range list {
		s.List[i] = v.(string)
	}

	anotherList := d["another_list"].([]interface{})
	s.AnotherList = make([]Baz, len(anotherList))
	for i, v := range anotherList {
		s.AnotherList[i].UnmarshalResource(v.(map[string]interface{}))
	}

	map_ := d["map"].(map[string]interface{})
	s.Map = make(map[string]int, len(map_))
	for k, v := range map_ {
		s.Map[k] = v.(int)
	}

	set := d["set"].(*schema.Set)
	s.Set = make(map[int]Bar, set.Len())
	for _, v := range set.List() {
		bar := Bar{}
		bar.UnmarshalResource(v.(map[string]interface{}))
		s.Set[set.F(v)] = bar
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
