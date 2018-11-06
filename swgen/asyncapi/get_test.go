package asyncapi_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/swaggest/go-asyncapi/spec"
	"github.com/swaggest/go-asyncapi/swgen/asyncapi"
	"io/ioutil"
	"testing"
	"time"
)

func TestGenerator_AddTopic(t *testing.T) {
	type SubItem struct {
		Key    string  `json:"key"`
		Values []int64 `json:"values" uniqueItems:"true"`
	}

	type MyMessage struct {
		Name      string    `path:"name"`
		CreatedAt time.Time `json:"createdAt"`
		Items     []SubItem `json:"items"`
	}

	type MyAnotherMessage struct {
		TraceID string  `header:"X-Trace-ID"`
		Item    SubItem `json:"item"`
	}

	g := asyncapi.Generator{
		Data: &spec.AsyncAPI{
			Asyncapi: spec.Asyncapi120,
			Servers: []spec.Server{
				{
					URL:    "api.streetlights.smartylighting.com:{port}",
					Scheme: spec.ServerSchemeAmqp,
				},
			},
			BaseTopic: "smartylighting.streetlights.1.0",
			Info: &spec.Info{
				Title: "My Lovely Messaging API",
			},
		},
	}
	err := g.AddTopic(asyncapi.TopicInfo{
		Topic: "one.{name}.two",
		Publish: &asyncapi.Message{
			Message: spec.Message{
				Description: "This is a sample schema",
				Summary:     "Sample publisher",
			},
			MessageSample: new(MyMessage),
		},
	})
	assert.NoError(t, err)

	err = g.AddTopic(asyncapi.TopicInfo{
		Topic: "another.one",
		Subscribe: &asyncapi.Message{
			Message: spec.Message{
				Description: "This is another sample schema",
				Summary:     "Sample consumer",
			},
			MessageSample: new(MyAnotherMessage),
		},
	})
	assert.NoError(t, err)

	yaml, err := g.Data.MarshalYAML()
	assert.NoError(t, err)
	//nolint:errcheck
	ioutil.WriteFile("sample.yaml", yaml, 0666)
	assert.Equal(t, `asyncapi: 1.2.0
baseTopic: smartylighting.streetlights.1.0
components:
  schemas:
    MyAnotherMessage:
      properties:
        item:
          $ref: '#/components/schemas/SubItem'
      type: object
    MyMessage:
      properties:
        createdAt:
          format: date-time
          type: string
        items:
          items:
            $ref: '#/components/schemas/SubItem'
          type: array
      type: object
    SubItem:
      properties:
        key:
          type: string
        values:
          items:
            format: int64
            type: integer
          type: array
          uniqueItems: true
      type: object
info:
  title: My Lovely Messaging API
servers:
- scheme: amqp
  url: api.streetlights.smartylighting.com:{port}
topics:
  another.one:
    subscribe:
      description: This is another sample schema
      headers:
        $schema: http://json-schema.org/draft-04/schema#
        properties:
          X-Trace-ID:
            type: string
        required:
        - X-Trace-ID
        type: object
      payload:
        $ref: '#/components/schemas/MyAnotherMessage'
      summary: Sample consumer
  one.{name}.two:
    parameters:
    - name: name
      schema:
        type: string
    publish:
      description: This is a sample schema
      payload:
        $ref: '#/components/schemas/MyMessage'
      summary: Sample publisher
`, string(yaml))
}
