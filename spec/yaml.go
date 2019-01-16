package spec

import (
	"encoding/json"

	"github.com/icza/dyno" //todo remove when gopkg.in/yaml provide a solution to avoid `map[interface{}]interface{}`
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

// UnmarshalYAML reads from YAML bytes
func (i *AsyncAPI) UnmarshalYAML(data []byte) error {
	var v interface{}
	err := yaml.Unmarshal(data, &v)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal YAML")
	}
	v = dyno.ConvertMapI2MapS(v)
	data, err = json.Marshal(v)
	if err != nil {
		return errors.Wrap(err, "failed to marshal JSON")
	}
	return i.UnmarshalJSON(data)
}

// MarshalYAML produces YAML bytes
func (i *AsyncAPI) MarshalYAML() ([]byte, error) {
	jsonData, err := i.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal JSON")
	}
	var v interface{}
	err = json.Unmarshal(jsonData, &v)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal JSON")
	}
	return yaml.Marshal(v)
}
