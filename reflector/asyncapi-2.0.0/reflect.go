// Package asyncapi provides schema reflector.
package asyncapi

import (
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
	"github.com/swaggest/go-asyncapi/spec-2.0.0"
	"github.com/swaggest/jsonschema-go"
)

// Reflector generates AsyncAPI definitions from provided message samples.
type Reflector struct {
	jsonschema.Reflector
	Schema *spec.AsyncAPI
}

// DataEns ensures AsyncAPI Schema.
//
// Deprecated: use SchemaEns().
func (r *Reflector) DataEns() *spec.AsyncAPI {
	return r.SchemaEns()
}

// SchemaEns ensures AsyncAPI Schema.
func (r *Reflector) SchemaEns() *spec.AsyncAPI {
	if r.Schema == nil {
		r.Schema = &spec.AsyncAPI{}
	}

	return r.Schema
}

// MessageSample is a structure that keeps general info and message sample (optional).
type MessageSample struct {
	// pkg.Message holds general message info.
	spec.MessageEntity

	// MessageSample holds a sample of message to be converted to JSON Schema, e.g. `new(MyMessage)`.
	MessageSample interface{}
}

// ChannelInfo keeps user-defined information about channel.
type ChannelInfo struct {
	Name            string // event.{streetlightId}.lighting.measured
	Publish         *MessageSample
	Subscribe       *MessageSample
	BaseChannelItem *spec.ChannelItem // Optional, if set is used as a base to fill with Message data
}

// AddChannel adds user-defined channel to AsyncAPI definition.
func (r *Reflector) AddChannel(info ChannelInfo) error {
	if info.Name == "" {
		return errors.New("name is required")
	}

	var (
		channelItem = spec.ChannelItem{}
		err         error
	)

	if info.BaseChannelItem != nil {
		channelItem = *info.BaseChannelItem
	}

	if info.Publish != nil {
		channelItem.Publish, err = r.makeOperation(&channelItem, info.Publish)
		if err != nil {
			return errors.Wrapf(err, "failed process publish operation for channel %s", info.Name)
		}
	}

	if info.Subscribe != nil {
		channelItem.Subscribe, err = r.makeOperation(&channelItem, info.Subscribe)
		if err != nil {
			return errors.Wrapf(err, "failed process subscribe operation for channel %s", info.Name)
		}
	}

	r.SchemaEns().WithChannelsItem(info.Name, channelItem)

	return nil
}

func schemaToMap(schema jsonschema.Schema) map[string]interface{} {
	var m map[string]interface{}

	j, err := json.Marshal(schema)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(j, &m)
	if err != nil {
		panic(err)
	}

	return m
}

func (r *Reflector) collectDefinition(name string, schema jsonschema.Schema) {
	if r.SchemaEns().ComponentsEns().Schemas == nil {
		r.SchemaEns().ComponentsEns().Schemas = make(map[string]map[string]interface{}, 1)
	}

	r.SchemaEns().ComponentsEns().Schemas[name] = schemaToMap(schema)
}

func (r *Reflector) makeOperation(channelItem *spec.ChannelItem, m *MessageSample) (*spec.Operation, error) {
	if m.MessageSample == nil {
		return &spec.Operation{
			Message: &spec.Message{
				Entity: &m.MessageEntity,
			},
		}, nil
	}

	payloadSchema, err := r.Reflect(m.MessageSample,
		jsonschema.RootRef,
		jsonschema.DefinitionsPrefix("#/components/schemas/"),
		jsonschema.CollectDefinitions(r.collectDefinition),
	)
	if err != nil {
		return nil, err
	}

	m.MessageEntity.Payload = schemaToMap(payloadSchema)

	headerSchema, err := r.Reflect(m.MessageSample,
		jsonschema.PropertyNameTag("header"),
		jsonschema.DefinitionsPrefix("#/components/schemas/"),
		jsonschema.CollectDefinitions(r.collectDefinition),
	)
	if err != nil {
		return nil, err
	}

	if len(headerSchema.Properties) > 0 {
		m.MessageEntity.Headers = schemaToMap(headerSchema)
	}

	pathSchema, err := r.Reflect(m.MessageSample,
		jsonschema.PropertyNameTag("path"),
		jsonschema.DefinitionsPrefix("#/components/schemas/"),
		jsonschema.CollectDefinitions(r.collectDefinition),
	)
	if err != nil {
		return nil, err
	}

	if len(pathSchema.Properties) > 0 {
		if channelItem.Parameters == nil {
			channelItem.Parameters = make(map[string]spec.Parameter, len(pathSchema.Properties))
		}

		for name, paramSchema := range pathSchema.Properties {
			param := spec.Parameter{
				Schema: schemaToMap(*paramSchema.TypeObjectEns()),
			}

			if payloadSchema.Description != nil {
				param.Description = *payloadSchema.Description
			}

			channelItem.Parameters[name] = param
		}
	}

	if payloadSchema.Ref != nil {
		messageName := strings.TrimPrefix(*payloadSchema.Ref, "#/components/schemas/")
		r.SchemaEns().ComponentsEns().WithMessagesItem(messageName, spec.Message{
			Entity: &m.MessageEntity,
		})

		return &spec.Operation{
			Message: &spec.Message{
				Reference: &spec.Reference{Ref: "#/components/messages/" + messageName},
			},
		}, nil
	}

	return &spec.Operation{
		Message: &spec.Message{
			Entity: &m.MessageEntity,
		},
	}, nil
}
