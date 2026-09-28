// Package asyncapi provides schema reflector.
package asyncapi

import (
	"encoding/json"
	"fmt"
	"strings"

	spec "github.com/swaggest/go-asyncapi/spec-3.0.0"
	"github.com/swaggest/jsonschema-go"
)

// Reflector generates AsyncAPI definitions from provided message samples.
type Reflector struct {
	jsonschema.Reflector
	Schema *spec.AsyncAPI
}

// SchemaEns ensures AsyncAPI Schema.
func (r *Reflector) SchemaEns() *spec.AsyncAPI {
	if r.Schema == nil {
		r.Schema = &spec.AsyncAPI{}
	}

	return r.Schema
}

// MessageSample is a structure that keeps general message info and message sample (optional).
type MessageSample struct {
	// MessageObject holds general message info.
	spec.MessageObject

	// Sample holds a sample of message to be converted to JSON Schema, e.g. `new(MyMessage)`.
	Sample interface{}
}

// AddChannel adds a channel to AsyncAPI definition and returns a reference to it.
func (r *Reflector) AddChannel(name string, prepare ...func(ch *spec.Channel)) spec.Reference {
	ch := spec.Channel{}
	ch.WithAddress(name)

	for _, fn := range prepare {
		fn(&ch)
	}

	r.SchemaEns().WithChannelsItem(name, spec.ChannelOrRef{Channel: &ch})

	return spec.Reference{Ref: "#/channels/" + name}
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
	r.SchemaEns().ComponentsEns().WithSchemasItem(name, schemaToMap(schema))
}

// AddOperation reflects a message sample and adds a send/receive operation to a channel
// previously created with AddChannel.
func (r *Reflector) AddOperation(channel spec.Reference, action spec.OperationAction, prepareMessage func(message *MessageSample), prepare ...func(op *spec.Operation)) error {
	channelName := strings.TrimPrefix(channel.Ref, "#/channels/")

	chOrRef, found := r.SchemaEns().Channels[channelName]
	if !found || chOrRef.Channel == nil {
		return fmt.Errorf("channel %s not found, add it with AddChannel first", channelName)
	}

	ch := chOrRef.Channel

	m := MessageSample{}

	prepareMessage(&m)

	messageName, err := r.reflectMessageSample(&m, ch)
	if err != nil {
		return err
	}

	if messageName == "" {
		messageName = capitalize(string(action)) + capitalize(channelName)
	}

	if ch.Messages == nil {
		ch.Messages = make(map[string]spec.MessageObjectOrRef, 1)
	}

	ch.Messages[messageName] = spec.MessageObjectOrRef{MessageObject: &m.MessageObject}

	op := spec.Operation{}
	op.WithAction(action).
		WithChannel(channel).
		WithMessages(spec.Reference{Ref: channel.Ref + "/messages/" + messageName})

	for _, fn := range prepare {
		fn(&op)
	}

	r.SchemaEns().WithOperationsItem(string(action)+":"+channelName, spec.OperationOrRef{Operation: &op})

	return nil
}

// reflectMessageSample reflects m.Sample, if set, into payload, headers and channel path
// parameters, and returns the message name derived from the payload schema, if available.
func (r *Reflector) reflectMessageSample(m *MessageSample, ch *spec.Channel) (string, error) {
	if m.Sample == nil {
		return m.Name, nil
	}

	payloadSchema, err := r.Reflect(m.Sample,
		jsonschema.RootRef,
		jsonschema.DefinitionsPrefix("#/components/schemas/"),
		jsonschema.CollectDefinitions(r.collectDefinition),
	)
	if err != nil {
		return "", fmt.Errorf("reflecting payload schema: %w", err)
	}

	payload := interface{}(schemaToMap(payloadSchema))
	m.Payload = &payload

	messageName := m.Name
	if messageName == "" && payloadSchema.Ref != nil {
		messageName = strings.TrimPrefix(*payloadSchema.Ref, "#/components/schemas/")
	}

	headerSchema, err := r.Reflect(m.Sample,
		jsonschema.PropertyNameTag("header"),
		jsonschema.DefinitionsPrefix("#/components/schemas/"),
		jsonschema.CollectDefinitions(r.collectDefinition),
	)
	if err != nil {
		return "", fmt.Errorf("reflecting headers schema: %w", err)
	}

	if len(headerSchema.Properties) > 0 {
		headers := interface{}(schemaToMap(headerSchema))
		m.Headers = &headers
	}

	pathSchema, err := r.Reflect(m.Sample,
		jsonschema.PropertyNameTag("path"),
		jsonschema.DefinitionsPrefix("#/components/schemas/"),
		jsonschema.CollectDefinitions(r.collectDefinition),
	)
	if err != nil {
		return "", fmt.Errorf("reflecting path parameters schema: %w", err)
	}

	addPathParameters(ch, pathSchema)

	return messageName, nil
}

// addPathParameters fills channel parameters from properties of a schema reflected
// with the "path" property name tag.
func addPathParameters(ch *spec.Channel, pathSchema jsonschema.Schema) {
	if len(pathSchema.Properties) == 0 {
		return
	}

	if ch.Parameters == nil {
		ch.Parameters = make(map[string]spec.ParameterOrRef, len(pathSchema.Properties))
	}

	for name, paramSchema := range pathSchema.Properties {
		param := spec.Parameter{}

		if d := paramSchema.TypeObjectEns().Description; d != nil {
			param.Description = *d
		}

		ch.Parameters[name] = spec.ParameterOrRef{Parameter: &param}
	}
}

func capitalize(s string) string {
	if s == "" {
		return s
	}

	return strings.ToUpper(s[:1]) + s[1:]
}
