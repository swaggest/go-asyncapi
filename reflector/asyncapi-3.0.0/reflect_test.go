package asyncapi_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/swaggest/assertjson"
	"github.com/swaggest/go-asyncapi/reflector/asyncapi-3.0.0"
	spec "github.com/swaggest/go-asyncapi/spec-3.0.0"
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
	asyncAPI.AddServer("production", func(srv *spec.Server) {
		srv.Host = "api.lovely.com:{port}"
		srv.Protocol = "amqp"
		srv.ProtocolVersion = "AMQP 0.9.1"
	})

	asyncAPI.UpdateInfo(func(i *spec.Info) {
		i.Version = "1.2.3"
		i.Title = "My Lovely Messaging API"
	})

	r := asyncapi.Reflector{Schema: &asyncAPI}

	chRef := r.AddChannel("one.{name}.two")
	assert.NoError(t, r.AddOperation(chRef, spec.OperationActionSend, func(message *asyncapi.MessageSample) {
		message.Sample = new(MyMessage)
		message.Description = "This is a sample schema"
		message.Summary = "Sample publisher"
	}))

	anotherRef := r.AddChannel("another.one")
	assert.NoError(t, r.AddOperation(anotherRef, spec.OperationActionReceive, func(message *asyncapi.MessageSample) {
		message.Sample = new(MyAnotherMessage)
		message.Description = "This is another sample schema"
		message.Summary = "Sample consumer"
	}))

	require.Error(t, r.AddOperation(spec.Reference{Ref: "#/channels/missing"}, spec.OperationActionSend,
		func(message *asyncapi.MessageSample) {}))

	assertjson.EqualMarshal(t, []byte(`{
	  "asyncapi":"3.0.0",
	  "info":{"title":"My Lovely Messaging API","version":"1.2.3"},
	  "servers":{
		"production":{
		  "host":"api.lovely.com:{port}","protocol":"amqp",
		  "protocolVersion":"AMQP 0.9.1"
		}
	  },
	  "channels":{
		"another.one":{
		  "address":"another.one",
		  "messages":{
			"Asyncapi300TestMyAnotherMessage":{
			  "headers":{
				"properties":{"X-Trace-ID":{"description":"Tracing header","type":"string"}},
				"required":["X-Trace-ID"],"type":"object"
			  },
			  "payload":{"$ref":"#/components/schemas/Asyncapi300TestMyAnotherMessage"},
			  "summary":"Sample consumer",
			  "description":"This is another sample schema"
			}
		  }
		},
		"one.{name}.two":{
		  "address":"one.{name}.two",
		  "messages":{
			"Asyncapi300TestMyMessage":{
			  "payload":{"$ref":"#/components/schemas/Asyncapi300TestMyMessage"},
			  "summary":"Sample publisher","description":"This is a sample schema"
			}
		  },
		  "parameters":{"name":{"description":"Name"}}
		}
	  },
	  "operations":{
		"receive:another.one":{
		  "action":"receive",
		  "channel":{"$ref":"#/channels/another.one"},
		  "messages":[{"$ref":"#/channels/another.one/messages/Asyncapi300TestMyAnotherMessage"}]
		},
		"send:one.{name}.two":{
		  "action":"send",
		  "channel":{"$ref":"#/channels/one.{name}.two"},
		  "messages":[{"$ref":"#/channels/one.{name}.two/messages/Asyncapi300TestMyMessage"}]
		}
	  },
	  "components":{
		"schemas":{
		  "Asyncapi300TestMyAnotherMessage":{
			"properties":{
			  "item":{
				"$ref":"#/components/schemas/Asyncapi300TestSubItem",
				"description":"Some item"
			  }
			},
			"type":"object"
		  },
		  "Asyncapi300TestMyMessage":{
			"properties":{
			  "createdAt":{"description":"Creation time","format":"date-time","type":"string"},
			  "items":{
				"description":"List of items",
				"items":{"$ref":"#/components/schemas/Asyncapi300TestSubItem"},
				"type":["array","null"]
			  }
			},
			"type":"object"
		  },
		  "Asyncapi300TestSubItem":{
			"properties":{
			  "key":{"description":"Item key","type":"string"},
			  "values":{
				"description":"List of item values","items":{"type":"integer"},
				"type":["array","null"],"uniqueItems":true
			  }
			},
			"type":"object"
		  }
		}
	  }
	}`), r.Schema)
}
