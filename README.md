# AsyncAPI Generator for Go

[![Build Status](https://github.com/swaggest/go-asyncapi/workflows/test-unit/badge.svg)](https://github.com/swaggest/go-asyncapi/actions?query=branch%3Amaster+workflow%3Atest-unit)
[![Coverage Status](https://codecov.io/gh/swaggest/go-asyncapi/branch/master/graph/badge.svg)](https://codecov.io/gh/swaggest/go-asyncapi)
[![GoDoc](https://godoc.org/github.com/swaggest/go-asyncapi?status.svg)](https://godoc.org/github.com/swaggest/go-asyncapi)
![Code lines](https://sloc.xyz/github/swaggest/go-asyncapi/?category=code)
![Comments](https://sloc.xyz/github/swaggest/go-asyncapi/?category=comments)

This library helps to create [AsyncAPI](https://www.asyncapi.com/) spec from your Go message structures.

Supported AsyncAPI versions:
* `v3.0.0` 
* `v2.4.0` 
* `v2.1.0` 
* `v2.0.0`
* `v1.2.0`

## Example

```go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/swaggest/go-asyncapi/reflector/asyncapi-3.0.0"
	"github.com/swaggest/go-asyncapi/spec-3.0.0"
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

	asyncAPI.UpdateInfo(func(i *spec.Info) {
		i.Version = "1.2.3"
		i.Title = "My Lovely Messaging API"
	})

	asyncAPI.AddServer("live", func(srv *spec.Server) {
		srv.Description = "Production instance."
		srv.Host = "api.{country}.lovely.com:5672"
		srv.ProtocolVersion = "0.9.1"
		srv.Protocol = "amqp"
		srv.AddVariable("country", func(v *spec.ServerVariable) {
			v.WithEnum("RU", "US", "DE", "FR")
			v.Default = "US"
			v.Description = "Country code."
		})
	})

	reflector := asyncapi.Reflector{}
	reflector.Schema = &asyncAPI

	mustNotFail := func(err error) {
		if err != nil {
			panic(err.Error())
		}
	}

	chRef := reflector.AddChannel("one.{name}.two", func(ch *spec.Channel) {
		ach := ch.BindingsEns().ChannelBindingsObjectEns().AmqpEns()
		// is: routingKey pairs with "exchange" (not "queue"), per the AMQP channel binding spec.
		ach.Is = spec.AMQPChannelBindingsObjectIsRoutingKey
		ach.ExchangeEns().Name = "some-exchange"
	})

	mustNotFail(reflector.AddOperation(chRef, spec.OperationActionSend, func(message *asyncapi.MessageSample) {
		message.Sample = new(MyMessage)

		message.Title = "Sample publisher"
		message.Description = "This is a sample schema."
	}))

	anotherRef := reflector.AddChannel("another.one", func(ch *spec.Channel) {
		ach := ch.BindingsEns().ChannelBindingsObjectEns().AmqpEns()
		// is: queue pairs with "queue" (not "exchange"), per the AMQP channel binding spec.
		ach.Is = spec.AMQPChannelBindingsObjectIsQueue
		ach.QueueEns().Name = "some-queue"
	})

	mustNotFail(reflector.AddOperation(anotherRef, spec.OperationActionReceive, func(message *asyncapi.MessageSample) {
		message.Sample = new(MyAnotherMessage)

		message.Title = "Sample consumer"
		message.Description = "This is another sample schema."
	}))

	yaml, err := reflector.Schema.MarshalYAML()
	mustNotFail(err)

	fmt.Println(string(yaml))
	mustNotFail(os.WriteFile("sample.yaml", yaml, 0o600))
	// output:
	// asyncapi: 3.0.0
	// info:
	//   title: My Lovely Messaging API
	//   version: 1.2.3
	// servers:
	//   live:
	//     host: api.{country}.lovely.com:5672
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
	//     address: another.one
	//     messages:
	//       Asyncapi300TestMyAnotherMessage:
	//         headers:
	//           properties:
	//             X-Trace-ID:
	//               description: Tracing header
	//               type: string
	//           required:
	//           - X-Trace-ID
	//           type: object
	//         payload:
	//           $ref: '#/components/schemas/Asyncapi300TestMyAnotherMessage'
	//         title: Sample consumer
	//         description: This is another sample schema.
	//     bindings:
	//       amqp:
	//         bindingVersion: 0.3.0
	//         is: queue
	//         queue:
	//           name: some-queue
	//   one.{name}.two:
	//     address: one.{name}.two
	//     messages:
	//       Asyncapi300TestMyMessage:
	//         payload:
	//           $ref: '#/components/schemas/Asyncapi300TestMyMessage'
	//         title: Sample publisher
	//         description: This is a sample schema.
	//     parameters:
	//       name:
	//         description: Name
	//     bindings:
	//       amqp:
	//         bindingVersion: 0.3.0
	//         is: routingKey
	//         exchange:
	//           name: some-exchange
	// operations:
	//   receive:another.one:
	//     action: receive
	//     channel:
	//       $ref: '#/channels/another.one'
	//     messages:
	//     - $ref: '#/channels/another.one/messages/Asyncapi300TestMyAnotherMessage'
	//   send:one.{name}.two:
	//     action: send
	//     channel:
	//       $ref: '#/channels/one.{name}.two'
	//     messages:
	//     - $ref: '#/channels/one.{name}.two/messages/Asyncapi300TestMyMessage'
	// components:
	//   schemas:
	//     Asyncapi300TestMyAnotherMessage:
	//       properties:
	//         item:
	//           $ref: '#/components/schemas/Asyncapi300TestSubItem'
	//           description: Some item
	//       type: object
	//     Asyncapi300TestMyMessage:
	//       properties:
	//         createdAt:
	//           description: Creation time
	//           format: date-time
	//           type: string
	//         items:
	//           description: List of items
	//           items:
	//             $ref: '#/components/schemas/Asyncapi300TestSubItem'
	//           type:
	//           - array
	//           - "null"
	//       type: object
	//     Asyncapi300TestSubItem:
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
}
```
