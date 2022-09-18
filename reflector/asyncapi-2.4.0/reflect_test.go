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

	assertjson.EqualMarshal(t, []byte(`{
	 "asyncapi": "2.1.0",
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
		 "$ref": "#/components/messages/Asyncapi210TestMyAnotherMessage"
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
		 "$ref": "#/components/messages/Asyncapi210TestMyMessage"
		}
	   }
	  }
	 },
	 "components": {
	  "schemas": {
	   "Asyncapi210TestMyAnotherMessage": {
		"properties": {
		 "item": {
		  "$ref": "#/components/schemas/Asyncapi210TestSubItem",
		  "description": "Some item"
		 }
		},
		"type": "object"
	   },
	   "Asyncapi210TestMyMessage": {
		"properties": {
		 "createdAt": {
		  "description": "Creation time",
		  "format": "date-time",
		  "type": "string"
		 },
		 "items": {
		  "description": "List of items",
		  "items": {
		   "$ref": "#/components/schemas/Asyncapi210TestSubItem"
		  },
		  "type": [
		   "array",
		   "null"
		  ]
		 }
		},
		"type": "object"
	   },
	   "Asyncapi210TestSubItem": {
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
	   "Asyncapi210TestMyAnotherMessage": {
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
		 "$ref": "#/components/schemas/Asyncapi210TestMyAnotherMessage"
		},
		"summary": "Sample consumer",
		"description": "This is another sample schema"
	   },
	   "Asyncapi210TestMyMessage": {
		"payload": {
		 "$ref": "#/components/schemas/Asyncapi210TestMyMessage"
		},
		"summary": "Sample publisher",
		"description": "This is a sample schema"
	   }
	  }
	 }
	}
	`), r.Schema)
}
