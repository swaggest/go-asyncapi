# AsyncAPI Generator for Go

[![Build Status](https://travis-ci.org/swaggest/go-asyncapi.svg?branch=master)](https://travis-ci.org/swaggest/go-asyncapi)
[![Coverage Status](https://codecov.io/gh/swaggest/go-asyncapi/branch/master/graph/badge.svg)](https://codecov.io/gh/swaggest/go-asyncapi)
[![GoDoc](https://godoc.org/github.com/swaggest/go-asyncapi?status.svg)](https://godoc.org/github.com/swaggest/go-asyncapi)
![Code lines](https://sloc.xyz/github/swaggest/go-asyncapi/?category=code)
![Comments](https://sloc.xyz/github/swaggest/go-asyncapi/?category=comments)

This library helps to create [AsyncAPI](https://www.asyncapi.com/) spec from your Go message structures.

## Example

```go
package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/swaggest/go-asyncapi/spec-2.0.0"
	"github.com/swaggest/go-asyncapi/swgen/asyncapi-2.0.0"
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

	g := asyncapi.Generator{
		Data: spec.AsyncAPI{
			Servers: map[string]spec.Server{
				"production": {
					URL:      "api.lovely.com:{port}",
					Protocol: "amqp",
				},
			},
			Info: &spec.Info{
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
	must(g.AddChannel(asyncapi.ChannelInfo{
		Name: "one.{name}.two",
		BaseChannelItem: &spec.ChannelItem{
			Bindings: &spec.ChannelBindingsObject{
				Amqp: &spec.AMQP091ChannelBindingObject{
					Is: spec.AMQP091ChannelBindingObjectIsRoutingKey,
					Exchange: &spec.Exchange{
						Name: "some-exchange",
					},
				},
			},
		},
		Publish: &asyncapi.Message{
			MessageEntity: spec.MessageEntity{
				Description: "This is a sample schema.",
				Summary:     "Sample publisher",
			},
			MessageSample: new(MyMessage),
		},
	}))

	must(g.AddChannel(asyncapi.ChannelInfo{
		Name: "another.one",
		Subscribe: &asyncapi.Message{
			MessageEntity: spec.MessageEntity{
				Description: "This is another sample schema.",
				Summary:     "Sample consumer",
			},
			MessageSample: new(MyAnotherMessage),
		},
	}))

	yaml, err := g.Data.MarshalYAML()
	must(err)

	fmt.Println(string(yaml))
	must(ioutil.WriteFile("sample.yaml", yaml, 0644))
	// output:
	// asyncapi: 2.0.0
	// info:
	//   title: My Lovely Messaging API
	//   version: 1.2.3
	// servers:
	//   production:
	//     url: api.lovely.com:{port}
	//     protocol: amqp
	// channels:
	//   another.one:
	//     subscribe:
	//       message:
	//         $ref: '#/components/messages/MyAnotherMessage'
	//   one.{name}.two:
	//     parameters:
	//       name:
	//         description: Name
	//         schema:
	//           description: Name
	//           type: string
	//     publish:
	//       message:
	//         $ref: '#/components/messages/MyMessage'
	//     bindings:
	//       amqp:
	//         is: routingKey
	//         exchange:
	//           name: some-exchange
	// components:
	//   schemas:
	//     MyAnotherMessage:
	//       properties:
	//         item:
	//           $ref: '#/components/schemas/SubItem'
	//       type: object
	//     MyMessage:
	//       properties:
	//         createdAt:
	//           description: Creation time
	//           format: date-time
	//           type: string
	//         items:
	//           description: List of items
	//           items:
	//             $ref: '#/components/schemas/SubItem'
	//           type: array
	//       type: object
	//     SubItem:
	//       properties:
	//         key:
	//           description: Item key
	//           type: string
	//         values:
	//           description: List of item values
	//           items:
	//             format: int64
	//             type: integer
	//           type: array
	//           uniqueItems: true
	//       type: object
	//   messages:
	//     MyAnotherMessage:
	//       headers:
	//         properties:
	//           X-Trace-ID:
	//             description: Tracing header
	//             type: string
	//         required:
	//         - X-Trace-ID
	//         type: object
	//       payload:
	//         $ref: '#/components/schemas/MyAnotherMessage'
	//       summary: Sample consumer
	//       description: This is another sample schema.
	//     MyMessage:
	//       payload:
	//         $ref: '#/components/schemas/MyMessage'
	//       summary: Sample publisher
	//       description: This is a sample schema.
}
```
