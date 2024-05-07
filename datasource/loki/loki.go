package loki

import (
	"encoding/json"

	"github.com/K-Phoen/sdk"
	"github.com/lueurxax/grabana/datasource"
)

var _ datasource.Datasource = Loki{}

type Loki struct {
	builder *sdk.Datasource
}

type Option func(datasource *Loki)

func New(name string, url string, options ...Option) Loki {
	jaeger := &Loki{
		builder: &sdk.Datasource{
			Name:           name,
			Type:           "loki",
			Access:         "proxy",
			URL:            url,
			JSONData:       map[string]interface{}{},
			SecureJSONData: map[string]interface{}{},
		},
	}

	for _, opt := range options {
		opt(jaeger)
	}

	return *jaeger
}

func (datasource Loki) Name() string {
	return datasource.builder.Name
}

func (datasource Loki) MarshalJSON() ([]byte, error) {
	return json.Marshal(datasource.builder)
}
