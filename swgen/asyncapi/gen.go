package asyncapi

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/swaggest/go-asyncapi/spec" // nolint:staticcheck // Deprecated.
	"github.com/swaggest/swgen"
)

// Generator generates AsyncAPI definitions from provided message samples.
type Generator struct {
	Swg  *swgen.Generator
	Data spec.AsyncAPI

	pathItems map[string]TopicInfo
}

// Message is a structure that keeps general info and message sample (optional).
type Message struct {
	// Message holds general message info.
	spec.Message

	// MessageSample holds a sample of message to be converted to JSON Schema, e.g. `new(MyMessage)`.
	MessageSample interface{}
}

// TopicInfo keeps user-defined information about topic.
type TopicInfo struct {
	Topic         string // event.{streetlightId}.lighting.measured
	Publish       *Message
	Subscribe     *Message
	BaseTopicItem *spec.TopicItem // Optional, if set is used as a base to fill with Message data.
}

// AddTopic adds user-defined topic to AsyncAPI definition.
func (g *Generator) AddTopic(info TopicInfo) error {
	if info.Topic == "" {
		return errors.New("topic is required")
	}

	var (
		topicItem = spec.TopicItem{}
		err       error
	)

	if info.BaseTopicItem != nil {
		topicItem = *info.BaseTopicItem
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
		topicItem.Publish, err = g.makeOperation("publish", info, &topicItem, info.Publish)
		if err != nil {
			return errors.Wrapf(err, "failed process publish operation for topic %s", info.Topic)
		}
	}

	if info.Subscribe != nil {
		topicItem.Subscribe, err = g.makeOperation("subscribe", info, &topicItem, info.Subscribe)
		if err != nil {
			return errors.Wrapf(err, "failed process subscribe operation for topic %s", info.Topic)
		}
	}

	if g.Data.Topics == nil {
		g.Data.Topics = &spec.Topics{}
	}

	if g.Data.Topics.MapOfTopicItemValues == nil {
		g.Data.Topics.MapOfTopicItemValues = make(map[string]spec.TopicItem)
	}

	g.Data.Topics.MapOfTopicItemValues[info.Topic] = topicItem

	return nil
}

func (g *Generator) makeOperation(intent string, info TopicInfo, topicItem *spec.TopicItem, m *Message) (*spec.Operation, error) {
	if m.MessageSample == nil {
		return &spec.Operation{
			Message: &m.Message,
		}, nil
	}

	if g.pathItems == nil {
		g.pathItems = make(map[string]TopicInfo)
	}

	path := "/" + intent + "/" + info.Topic
	g.pathItems[path] = info

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

	msg := m.Message
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

			topicItem.Parameters = append(topicItem.Parameters, spec.Parameter{
				Name:        param.Name,
				Description: param.Description,
				Schema:      schema,
			})
		}
	}

	if ref, ok := msg.Payload["$ref"].(string); ok && ref != "" {
		messageName := strings.TrimPrefix(ref, "#/components/schemas/")
		g.Data.Components.Messages[messageName] = msg

		return &spec.Operation{
			Message: &spec.Message{
				Ref: "#/components/messages/" + messageName,
			},
		}, nil
	}

	return &spec.Operation{
		Message: &msg,
	}, nil
}

// WalkJSONSchemas iterates thorough message payload schemas.
func (g *Generator) WalkJSONSchemas(w func(isPublishing bool, info TopicInfo, schema map[string]interface{})) error {
	return g.Swg.WalkJSONSchemaResponses(func(path, _ string, _ int, schema map[string]interface{}) {
		intent := strings.Split(path, "/")[1]
		w(intent == "publish", g.pathItems[path], schema)
	})
}
