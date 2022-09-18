package spec

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"gopkg.in/yaml.v2"
)

// UnmarshalYAML reads from YAML bytes.
func (i *AsyncAPI) UnmarshalYAML(data []byte) error {
	var v interface{}

	err := yaml.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	v = convertMapI2MapS(v)

	data, err = json.Marshal(v)
	if err != nil {
		return err
	}

	return i.UnmarshalJSON(data)
}

// MarshalYAML produces YAML bytes.
func (i *AsyncAPI) MarshalYAML() ([]byte, error) {
	jsonData, err := i.MarshalJSON()
	if err != nil {
		return nil, err
	}

	var v orderedMap

	err = json.Unmarshal(jsonData, &v)
	if err != nil {
		return nil, err
	}

	return yaml.Marshal(yaml.MapSlice(v))
}

type orderedMap []yaml.MapItem

func (om *orderedMap) UnmarshalJSON(data []byte) error {
	keys, err := objectKeys(data)
	if err != nil {
		return err
	}

	var mapData map[string]json.RawMessage

	err = json.Unmarshal(data, &mapData)
	if err != nil {
		return err
	}

	for _, key := range keys {
		jsonVal := mapData[key]
		_, err = objectKeys(jsonVal)

		var val interface{}

		if err == nil {
			v := make(orderedMap, 0)

			err = json.Unmarshal(jsonVal, &v)
			if err != nil {
				return err
			}

			val = yaml.MapSlice(v)
		} else {
			err = json.Unmarshal(jsonVal, &val)
			if err != nil {
				return err
			}
		}

		*om = append(*om, yaml.MapItem{
			Key:   key,
			Value: val,
		})
	}

	return nil
}

func objectKeys(b []byte) ([]string, error) {
	d := json.NewDecoder(bytes.NewReader(b))

	t, err := d.Token()
	if err != nil {
		return nil, err
	}

	if t != json.Delim('{') {
		return nil, errors.New("expected start of object")
	}

	var keys []string

	for {
		t, err := d.Token()
		if err != nil {
			return nil, err
		}

		if t == json.Delim('}') {
			return keys, nil
		}

		keys = append(keys, t.(string))

		if err := skipValue(d); err != nil {
			return nil, err
		}
	}
}

var errEnd = errors.New("invalid errEnd of array or object")

func skipValue(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}

	switch t {
	case json.Delim('['), json.Delim('{'):
		for {
			if err := skipValue(d); err != nil {
				if errors.Is(err, errEnd) {
					break
				}

				return err
			}
		}
	case json.Delim(']'), json.Delim('}'):
		return errEnd
	}

	return nil
}

// convertMapI2MapS walks the given dynamic object recursively, and
// converts maps with interface{} key type to maps with string key type.
// This function comes handy if you want to marshal a dynamic object into
// JSON where maps with interface{} key type are not allowed.
//
// Recursion is implemented into values of the following types:
//
//	-map[interface{}]interface{}
//	-map[string]interface{}
//	-[]interface{}
//
// When converting map[interface{}]interface{} to map[string]interface{},
// fmt.Sprint() with default formatting is used to convert the key to a string key.
//
// See github.com/icza/dyno.
func convertMapI2MapS(v interface{}) interface{} {
	switch x := v.(type) {
	case map[interface{}]interface{}:
		m := map[string]interface{}{}

		for k, v2 := range x {
			switch k2 := k.(type) {
			case string: // Fast check if it's already a string
				m[k2] = convertMapI2MapS(v2)
			default:
				m[fmt.Sprint(k)] = convertMapI2MapS(v2)
			}
		}

		v = m

	case []interface{}:
		for i, v2 := range x {
			x[i] = convertMapI2MapS(v2)
		}

	case map[string]interface{}:
		for k, v2 := range x {
			x[k] = convertMapI2MapS(v2)
		}
	}

	return v
}
