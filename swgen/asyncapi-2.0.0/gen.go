package asyncapi

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/swaggest/go-asyncapi/spec-2.0.0"
	"github.com/swaggest/swgen"
)

// Generator generates AsyncAPI definitions from provided message samples.
type Generator struct {
	Swg  *swgen.Generator
	Data spec.AsyncAPI

	channelInfos map[string]ChannelInfo
}

// Message is a structure that keeps general info and message sample (optional).
type Message struct {
	// pkg.Message holds general message info.
	spec.MessageEntity

	// MessageSample holds a sample of message to be converted to JSON Schema, e.g. `new(MyMessage)`.
	MessageSample interface{}
}

// ChannelInfo keeps user-defined information about channel.
type ChannelInfo struct {
	Name            string // event.{streetlightId}.lighting.measured
	Publish         *Message
	Subscribe       *Message
	BaseChannelItem *spec.ChannelItem // Optional, if set is used as a base to fill with Message data
}

// AddChannel adds user-defined channel to AsyncAPI definition.
func (g *Generator) AddChannel(info ChannelInfo) error {
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

	if g.Swg == nil {
		g.Swg = swgen.NewGenerator()
	}

	if g.Data.Components == nil {
		g.Data.Components = &spec.Components{}
	}

	if g.Data.Components.Schemas == nil {
		g.Data.Components.Schemas = make(map[string]map[string]interface{})
	}

	if g.Data.Components.Messages == nil {
		g.Data.Components.Messages = make(map[string]spec.Message)
	}

	if info.Publish != nil {
		channelItem.Publish, err = g.makeOperation("publish", info, &channelItem, info.Publish)
		if err != nil {
			return errors.Wrapf(err, "failed process publish operation for channel %s", info.Name)
		}
	}

	if info.Subscribe != nil {
		channelItem.Subscribe, err = g.makeOperation("subscribe", info, &channelItem, info.Subscribe)
		if err != nil {
			return errors.Wrapf(err, "failed process subscribe operation for channel %s", info.Name)
		}
	}

	if g.Data.Channels == nil {
		g.Data.Channels = make(map[string]spec.ChannelItem)
	}

	g.Data.Channels[info.Name] = channelItem
	return nil
}

func (g *Generator) makeOperation(intent string, info ChannelInfo, channelItem *spec.ChannelItem, m *Message) (*spec.Operation, error) {
	if m.MessageSample == nil {
		return &spec.Operation{
			Message: &spec.Message{
				Entity: &m.MessageEntity,
			},
		}, nil
	}

	if g.channelInfos == nil {
		g.channelInfos = make(map[string]ChannelInfo)
	}
	path := "/" + intent + "/" + info.Name
	g.channelInfos[path] = info

	fakeInfo := swgen.PathItemInfo{
		Path:     path,
		Method:   http.MethodPost,
		Request:  m.MessageSample,
		Response: m.MessageSample,
	}
	obj := g.Swg.SetPathItem(fakeInfo)

	cfg := swgen.JSONSchemaConfig{
		DefinitionsPrefix:  "#/components/schemas/",
		StripDefinitions:   true,
		CollectDefinitions: g.Data.Components.Schemas,
	}
	groups, err := g.Swg.GetJSONSchemaRequestGroups(obj, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make schema")
	}

	msg := m.MessageEntity
	if headerSchema, ok := groups[`header`]; ok {
		msg.Headers, err = headerSchema.ToMap()
		if err != nil {
			return nil, err
		}
		delete(msg.Headers, "$schema")
	}

	body, err := g.Swg.GetJSONSchemaRequestBody(obj, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make body schema")
	}

	msg.Payload = body

	for _, param := range obj.Parameters {
		if param.In == `path` {
			schema, err := g.Swg.ParamJSONSchema(param, cfg)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to process param schema for %s", param.Name)
			}

			if channelItem.Parameters == nil {
				channelItem.Parameters = map[string]spec.Parameter{}
			}

			channelItem.Parameters[param.Name] = spec.Parameter{
				Description: param.Description,
				Schema:      schema,
			}
		}
	}

	if ref, ok := msg.Payload["$ref"].(string); ok && ref != "" {
		messageName := strings.TrimPrefix(ref, "#/components/schemas/")
		g.Data.Components.Messages[messageName] = spec.Message{
			Entity: &msg,
		}

		return &spec.Operation{
			Message: &spec.Message{
				Reference: &spec.Reference{Ref: "#/components/messages/" + messageName},
			},
		}, nil
	}

	return &spec.Operation{
		Message: &spec.Message{
			Entity: &msg,
		},
	}, nil
}

// WalkJSONSchemas iterates thorough message payload schemas
func (g *Generator) WalkJSONSchemas(w func(isPublishing bool, info ChannelInfo, schema map[string]interface{})) error {
	return g.Swg.WalkJSONSchemaResponses(func(path, _ string, _ int, schema map[string]interface{}) {
		intent := strings.Split(path, "/")[1]
		w(intent == "publish", g.channelInfos[path], schema)
	})
}
