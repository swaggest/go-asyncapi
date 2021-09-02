package asyncapi_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/swaggest/assertjson"
	"github.com/swaggest/go-asyncapi/reflector/asyncapi-2.0.0"
	"github.com/swaggest/go-asyncapi/spec-2.0.0"
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

	r := asyncapi.Reflector{
		Schema: &spec.AsyncAPI{
			Servers: map[string]spec.Server{
				"production": {
					URL:             "api.lovely.com:{port}",
					Protocol:        "amqp",
					ProtocolVersion: "AMQP 0.9.1",
				},
			},
			Info: spec.Info{
				Version: "1.2.3", // required
				Title:   "My Lovely Messaging API",
			},
		},
	}
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

	j, err := json.MarshalIndent(r.Schema, "", " ")
	require.NoError(t, err)

	assertjson.Equal(t, []byte(`{
	 "asyncapi": "2.0.0",
	 "info": {
	  "title": "My Lovely Messaging API",
	  "version": "1.2.3"
	 },
	 "servers": {
	  "production": {
	   "url": "api.lovely.com:{port}",
	   "protocol": "amqp",
	   "protocolVersion": "AMQP 0.9.1"
	  }
	 },
	 "channels": {
	  "another.one": {
	   "subscribe": {
		"message": {
		 "$ref": "#/components/messages/Asyncapi200TestMyAnotherMessage"
		}
	   }
	  },
	  "one.{name}.two": {
	   "parameters": {
		"name": {
		 "schema": {
		  "description": "Name",
		  "type": "string"
		 }
		}
	   },
	   "publish": {
		"message": {
		 "$ref": "#/components/messages/Asyncapi200TestMyMessage"
		}
	   }
	  }
	 },
	 "components": {
	  "schemas": {
	   "Asyncapi200TestMyAnotherMessage": {
		"properties": {
		 "item": {
		  "$ref": "#/components/schemas/Asyncapi200TestSubItem",
		  "description": "Some item"
		 }
		},
		"type": "object"
	   },
	   "Asyncapi200TestMyMessage": {
		"properties": {
		 "createdAt": {
		  "description": "Creation time",
		  "format": "date-time",
		  "type": "string"
		 },
		 "items": {
		  "description": "List of items",
		  "items": {
		   "$ref": "#/components/schemas/Asyncapi200TestSubItem"
		  },
		  "type": [
		   "array",
		   "null"
		  ]
		 }
		},
		"type": "object"
	   },
	   "Asyncapi200TestSubItem": {
		"properties": {
		 "key": {
		  "description": "Item key",
		  "type": "string"
		 },
		 "values": {
		  "description": "List of item values",
		  "items": {
		   "type": "integer"
		  },
		  "type": [
		   "array",
		   "null"
		  ],
		  "uniqueItems": true
		 }
		},
		"type": "object"
	   }
	  },
	  "messages": {
	   "Asyncapi200TestMyAnotherMessage": {
		"headers": {
		 "properties": {
		  "X-Trace-ID": {
		   "description": "Tracing header",
		   "type": "string"
		  }
		 },
		 "required": [
		  "X-Trace-ID"
		 ],
		 "type": "object"
		},
		"payload": {
		 "$ref": "#/components/schemas/Asyncapi200TestMyAnotherMessage"
		},
		"summary": "Sample consumer",
		"description": "This is another sample schema"
	   },
	   "Asyncapi200TestMyMessage": {
		"payload": {
		 "$ref": "#/components/schemas/Asyncapi200TestMyMessage"
		},
		"summary": "Sample publisher",
		"description": "This is a sample schema"
	   }
	  }
	 }
	}`), j, string(j))
}
