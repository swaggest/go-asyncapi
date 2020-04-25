package asyncapi

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/swaggest/go-asyncapi/spec-2.0.0"
	"github.com/swaggest/jsonschema-go"
)

// Reflector generates AsyncAPI definitions from provided message samples.
type Reflector struct {
	jsonschema.Reflector
	Data *spec.AsyncAPI
}

// DataEns ensures AsyncAPI Data.
func (r *Reflector) DataEns() *spec.AsyncAPI {
	if r.Data == nil {
		r.Data = &spec.AsyncAPI{}
	}

	return r.Data
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

	r.DataEns().WithChannelsItem(info.Name, channelItem)

	return nil
}

func (r *Reflector) collectDefinition(name string, schema jsonschema.Schema) {
	if r.Data.ComponentsEns().Schemas == nil {
		r.Data.ComponentsEns().Schemas = make(map[string]jsonschema.Schema, 1)
	}

	r.Data.ComponentsEns().Schemas[name] = schema
}

func (r *Reflector) makeOperation(channelItem *spec.ChannelItem, m *MessageSample) (*spec.Operation, error) {
	if m.MessageSample == nil {
		return &spec.Operation{
			Message: &spec.Message{
				Entity: &m.MessageEntity,
			},
		}, nil
	}

	schema, err := r.Reflect(m.MessageSample,
		jsonschema.RootRef,
		jsonschema.DefinitionsPrefix("#/components/schemas/"),
		jsonschema.CollectDefinitions(r.collectDefinition),
	)

	if err != nil {
		return nil, err
	}

	m.MessageEntity.Payload = &schema

	headerSchema, err := r.Reflect(m.MessageSample,
		jsonschema.PropertyNameTag("header"),
		jsonschema.DefinitionsPrefix("#/components/schemas/"),
		jsonschema.CollectDefinitions(r.collectDefinition),
	)

	if err != nil {
		return nil, err
	}

	if len(headerSchema.Properties) > 0 {
		m.MessageEntity.Headers = &headerSchema
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
				Schema: paramSchema.TypeObjectEns(),
			}

			if schema.Description != nil {
				param.Description = *schema.Description
			}

			channelItem.Parameters[name] = param
		}
	}

	if m.Payload.Ref != nil {
		messageName := strings.TrimPrefix(*m.Payload.Ref, "#/components/schemas/")
		r.Data.ComponentsEns().WithMessagesItem(messageName, spec.Message{
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
