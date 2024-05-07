package tempo

import (
	"encoding/json"

	"github.com/K-Phoen/sdk"
	"github.com/lueurxax/grabana/datasource"
)

var _ datasource.Datasource = Tempo{}

type Tempo struct {
	builder *sdk.Datasource
}

type Option func(datasource *Tempo)
type TraceToLogsOption func(settings map[string]interface{})

func New(name string, url string, options ...Option) Tempo {
	jaeger := &Tempo{
		builder: &sdk.Datasource{
			Name:           name,
			Type:           "tempo",
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

func (datasource Tempo) Name() string {
	return datasource.builder.Name
}

func (datasource Tempo) MarshalJSON() ([]byte, error) {
	return json.Marshal(datasource.builder)
}
