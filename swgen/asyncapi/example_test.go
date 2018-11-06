package asyncapi_test

import (
	"fmt"
	"time"

	"github.com/swaggest/go-asyncapi/spec"
	"github.com/swaggest/go-asyncapi/swgen/asyncapi"
)

func ExampleGenerator_GenDocument() {
	type SubItem struct {
		Key    string  `json:"key"`
		Values []int64 `json:"values" uniqueItems:"true"`
	}

	type MyMessage struct {
		Name      string    `path:"name"`
		CreatedAt time.Time `json:"createdAt"`
		Items     []SubItem `json:"items"`
	}

	type MyAnotherMessage struct {
		TraceID string  `header:"X-Trace-ID"`
		Item    SubItem `json:"item"`
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
	// output:
	// asyncapi: 1.2.0
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
	//           format: date-time
	//           type: string
	//         items:
	//           items:
	//             $ref: '#/components/schemas/SubItem'
	//           type: array
	//       type: object
	//     SubItem:
	//       properties:
	//         key:
	//           type: string
	//         values:
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
	//       description: This is another sample schema
	//       headers:
	//         properties:
	//           X-Trace-ID:
	//             type: string
	//         required:
	//         - X-Trace-ID
	//         type: object
	//       payload:
	//         $ref: '#/components/schemas/MyAnotherMessage'
	//       summary: Sample consumer
	//   one.{name}.two:
	//     parameters:
	//     - name: name
	//       schema:
	//         type: string
	//     publish:
	//       description: This is a sample schema
	//       payload:
	//         $ref: '#/components/schemas/MyMessage'
	//       summary: Sample publisher
}
