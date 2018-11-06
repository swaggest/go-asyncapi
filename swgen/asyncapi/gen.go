package asyncapi

import (
	"github.com/pkg/errors"
	"github.com/swaggest/go-asyncapi/spec"
	"github.com/swaggest/swgen"
	"net/http"
)

// Generator generates AsyncAPI definitions from provided message samples
type Generator struct {
	Swg  *swgen.Generator
	Data *spec.AsyncAPI
}

// Message is a structure that keeps general info and message sample (optional)
type Message struct {
	// pkg.Message holds general message info
	spec.Message

	// MessageSample holds a sample of message to be converted to JSON Schema, e.g. `new(MyMessage)`
	MessageSample interface{}
}

// TopicInfo keeps user-defined information about topic
type TopicInfo struct {
	Topic      string // event.{streetlightId}.lighting.measured
	Deprecated bool
	Publish    *Message
	Subscribe  *Message
}

// AddTopic adds user-defined topic to AsyncAPI definition
func (g Generator) AddTopic(info TopicInfo) error {
	var err error
	topicItem := spec.TopicItem{
		Deprecated: info.Deprecated,
	}

	if g.Swg == nil {
		g.Swg = swgen.NewGenerator()
	}

	if g.Data == nil {
		g.Data = &spec.AsyncAPI{}
	}

	if g.Data.Components == nil {
		g.Data.Components = &spec.Components{}
	}

	if g.Data.Components.Schemas == nil {
		g.Data.Components.Schemas = make(map[string]map[string]interface{})
	}

	if info.Publish != nil {
		topicItem.Publish, err = g.makeOperation(&topicItem, info.Publish)
		if err != nil {
			return errors.Wrapf(err, "failed process publish operation for topic %s", info.Topic)
		}
	}

	if info.Subscribe != nil {
		topicItem.Subscribe, err = g.makeOperation(&topicItem, info.Subscribe)
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

func (g Generator) makeOperation(topicItem *spec.TopicItem, m *Message) (*spec.Operation, error) {
	if m.MessageSample == nil {
		return &spec.Operation{
			Message: &m.Message,
		}, nil
	}

	fakeInfo := swgen.PathItemInfo{
		Method:  http.MethodPost,
		Request: m.MessageSample,
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
	}

	if _, ok := groups[`body`]; ok {
		msg.Payload = groups[`body`].Properties[`body`]
	}

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

	return &spec.Operation{
		Message: &msg,
	}, nil
}
