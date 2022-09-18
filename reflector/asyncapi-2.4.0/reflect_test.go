package asyncapi_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/swaggest/assertjson"
	"github.com/swaggest/go-asyncapi/reflector/asyncapi-2.4.0"
	"github.com/swaggest/go-asyncapi/spec-2.4.0"
)

func TestReflector_AddChannel(t *testing.T) {
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

	asyncAPI := spec.AsyncAPI{}
	asyncAPI.AddServer("production", spec.Server{
		URL:             "api.lovely.com:{port}",
		Protocol:        "amqp",
		ProtocolVersion: "AMQP 0.9.1",
	})

	asyncAPI.Info.Version = "1.2.3"
	asyncAPI.Info.Title = "My Lovely Messaging API"

	r := asyncapi.Reflector{Schema: &asyncAPI}
	assert.NoError(t, r.AddChannel(asyncapi.ChannelInfo{
		Name: "one.{name}.two",
		Publish: &asyncapi.MessageSample{
			MessageEntity: spec.MessageEntity{
				Description: "This is a sample schema",
				Summary:     "Sample publisher",
			},
			MessageSample: new(MyMessage),
		},
	}))

	assert.NoError(t, r.AddChannel(asyncapi.ChannelInfo{
		Name: "another.one",
		Subscribe: &asyncapi.MessageSample{
			MessageEntity: spec.MessageEntity{
				Description: "This is another sample schema",
				Summary:     "Sample consumer",
			},
			MessageSample: new(MyAnotherMessage),
		},
	}))

	assertjson.EqualMarshal(t, []byte(`{
	  "asyncapi":"2.4.0",
	  "info":{"title":"My Lovely Messaging API","version":"1.2.3"},
	  "servers":{
		"production":{
		  "url":"api.lovely.com:{port}","protocol":"amqp",
		  "protocolVersion":"AMQP 0.9.1"
		}
	  },
	  "channels":{
		"another.one":{
		  "subscribe":{
			"message":{"$ref":"#/components/messages/Asyncapi240TestMyAnotherMessage"}
		  }
		},
		"one.{name}.two":{
		  "parameters":{"name":{"schema":{"description":"Name","type":"string"}}},
		  "publish":{"message":{"$ref":"#/components/messages/Asyncapi240TestMyMessage"}}
		}
	  },
	  "components":{
		"schemas":{
		  "Asyncapi240TestMyAnotherMessage":{
			"properties":{
			  "item":{
				"$ref":"#/components/schemas/Asyncapi240TestSubItem",
				"description":"Some item"
			  }
			},
			"type":"object"
		  },
		  "Asyncapi240TestMyMessage":{
			"properties":{
			  "createdAt":{"description":"Creation time","format":"date-time","type":"string"},
			  "items":{
				"description":"List of items",
				"items":{"$ref":"#/components/schemas/Asyncapi240TestSubItem"},
				"type":["array","null"]
			  }
			},
			"type":"object"
		  },
		  "Asyncapi240TestSubItem":{
			"properties":{
			  "key":{"description":"Item key","type":"string"},
			  "values":{
				"description":"List of item values","items":{"type":"integer"},
				"type":["array","null"],"uniqueItems":true
			  }
			},
			"type":"object"
		  }
		},
		"messages":{
		  "Asyncapi240TestMyAnotherMessage":{
			"headers":{
			  "properties":{"X-Trace-ID":{"description":"Tracing header","type":"string"}},
			  "required":["X-Trace-ID"],"type":"object"
			},
			"payload":{"$ref":"#/components/schemas/Asyncapi240TestMyAnotherMessage"},
			"summary":"Sample consumer",
			"description":"This is another sample schema"
		  },
		  "Asyncapi240TestMyMessage":{
			"payload":{"$ref":"#/components/schemas/Asyncapi240TestMyMessage"},
			"summary":"Sample publisher","description":"This is a sample schema"
		  }
		}
	  }
	}`), r.Schema)
}
