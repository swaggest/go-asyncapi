# AsyncAPI Generator for Go

This library helps to create [AsyncAPI](https://www.asyncapi.com/) spec from your Go message structures.

## Example

```go
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
					URL:    "api.streetlights.smartylighting.com:{port}",
					Scheme: spec.ServerSchemeAmqp,
				},
			},
			BaseTopic: "smartylighting.streetlights.1.0",
			Info: &spec.Info{
				Title: "My Lovely Messaging API",
			},
		},
	}
	err := g.AddTopic(asyncapi.TopicInfo{
		Topic: "one.{name}.two",
		Publish: &asyncapi.Message{
			Message: spec.Message{
				Description: "This is a sample schema",
				Summary:     "Sample publisher",
			},
			MessageSample: new(MyMessage),
		},
	})
	assert.NoError(t, err)

	err = g.AddTopic(asyncapi.TopicInfo{
		Topic: "another.one",
		Subscribe: &asyncapi.Message{
			Message: spec.Message{
				Description: "This is another sample schema",
				Summary:     "Sample consumer",
			},
			MessageSample: new(MyAnotherMessage),
		},
	})
	assert.NoError(t, err)

	yaml, err := g.Data.MarshalYAML()
	assert.NoError(t, err)
	//nolint:errcheck
	ioutil.WriteFile("sample.yaml", yaml, 0666)
```