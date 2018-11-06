package asyncapi_test

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/swaggest/go-asyncapi/spec"
	"github.com/swaggest/go-asyncapi/swgen/asyncapi"
)

func ExampleGenerator_AddTopic() {
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
		TraceID string  `header:"X-Trace-ID" description:"Tracing header"`
		Item    SubItem `json:"item" description:"Some item"`
	}

	g := asyncapi.Generator{
		Data: &spec.AsyncAPI{
			Asyncapi: spec.Asyncapi120,
			Servers: []spec.Server{
				{
					URL:    "api.lovely.com:{port}",
					Scheme: spec.ServerSchemeAmqp,
				},
			},
			Info: &spec.Info{
				Version: "0.0.0", //required
				Title:   "My Lovely Messaging API",
			},
		},
	}
	must := func(err error) {
		if err != nil {
			panic(err.Error())
		}
	}
	must(g.AddTopic(asyncapi.TopicInfo{
		Topic: "one.{name}.two",
		Publish: &asyncapi.Message{
			Message: spec.Message{
				Description: "This is a sample schema",
				Summary:     "Sample publisher",
			},
			MessageSample: new(MyMessage),
		},
	}))

	must(g.AddTopic(asyncapi.TopicInfo{
		Topic: "another.one",
		Subscribe: &asyncapi.Message{
			Message: spec.Message{
				Description: "This is another sample schema",
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
	// asyncapi: 1.2.0
	// components:
	//   messages:
	//     MyAnotherMessage:
	//       description: This is another sample schema
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
	//     MyMessage:
	//       description: This is a sample schema
	//       payload:
	//         $ref: '#/components/schemas/MyMessage'
	//       summary: Sample publisher
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
	// info:
	//   title: My Lovely Messaging API
	//   version: 0.0.0
	// servers:
	// - scheme: amqp
	//   url: api.lovely.com:{port}
	// topics:
	//   another.one:
	//     subscribe:
	//       $ref: '#/components/messages/MyAnotherMessage'
	//   one.{name}.two:
	//     parameters:
	//     - description: Name
	//       name: name
	//       schema:
	//         description: Name
	//         type: string
	//     publish:
	//       $ref: '#/components/messages/MyMessage'
}
