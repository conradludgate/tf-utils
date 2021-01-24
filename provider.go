package tfutils

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

type Provider interface {
	Schema
	Configure() error
	Client() interface{}
	Resources() map[string]CRUD
	Datasources() map[string]Data
}

func ProviderFunc(p Provider) plugin.ProviderFunc {
	return func() *schema.Provider {
		sm := p.IntoSchemaMap()

		pr := p.Resources()
		resources := make(map[string]*schema.Resource, len(pr))
		for key, v := range pr {
			resources[key] = BuildCRUD(v)
		}
		pd := p.Datasources()
		datasources := make(map[string]*schema.Resource, len(pd))
		for key, v := range pd {
			datasources[key] = BuildData(v)
		}

		return &schema.Provider{
			Schema: sm,
			ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
				p.UnmarshalResource(load(d, sm))
				if err := p.Configure(); err != nil {
					return nil, err
				}
				return p.Client(), nil
			},
			ResourcesMap:   resources,
			DataSourcesMap: datasources,
		}
	}
}
