package asyncapi_test

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/swaggest/go-asyncapi/reflector/asyncapi-2.1.0"
	"github.com/swaggest/go-asyncapi/spec-2.1.0"
)

func ExampleReflector_AddChannel() {
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

	reflector := asyncapi.Reflector{
		Schema: &spec.AsyncAPI{
			Servers: map[string]spec.Server{
				"live": {
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
				},
			},
			Info: spec.Info{
				Version: "1.2.3", // required
				Title:   "My Lovely Messaging API",
			},
		},
	}
	mustNotFail := func(err error) {
		if err != nil {
			panic(err.Error())
		}
	}

	mustNotFail(reflector.AddChannel(asyncapi.ChannelInfo{
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
	mustNotFail(ioutil.WriteFile("sample.yaml", yaml, 0o600))
	// output:
	// asyncapi: 2.1.0
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
	//         $ref: '#/components/messages/Asyncapi210TestMyAnotherMessage'
	//   one.{name}.two:
	//     parameters:
	//       name:
	//         schema:
	//           description: Name
	//           type: string
	//     publish:
	//       message:
	//         $ref: '#/components/messages/Asyncapi210TestMyMessage'
	//     bindings:
	//       amqp:
	//         is: routingKey
	//         exchange:
	//           name: some-exchange
	// components:
	//   schemas:
	//     Asyncapi210TestMyAnotherMessage:
	//       properties:
	//         item:
	//           $ref: '#/components/schemas/Asyncapi210TestSubItem'
	//           description: Some item
	//       type: object
	//     Asyncapi210TestMyMessage:
	//       properties:
	//         createdAt:
	//           description: Creation time
	//           format: date-time
	//           type: string
	//         items:
	//           description: List of items
	//           items:
	//             $ref: '#/components/schemas/Asyncapi210TestSubItem'
	//           type:
	//           - array
	//           - "null"
	//       type: object
	//     Asyncapi210TestSubItem:
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
	//     Asyncapi210TestMyAnotherMessage:
	//       headers:
	//         properties:
	//           X-Trace-ID:
	//             description: Tracing header
	//             type: string
	//         required:
	//         - X-Trace-ID
	//         type: object
	//       payload:
	//         $ref: '#/components/schemas/Asyncapi210TestMyAnotherMessage'
	//       summary: Sample consumer
	//       description: This is another sample schema.
	//     Asyncapi210TestMyMessage:
	//       payload:
	//         $ref: '#/components/schemas/Asyncapi210TestMyMessage'
	//       summary: Sample publisher
	//       description: This is a sample schema.
}
