package spec_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/swaggest/go-asyncapi/spec"
)

func TestInfo_MarshalJSON(t *testing.T) {
	i := spec.Info{
		Version: "v1",
		MapOfAnythingValues: map[string]interface{}{
			"x-two": "two",
			"x-one": 1,
		},
	}

	res, err := json.Marshal(i)
	assert.NoError(t, err)
	assert.Equal(t, `{"version":"v1","x-one":1,"x-two":"two"}`, string(res))
}

func TestInfo_MarshalJSON_Nil(t *testing.T) {
	i := spec.Info{
		Version: "v1",
	}

	res, err := json.Marshal(i)
	assert.NoError(t, err)
	assert.Equal(t, `{"version":"v1"}`, string(res))
}

func TestInfo_UnmarshalJSON(t *testing.T) {
	i := spec.Info{}

	err := json.Unmarshal([]byte(`{"version":"v1","x-one":1,"x-two":"two"}`), &i)
	assert.NoError(t, err)
	assert.Equal(t, 1.0, i.MapOfAnythingValues["x-one"].(float64))
	assert.Equal(t, "two", i.MapOfAnythingValues["x-two"])
}

func TestAPIKey_UnmarshalJSON(t *testing.T) {
	data := []byte(`asyncapi: '1.2.0'
info:
  title: Streetlights API
  version: '1.0.0'
  description: |
    The Smartylighting Streetlights API allows you to remotely manage the city lights.

    ### Check out its awesome features:

    * Turn a specific streetlight on/off 🌃
    * Dim a specific streetlight 😎
    * Receive real-time information about environmental lighting conditions 📈
  license:
    name: Apache 2.0
    url: https://www.apache.org/licenses/LICENSE-2.0
baseTopic: smartylighting.streetlights.1.0

servers:
- url: api.streetlights.smartylighting.com:{port}
  scheme: mqtt
  description: Test broker
  variables:
    port:
      description: Secure connection (TLS) is available through port 8883.
      default: '1883'
      enum:
      - '1883'
      - '8883'

security:
- apiKey: []

topics:
  event.{streetlightId}.lighting.measured:
    parameters:
    - $ref: '#/components/parameters/streetlightId'
    publish:
      $ref: '#/components/messages/lightMeasured'

  action.{streetlightId}.turn.on:
    parameters:
    - $ref: '#/components/parameters/streetlightId'
    subscribe:
      $ref: '#/components/messages/turnOnOff'

  action.{streetlightId}.turn.off:
    parameters:
    - $ref: '#/components/parameters/streetlightId'
    subscribe:
      $ref: '#/components/messages/turnOnOff'

  action.{streetlightId}.dim:
    parameters:
    - $ref: '#/components/parameters/streetlightId'
    subscribe:
      $ref: '#/components/messages/dimLight'

components:
  messages:
    lightMeasured:
      summary: Inform about environmental lighting conditions for a particular streetlight.
      payload:
        $ref: "#/components/schemas/lightMeasuredPayload"
    turnOnOff:
      summary: Command a particular streetlight to turn the lights on or off.
      payload:
        $ref: "#/components/schemas/turnOnOffPayload"
    dimLight:
      summary: Command a particular streetlight to dim the lights.
      payload:
        $ref: "#/components/schemas/dimLightPayload"

  schemas:
    lightMeasuredPayload:
      type: object
      properties:
        lumens:
          type: integer
          minimum: 0
          description: Light intensity measured in lumens.
        sentAt:
          $ref: "#/components/schemas/sentAt"
    turnOnOffPayload:
      type: object
      properties:
        command:
          type: string
          enum:
          - on
          - off
          description: Whether to turn on or off the light.
        sentAt:
          $ref: "#/components/schemas/sentAt"
    dimLightPayload:
      type: object
      properties:
        percentage:
          type: integer
          description: Percentage to which the light should be dimmed to.
          minimum: 0
          maximum: 100
        sentAt:
          $ref: "#/components/schemas/sentAt"
    sentAt:
      type: string
      format: date-time
      description: Date and time when the message was sent.

  securitySchemes:
    apiKey:
      type: apiKey
      in: user
      description: Provide your API key as the user and leave the password empty.

  parameters:
    streetlightId:
      name: streetlightId
      description: The ID of the streetlight.
      schema:
        type: string

`)

	var a spec.AsyncAPI
	err := a.UnmarshalYAML(data)
	assert.NoError(t, err)

	assert.Equal(t, "#/components/messages/lightMeasured", a.Topics.MapOfTopicItemValues["event.{streetlightId}.lighting.measured"].Publish.Message.Ref)

	data, err = json.MarshalIndent(a, "", " ")
	assert.NoError(t, err)
	expected := `{
 "asyncapi": "1.2.0",
 "info": {
  "title": "Streetlights API",
  "version": "1.0.0",
  "description": "The Smartylighting Streetlights API allows you to remotely manage the city lights.\n\n### Check out its awesome features:\n\n* Turn a specific streetlight on/off 🌃\n* Dim a specific streetlight 😎\n* Receive real-time information about environmental lighting conditions 📈\n",
  "license": {
   "name": "Apache 2.0",
   "url": "https://www.apache.org/licenses/LICENSE-2.0"
  }
 },
 "baseTopic": "smartylighting.streetlights.1.0",
 "servers": [
  {
   "url": "api.streetlights.smartylighting.com:{port}",
   "description": "Test broker",
   "scheme": "mqtt",
   "variables": {
    "port": {
     "enum": [
      "1883",
      "8883"
     ],
     "default": "1883",
     "description": "Secure connection (TLS) is available through port 8883."
    }
   }
  }
 ],
 "topics": {
  "action.{streetlightId}.dim": {
   "parameters": [
    {
     "$ref": "#/components/parameters/streetlightId"
    }
   ],
   "subscribe": {
    "$ref": "#/components/messages/dimLight"
   }
  },
  "action.{streetlightId}.turn.off": {
   "parameters": [
    {
     "$ref": "#/components/parameters/streetlightId"
    }
   ],
   "subscribe": {
    "$ref": "#/components/messages/turnOnOff"
   }
  },
  "action.{streetlightId}.turn.on": {
   "parameters": [
    {
     "$ref": "#/components/parameters/streetlightId"
    }
   ],
   "subscribe": {
    "$ref": "#/components/messages/turnOnOff"
   }
  },
  "event.{streetlightId}.lighting.measured": {
   "parameters": [
    {
     "$ref": "#/components/parameters/streetlightId"
    }
   ],
   "publish": {
    "$ref": "#/components/messages/lightMeasured"
   }
  }
 },
 "components": {
  "schemas": {
   "dimLightPayload": {
    "properties": {
     "percentage": {
      "description": "Percentage to which the light should be dimmed to.",
      "maximum": 100,
      "minimum": 0,
      "type": "integer"
     },
     "sentAt": {
      "$ref": "#/components/schemas/sentAt"
     }
    },
    "type": "object"
   },
   "lightMeasuredPayload": {
    "properties": {
     "lumens": {
      "description": "Light intensity measured in lumens.",
      "minimum": 0,
      "type": "integer"
     },
     "sentAt": {
      "$ref": "#/components/schemas/sentAt"
     }
    },
    "type": "object"
   },
   "sentAt": {
    "description": "Date and time when the message was sent.",
    "format": "date-time",
    "type": "string"
   },
   "turnOnOffPayload": {
    "properties": {
     "command": {
      "description": "Whether to turn on or off the light.",
      "enum": [
       true,
       false
      ],
      "type": "string"
     },
     "sentAt": {
      "$ref": "#/components/schemas/sentAt"
     }
    },
    "type": "object"
   }
  },
  "messages": {
   "dimLight": {
    "payload": {
     "$ref": "#/components/schemas/dimLightPayload"
    },
    "summary": "Command a particular streetlight to dim the lights."
   },
   "lightMeasured": {
    "payload": {
     "$ref": "#/components/schemas/lightMeasuredPayload"
    },
    "summary": "Inform about environmental lighting conditions for a particular streetlight."
   },
   "turnOnOff": {
    "payload": {
     "$ref": "#/components/schemas/turnOnOffPayload"
    },
    "summary": "Command a particular streetlight to turn the lights on or off."
   }
  },
  "securitySchemes": {
   "apiKey": {
    "in": "user",
    "description": "Provide your API key as the user and leave the password empty.",
    "type": "apiKey"
   }
  },
  "parameters": {
   "streetlightId": {
    "description": "The ID of the streetlight.",
    "name": "streetlightId",
    "schema": {
     "type": "string"
    }
   }
  }
 },
 "security": [
  {
   "apiKey": []
  }
 ]
}`
	assert.Equal(t, expected, string(data))

	data, err = a.MarshalYAML()
	assert.NoError(t, err)

	expected = `asyncapi: 1.2.0
info:
  description: "The Smartylighting Streetlights API allows you to remotely manage
    the city lights.\n\n### Check out its awesome features:\n\n* Turn a specific streetlight
    on/off \U0001F303\n* Dim a specific streetlight \U0001F60E\n* Receive real-time
    information about environmental lighting conditions \U0001F4C8\n"
  license:
    name: Apache 2.0
    url: https://www.apache.org/licenses/LICENSE-2.0
  title: Streetlights API
  version: 1.0.0
baseTopic: smartylighting.streetlights.1.0
servers:
- description: Test broker
  scheme: mqtt
  url: api.streetlights.smartylighting.com:{port}
  variables:
    port:
      default: "1883"
      description: Secure connection (TLS) is available through port 8883.
      enum:
      - "1883"
      - "8883"
topics:
  action.{streetlightId}.dim:
    parameters:
    - $ref: '#/components/parameters/streetlightId'
    subscribe:
      $ref: '#/components/messages/dimLight'
  action.{streetlightId}.turn.off:
    parameters:
    - $ref: '#/components/parameters/streetlightId'
    subscribe:
      $ref: '#/components/messages/turnOnOff'
  action.{streetlightId}.turn.on:
    parameters:
    - $ref: '#/components/parameters/streetlightId'
    subscribe:
      $ref: '#/components/messages/turnOnOff'
  event.{streetlightId}.lighting.measured:
    parameters:
    - $ref: '#/components/parameters/streetlightId'
    publish:
      $ref: '#/components/messages/lightMeasured'
components:
  messages:
    dimLight:
      payload:
        $ref: '#/components/schemas/dimLightPayload'
      summary: Command a particular streetlight to dim the lights.
    lightMeasured:
      payload:
        $ref: '#/components/schemas/lightMeasuredPayload'
      summary: Inform about environmental lighting conditions for a particular streetlight.
    turnOnOff:
      payload:
        $ref: '#/components/schemas/turnOnOffPayload'
      summary: Command a particular streetlight to turn the lights on or off.
  parameters:
    streetlightId:
      description: The ID of the streetlight.
      name: streetlightId
      schema:
        type: string
  schemas:
    dimLightPayload:
      properties:
        percentage:
          description: Percentage to which the light should be dimmed to.
          maximum: 100
          minimum: 0
          type: integer
        sentAt:
          $ref: '#/components/schemas/sentAt'
      type: object
    lightMeasuredPayload:
      properties:
        lumens:
          description: Light intensity measured in lumens.
          minimum: 0
          type: integer
        sentAt:
          $ref: '#/components/schemas/sentAt'
      type: object
    sentAt:
      description: Date and time when the message was sent.
      format: date-time
      type: string
    turnOnOffPayload:
      properties:
        command:
          description: Whether to turn on or off the light.
          enum:
          - true
          - false
          type: string
        sentAt:
          $ref: '#/components/schemas/sentAt'
      type: object
  securitySchemes:
    apiKey:
      description: Provide your API key as the user and leave the password empty.
      in: user
      type: apiKey
security:
- apiKey: []
`

	assert.Equal(t, expected, string(data))
}
