package asyncapi_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/swaggest/go-asyncapi/amqp"
	"github.com/swaggest/go-asyncapi/spec"           // nolint:staticcheck // Deprecated package.
	"github.com/swaggest/go-asyncapi/swgen/asyncapi" // nolint:staticcheck // Deprecated package.
)

func TestGenerator_WalkJSONSchemas(t *testing.T) {
	type SubItem struct {
		Key    string  `json:"key" description:"Item key"`
		Values []int64 `json:"values" uniqueItems:"true" description:"List of item values"`
	}

	type MyMessage struct {
		Name      string    `path:"name" description:"Name"`
		CreatedAt time.Time `json:"createdAt" description:"Creation time"`
		Items     []SubItem `json:"items" description:"List of items"`
	}

	type MyAnotherMessage struct {
		TraceID string  `header:"X-Trace-ID" description:"Tracing header" required:"true"`
		Item    SubItem `json:"item" description:"Some item"`
	}

	g := asyncapi.Generator{
		Data: spec.AsyncAPI{
			Asyncapi: spec.Asyncapi120,
			Servers: []spec.Server{
				{
					URL:    "api.lovely.com:{port}",
					Scheme: spec.ServerSchemeAmqp,
				},
			},
			Info: spec.Info{
				Version: "1.2.3", //required
				Title:   "My Lovely Messaging API",
			},
		},
	}
	assert.NoError(t, g.AddTopic(asyncapi.TopicInfo{
		Topic: "one.{name}.two",
		Publish: &asyncapi.Message{
			Message: spec.Message{
				Description: "This is a sample schema",
				Summary:     "Sample publisher",
			},
			MessageSample: new(MyMessage),
		},
	}))

	assert.NoError(t, g.AddTopic(asyncapi.TopicInfo{
		Topic: "another.one",
		Subscribe: &asyncapi.Message{
			Message: spec.Message{
				Description: "This is another sample schema",
				Summary:     "Sample consumer",
			},
			MessageSample: new(MyAnotherMessage),
		},
	}))

	assert.NoError(t, g.WalkJSONSchemas(func(isPublishing bool, info asyncapi.TopicInfo, schema map[string]interface{}) {
		js, err := json.Marshal(schema)
		assert.NoError(t, err)
		switch info.Topic {
		case "one.{name}.two":
			assert.Equal(t, `{"$schema":"http://json-schema.org/draft-04/schema#","definitions":{"SubItem":{"properties":{"key":{"description":"Item key","type":"string"},"values":{"description":"List of item values","items":{"format":"int64","type":"integer"},"type":"array","uniqueItems":true}},"type":"object"}},"properties":{"createdAt":{"description":"Creation time","format":"date-time","type":"string"},"items":{"description":"List of items","items":{"$ref":"#/definitions/SubItem"},"type":"array"}},"type":"object"}`, string(js))
		case "another.one":
			assert.Equal(t, `{"$schema":"http://json-schema.org/draft-04/schema#","definitions":{"SubItem":{"properties":{"key":{"description":"Item key","type":"string"},"values":{"description":"List of item values","items":{"format":"int64","type":"integer"},"type":"array","uniqueItems":true}},"type":"object"}},"properties":{"item":{"$ref":"#/definitions/SubItem"}},"type":"object"}`, string(js))
		default:
			t.Fatal("should not get here")
		}
	}))
}

func TestExample(t *testing.T) {
	type SubItem struct {
		Key    string  `json:"key" description:"Item key"`
		Values []int64 `json:"values" uniqueItems:"true" description:"List of item values"`
	}

	type MyMessage struct {
		Name      string    `path:"name" description:"Name"`
		CreatedAt time.Time `json:"createdAt" description:"Creation time"`
		Items     []SubItem `json:"items" description:"List of items"`
	}

	type MyAnotherMessage struct {
		TraceID string  `header:"X-Trace-ID" description:"Tracing header" required:"true"`
		Item    SubItem `json:"item" description:"Some item"`
	}

	g := asyncapi.Generator{
		Data: spec.AsyncAPI{
			Asyncapi: spec.Asyncapi120,
			Servers: []spec.Server{
				{
					URL:    "api.lovely.com:{port}",
					Scheme: spec.ServerSchemeAmqp,
				},
			},
			Info: spec.Info{
				Version: "1.2.3", // required
				Title:   "My Lovely Messaging API",
			},
		},
	}
	must := func(err error) {
		if err != nil {
			panic(err.Error())
		}
	}
	must(g.AddTopic(asyncapi.TopicInfo{
		Topic: "one.{name}.two",
		Publish: amqp.MessageWithInfo(&asyncapi.Message{
			Message: spec.Message{
				Description: "This is a sample schema.",
				Summary:     "Sample publisher",
			},
			MessageSample: new(MyMessage),
		}, amqp.Info{
			Exchange: "some-exchange",
			VHost:    "some-vhost",
			Queue:    "some-queue",
		}),
	}))

	must(g.AddTopic(asyncapi.TopicInfo{
		Topic: "another.one",
		Subscribe: &asyncapi.Message{
			Message: *(&spec.Message{}).
				WithDescription("This is another sample schema.").
				WithSummary("Sample consumer"),
			MessageSample: new(MyAnotherMessage),
		},
	}))

	yaml, err := g.Data.MarshalYAML()
	must(err)

	expected := `asyncapi: 1.2.0
info:
  title: My Lovely Messaging API
  version: 1.2.3
servers:
- scheme: amqp
  url: api.lovely.com:{port}
topics:
  another.one:
    subscribe:
      $ref: '#/components/messages/MyAnotherMessage'
  one.{name}.two:
    parameters:
    - description: Name
      name: name
      schema:
        description: Name
        type: string
    publish:
      $ref: '#/components/messages/MyMessage'
components:
  messages:
    MyAnotherMessage:
      description: This is another sample schema.
      headers:
        properties:
          X-Trace-ID:
            description: Tracing header
            type: string
        required:
        - X-Trace-ID
        type: object
      payload:
        $ref: '#/components/schemas/MyAnotherMessage'
      summary: Sample consumer
    MyMessage:
      description: |-
        This is a sample schema.

        AMQP VHost: some-vhost.

        AMQP Exchange: some-exchange.

        AMQP Queue: some-queue.
      payload:
        $ref: '#/components/schemas/MyMessage'
      summary: Sample publisher
      x-amqp-exchange: some-exchange
      x-amqp-queue: some-queue
      x-amqp-vhost: some-vhost
  schemas:
    MyAnotherMessage:
      properties:
        item:
          $ref: '#/components/schemas/SubItem'
      type: object
    MyMessage:
      properties:
        createdAt:
          description: Creation time
          format: date-time
          type: string
        items:
          description: List of items
          items:
            $ref: '#/components/schemas/SubItem'
          type: array
      type: object
    SubItem:
      properties:
        key:
          description: Item key
          type: string
        values:
          description: List of item values
          items:
            format: int64
            type: integer
          type: array
          uniqueItems: true
      type: object
`
	assert.Equal(t, expected, string(yaml))
}
