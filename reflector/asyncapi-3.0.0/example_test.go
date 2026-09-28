package asyncapi_test

import (
	"fmt"
	"os"
	"time"

	"github.com/swaggest/go-asyncapi/reflector/asyncapi-3.0.0"
	spec "github.com/swaggest/go-asyncapi/spec-3.0.0"
)

func ExampleReflector_AddChannel_amqp() {
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
	mustNotFail(os.WriteFile("sample-amqp.yaml", yaml, 0o600))
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

func ExampleReflector_AddChannel_kafka() {
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
		srv.ProtocolVersion = "1.0.0"
		srv.Protocol = "kafka"
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
		kch := ch.BindingsEns().ChannelBindingsObjectEns().KafkaEns().BindingsKafka050ChannelEns()
		kch.Topic = "one.{name}.two"
	})

	mustNotFail(reflector.AddOperation(chRef, spec.OperationActionSend, func(message *asyncapi.MessageSample) {
		message.Sample = new(MyMessage)

		message.Title = "Sample publisher"
		message.Description = "This is a sample schema."

		km := message.BindingsEns().MessageBindingsObjectEns().KafkaEns().BindingsKafka050MessageEns()
		km.KeyEns().WithSchema(map[string]interface{}{"const": "my-key"})
	}, func(op *spec.Operation) {
		op.BindingsEns().OperationBindingsObjectEns().KafkaEns().BindingsKafka050OperationEns().WithGroupIDItem("const", "my-group-id")
	}))

	anotherRef := reflector.AddChannel("another.one")

	mustNotFail(reflector.AddOperation(anotherRef, spec.OperationActionReceive, func(message *asyncapi.MessageSample) {
		message.Sample = new(MyAnotherMessage)

		message.Title = "Sample consumer"
		message.Description = "This is another sample schema."
	}, func(op *spec.Operation) {
		op.BindingsEns().OperationBindingsObjectEns().KafkaEns().BindingsKafka050OperationEns().WithGroupIDItem("const", "my-group-id-2")
	}))

	yaml, err := reflector.Schema.MarshalYAML()
	mustNotFail(err)

	fmt.Println(string(yaml))
	mustNotFail(os.WriteFile("sample-kafka.yaml", yaml, 0o600))
	// output:
	// asyncapi: 3.0.0
	// info:
	//   title: My Lovely Messaging API
	//   version: 1.2.3
	// servers:
	//   live:
	//     host: api.{country}.lovely.com:5672
	//     description: Production instance.
	//     protocol: kafka
	//     protocolVersion: 1.0.0
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
	//   one.{name}.two:
	//     address: one.{name}.two
	//     messages:
	//       Asyncapi300TestMyMessage:
	//         payload:
	//           $ref: '#/components/schemas/Asyncapi300TestMyMessage'
	//         title: Sample publisher
	//         description: This is a sample schema.
	//         bindings:
	//           kafka:
	//             bindingVersion: 0.5.0
	//             key:
	//               const: my-key
	//     parameters:
	//       name:
	//         description: Name
	//     bindings:
	//       kafka:
	//         bindingVersion: 0.5.0
	//         topic: one.{name}.two
	// operations:
	//   receive:another.one:
	//     action: receive
	//     channel:
	//       $ref: '#/channels/another.one'
	//     messages:
	//     - $ref: '#/channels/another.one/messages/Asyncapi300TestMyAnotherMessage'
	//     bindings:
	//       kafka:
	//         bindingVersion: 0.5.0
	//         groupId:
	//           const: my-group-id-2
	//   send:one.{name}.two:
	//     action: send
	//     channel:
	//       $ref: '#/channels/one.{name}.two'
	//     messages:
	//     - $ref: '#/channels/one.{name}.two/messages/Asyncapi300TestMyMessage'
	//     bindings:
	//       kafka:
	//         bindingVersion: 0.5.0
	//         groupId:
	//           const: my-group-id
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

func ExampleReflector_AddChannel_sqlite() {
	type Job struct {
		ID   int64  `json:"id" description:"Row id"`
		Body string `json:"body" description:"Job payload"`
	}

	asyncAPI := spec.AsyncAPI{}

	asyncAPI.UpdateInfo(func(i *spec.Info) {
		i.Version = "1.0.0"
		i.Title = "My Private SQLite Queue"
	})

	// "sqlite" is not an officially registered AsyncAPI protocol, but the
	// server protocol field is a free string, so it can still be declared here.
	asyncAPI.AddServer("local", func(srv *spec.Server) {
		srv.Host = "file:./queue.db"
		srv.Protocol = "sqlite"
		srv.Description = "Local SQLite-backed queue."
	})

	reflector := asyncapi.Reflector{}
	reflector.Schema = &asyncAPI

	mustNotFail := func(err error) {
		if err != nil {
			panic(err.Error())
		}
	}

	// There is no first-class "sqlite" entry in the Channel Bindings Object
	// (it's a closed set of known protocols), so custom binding details are
	// attached via the sanctioned x- specification extension mechanism instead.
	chRef := reflector.AddChannel("jobs", func(ch *spec.Channel) {
		ch.BindingsEns().ChannelBindingsObjectEns().WithMapOfAnythingItem("x-sqlite", map[string]interface{}{
			"table":       "jobs_queue",
			"pollingMode": "poll",
		})
	})

	mustNotFail(reflector.AddOperation(chRef, spec.OperationActionSend, func(message *asyncapi.MessageSample) {
		message.Sample = new(Job)
		message.Title = "Job"
		message.Description = "A job row inserted into the queue table."
	}))

	yaml, err := reflector.Schema.MarshalYAML()
	mustNotFail(err)

	fmt.Println(string(yaml))
	mustNotFail(os.WriteFile("sample-sqlite.yaml", yaml, 0o600))
	// output:
	// asyncapi: 3.0.0
	// info:
	//   title: My Private SQLite Queue
	//   version: 1.0.0
	// servers:
	//   local:
	//     host: file:./queue.db
	//     description: Local SQLite-backed queue.
	//     protocol: sqlite
	// channels:
	//   jobs:
	//     address: jobs
	//     messages:
	//       Asyncapi300TestJob:
	//         payload:
	//           $ref: '#/components/schemas/Asyncapi300TestJob'
	//         title: Job
	//         description: A job row inserted into the queue table.
	//     bindings:
	//       x-sqlite:
	//         pollingMode: poll
	//         table: jobs_queue
	// operations:
	//   send:jobs:
	//     action: send
	//     channel:
	//       $ref: '#/channels/jobs'
	//     messages:
	//     - $ref: '#/channels/jobs/messages/Asyncapi300TestJob'
	// components:
	//   schemas:
	//     Asyncapi300TestJob:
	//       properties:
	//         body:
	//           description: Job payload
	//           type: string
	//         id:
	//           description: Row id
	//           type: integer
	//       type: object
}
