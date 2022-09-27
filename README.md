# AsyncAPI Generator for Go

[![Build Status](https://github.com/swaggest/go-asyncapi/workflows/test-unit/badge.svg)](https://github.com/swaggest/go-asyncapi/actions?query=branch%3Amaster+workflow%3Atest-unit)
[![Coverage Status](https://codecov.io/gh/swaggest/go-asyncapi/branch/master/graph/badge.svg)](https://codecov.io/gh/swaggest/go-asyncapi)
[![GoDoc](https://godoc.org/github.com/swaggest/go-asyncapi?status.svg)](https://godoc.org/github.com/swaggest/go-asyncapi)
![Code lines](https://sloc.xyz/github/swaggest/go-asyncapi/?category=code)
![Comments](https://sloc.xyz/github/swaggest/go-asyncapi/?category=comments)

This library helps to create [AsyncAPI](https://www.asyncapi.com/) spec from your Go message structures.

Supported AsyncAPI versions:
* `v2.4.0` 
* `v2.1.0` 
* `v2.0.0`
* `v1.2.0`

## Example

```go
package asyncapi_test

import (
	"fmt"
	"os"
	"time"

	"github.com/swaggest/go-asyncapi/reflector/asyncapi-2.4.0"
	"github.com/swaggest/go-asyncapi/spec-2.4.0"
)

func main() {
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
	asyncAPI.Info.Version = "1.2.3"
	asyncAPI.Info.Title = "My Lovely Messaging API"

	asyncAPI.AddServer("live", spec.Server{
		URL:             "api.{country}.lovely.com:5672",
		Description:     "Production instance.",
		ProtocolVersion: "0.9.1",
		Protocol:        "amqp",
		Variables: map[string]spec.ServerVariable{
			"country": {
				Enum:        []string{"RU", "US", "DE", "FR"},
				Default:     "US",
				Description: "Country code.",
			},
		},
	})

	reflector := asyncapi.Reflector{}
	reflector.Schema = &asyncAPI

	mustNotFail := func(err error) {
		if err != nil {
			panic(err.Error())
		}
	}

	mustNotFail(reflector.AddChannel(asyncapi.ChannelInfo{
		Name: "one.{name}.two",
		BaseChannelItem: &spec.ChannelItem{
			Bindings: &spec.ChannelBindingsObject{
				Amqp: &spec.AmqpChannel{
					Is: spec.AmqpChannelIsRoutingKey,
					Exchange: &spec.AmqpChannelExchange{
						Name: "some-exchange",
					},
				},
			},
		},
		Publish: &asyncapi.MessageSample{
			MessageEntity: spec.MessageEntity{
				Description: "This is a sample schema.",
				Summary:     "Sample publisher",
			},
			MessageSample: new(MyMessage),
		},
	}))

	mustNotFail(reflector.AddChannel(asyncapi.ChannelInfo{
		Name: "another.one",
		Subscribe: &asyncapi.MessageSample{
			MessageEntity: spec.MessageEntity{
				Description: "This is another sample schema.",
				Summary:     "Sample consumer",
			},
			MessageSample: new(MyAnotherMessage),
		},
	}))

	yaml, err := reflector.Schema.MarshalYAML()
	mustNotFail(err)

	fmt.Println(string(yaml))
	mustNotFail(os.WriteFile("sample.yaml", yaml, 0o600))
	// output:
	// asyncapi: 2.4.0
	// info:
	//   title: My Lovely Messaging API
	//   version: 1.2.3
	// servers:
	//   live:
	//     url: api.{country}.lovely.com:5672
	//     description: Production instance.
	//     protocol: amqp
	//     protocolVersion: 0.9.1
	//     variables:
	//       country:
	//         enum:
	//         - RU
	//         - US
	//         - DE
	//         - FR
	//         default: US
	//         description: Country code.
	// channels:
	//   another.one:
	//     subscribe:
	//       message:
	//         $ref: '#/components/messages/Asyncapi240TestMyAnotherMessage'
	//   one.{name}.two:
	//     parameters:
	//       name:
	//         schema:
	//           description: Name
	//           type: string
	//     publish:
	//       message:
	//         $ref: '#/components/messages/Asyncapi240TestMyMessage'
	//     bindings:
	//       amqp:
	//         bindingVersion: 0.2.0
	//         is: routingKey
	//         exchange:
	//           name: some-exchange
	// components:
	//   schemas:
	//     Asyncapi240TestMyAnotherMessage:
	//       properties:
	//         item:
	//           $ref: '#/components/schemas/Asyncapi240TestSubItem'
	//           description: Some item
	//       type: object
	//     Asyncapi240TestMyMessage:
	//       properties:
	//         createdAt:
	//           description: Creation time
	//           format: date-time
	//           type: string
	//         items:
	//           description: List of items
	//           items:
	//             $ref: '#/components/schemas/Asyncapi240TestSubItem'
	//           type:
	//           - array
	//           - "null"
	//       type: object
	//     Asyncapi240TestSubItem:
	//       properties:
	//         key:
	//           description: Item key
	//           type: string
	//         values:
	//           description: List of item values
	//           items:
	//             type: integer
	//           type:
	//           - array
	//           - "null"
	//           uniqueItems: true
	//       type: object
	//   messages:
	//     Asyncapi240TestMyAnotherMessage:
	//       headers:
	//         properties:
	//           X-Trace-ID:
	//             description: Tracing header
	//             type: string
	//         required:
	//         - X-Trace-ID
	//         type: object
	//       payload:
	//         $ref: '#/components/schemas/Asyncapi240TestMyAnotherMessage'
	//       summary: Sample consumer
	//       description: This is another sample schema.
	//     Asyncapi240TestMyMessage:
	//       payload:
	//         $ref: '#/components/schemas/Asyncapi240TestMyMessage'
	//       summary: Sample publisher
	//       description: This is a sample schema.
}
```
