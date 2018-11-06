package spec

import (
	"encoding/json"
	"github.com/pkg/errors"
	"regexp"
	"strings"
)

// AsyncAPI structure is generated from #
// AsyncAPI 1.2.0 schema.
type AsyncAPI struct {
	Asyncapi            Asyncapi               `json:"asyncapi,omitempty"`  // The AsyncAPI specification version of this document.
	Info                *Info                  `json:"info,omitempty"`      // General information about the API.
	BaseTopic           string                 `json:"baseTopic,omitempty"` // The base topic to the API. Example: 'hitch'.
	Servers             []Server               `json:"servers,omitempty"`
	Topics              *Topics                `json:"topics,omitempty"`     // Relative paths to the individual topics. They must be relative to the 'baseTopic'.
	Stream              *Stream                `json:"stream,omitempty"`     // Stream Object
	Events              *Events                `json:"events,omitempty"`     // Events Object
	Components          *Components            `json:"components,omitempty"` // An object to hold a set of reusable objects for different aspects of the AsyncAPI Specification.
	Tags                []Tag                  `json:"tags,omitempty"`
	Security            []map[string][]string  `json:"security,omitempty"`
	ExternalDocs        *ExternalDocs          `json:"externalDocs,omitempty"` // information about external documentation
	MapOfAnythingValues map[string]interface{} `json:"-"`                      // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *AsyncAPI) UnmarshalJSON(data []byte) error {
	type p AsyncAPI

	ii := p(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"asyncapi", "info", "baseTopic", "servers", "topics", "stream", "events", "components", "tags", "security", "externalDocs"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = AsyncAPI(ii)

	return err
}

// MarshalJSON encodes JSON
func (i AsyncAPI) MarshalJSON() ([]byte, error) {
	type p AsyncAPI

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// Info structure is generated from #/definitions/info
// General information about the API.
type Info struct {
	Title               string                 `json:"title,omitempty"`          // A unique and precise title of the API.
	Version             string                 `json:"version,omitempty"`        // A semantic version number of the API.
	Description         string                 `json:"description,omitempty"`    // A longer description of the API. Should be different from the title. CommonMark is allowed.
	TermsOfService      string                 `json:"termsOfService,omitempty"` // A URL to the Terms of Service for the API. MUST be in the format of a URL.
	Contact             *Contact               `json:"contact,omitempty"`        // Contact information for the owners of the API.
	License             *License               `json:"license,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"` // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *Info) UnmarshalJSON(data []byte) error {
	type p Info

	ii := p(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"title", "version", "description", "termsOfService", "contact", "license"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = Info(ii)

	return err
}

// MarshalJSON encodes JSON
func (i Info) MarshalJSON() ([]byte, error) {
	type p Info

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// Contact structure is generated from #/definitions/contact
// Contact information for the owners of the API.
type Contact struct {
	Name                string                 `json:"name,omitempty"`  // The identifying name of the contact person/organization.
	URL                 string                 `json:"url,omitempty"`   // The URL pointing to the contact information.
	Email               string                 `json:"email,omitempty"` // The email address of the contact person/organization.
	MapOfAnythingValues map[string]interface{} `json:"-"`               // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *Contact) UnmarshalJSON(data []byte) error {
	type p Contact

	ii := p(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"name", "url", "email"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = Contact(ii)

	return err
}

// MarshalJSON encodes JSON
func (i Contact) MarshalJSON() ([]byte, error) {
	type p Contact

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// License structure is generated from #/definitions/license
type License struct {
	Name                string                 `json:"name,omitempty"` // The name of the license type. It's encouraged to use an OSI compatible license.
	URL                 string                 `json:"url,omitempty"`  // The URL pointing to the license.
	MapOfAnythingValues map[string]interface{} `json:"-"`              // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *License) UnmarshalJSON(data []byte) error {
	type p License

	ii := p(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"name", "url"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = License(ii)

	return err
}

// MarshalJSON encodes JSON
func (i License) MarshalJSON() ([]byte, error) {
	type p License

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// Server structure is generated from #/definitions/server
// An object representing a Server.
type Server struct {
	URL                 string                    `json:"url,omitempty"`
	Description         string                    `json:"description,omitempty"`
	Scheme              ServerScheme              `json:"scheme,omitempty"` // The transfer protocol.
	SchemeVersion       string                    `json:"schemeVersion,omitempty"`
	Variables           map[string]ServerVariable `json:"variables,omitempty"`
	MapOfAnythingValues map[string]interface{}    `json:"-"` // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *Server) UnmarshalJSON(data []byte) error {
	type p Server

	ii := p(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"url", "description", "scheme", "schemeVersion", "variables"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = Server(ii)

	return err
}

// MarshalJSON encodes JSON
func (i Server) MarshalJSON() ([]byte, error) {
	type p Server

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// ServerVariable structure is generated from #/definitions/serverVariable
// An object representing a Server Variable for server URL template substitution.
type ServerVariable struct {
	Enum                []string               `json:"enum,omitempty"`
	Default             string                 `json:"default,omitempty"`
	Description         string                 `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"` // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *ServerVariable) UnmarshalJSON(data []byte) error {
	type p ServerVariable

	ii := p(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"enum", "default", "description"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = ServerVariable(ii)

	return err
}

// MarshalJSON encodes JSON
func (i ServerVariable) MarshalJSON() ([]byte, error) {
	type p ServerVariable

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// Topics structure is generated from #/definitions/topics
// Relative paths to the individual topics. They must be relative to the 'baseTopic'.
type Topics struct {
	MapOfAnythingValues  map[string]interface{} `json:"-"` // Key must match pattern: ^x-
	MapOfTopicItemValues map[string]TopicItem   `json:"-"` // Key must match pattern: ^[^.]
}

// UnmarshalJSON decodes JSON
func (i *Topics) UnmarshalJSON(data []byte) error {

	err := unmarshalUnion(
		nil,
		nil,
		nil,
		map[string]interface{}{
			`^x-`:   &i.MapOfAnythingValues,
			`^[^.]`: &i.MapOfTopicItemValues,
		},
		data,
	)

	return err
}

// MarshalJSON encodes JSON
func (i Topics) MarshalJSON() ([]byte, error) {
	type p Topics

	return marshalUnion(p(i), i.MapOfAnythingValues, i.MapOfTopicItemValues)
}

// TopicItem structure is generated from #/definitions/topicItem
type TopicItem struct {
	Ref                 string                 `json:"$ref,omitempty"`
	Parameters          []Parameter            `json:"parameters,omitempty"`
	Publish             *Operation             `json:"publish,omitempty"`
	Subscribe           *Operation             `json:"subscribe,omitempty"`
	Deprecated          bool                   `json:"deprecated,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"` // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *TopicItem) UnmarshalJSON(data []byte) error {
	type p TopicItem

	ii := p(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"$ref", "parameters", "publish", "subscribe", "deprecated"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = TopicItem(ii)

	return err
}

// MarshalJSON encodes JSON
func (i TopicItem) MarshalJSON() ([]byte, error) {
	type p TopicItem

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// Parameter structure is generated from #/definitions/parameter
type Parameter struct {
	Description         string                 `json:"description,omitempty"` // A brief description of the parameter. This could contain examples of use.  GitHub Flavored Markdown is allowed.
	Name                string                 `json:"name,omitempty"`        // The name of the parameter.
	Schema              map[string]interface{} `json:"schema,omitempty"`      // A deterministic version of a JSON Schema object.
	Ref                 string                 `json:"$ref,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"` // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *Parameter) UnmarshalJSON(data []byte) error {
	type p Parameter

	ii := p(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"description", "name", "schema", "$ref"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = Parameter(ii)

	return err
}

// MarshalJSON encodes JSON
func (i Parameter) MarshalJSON() ([]byte, error) {
	type p Parameter

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// Message structure is generated from #/definitions/message
type Message struct {
	Ref                 string                 `json:"$ref,omitempty"`
	Headers             map[string]interface{} `json:"headers,omitempty"` // A deterministic version of a JSON Schema object.
	Payload             map[string]interface{} `json:"payload,omitempty"` // A deterministic version of a JSON Schema object.
	Tags                []Tag                  `json:"tags,omitempty"`
	Summary             string                 `json:"summary,omitempty"`      // A brief summary of the message.
	Description         string                 `json:"description,omitempty"`  // A longer description of the message. CommonMark is allowed.
	ExternalDocs        *ExternalDocs          `json:"externalDocs,omitempty"` // information about external documentation
	Deprecated          bool                   `json:"deprecated,omitempty"`
	Example             interface{}            `json:"example,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"` // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *Message) UnmarshalJSON(data []byte) error {
	type p Message

	ii := p(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"$ref", "headers", "payload", "tags", "summary", "description", "externalDocs", "deprecated", "example"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = Message(ii)

	return err
}

// MarshalJSON encodes JSON
func (i Message) MarshalJSON() ([]byte, error) {
	type p Message

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// Tag structure is generated from #/definitions/tag
type Tag struct {
	Name                string                 `json:"name,omitempty"`
	Description         string                 `json:"description,omitempty"`
	ExternalDocs        *ExternalDocs          `json:"externalDocs,omitempty"` // information about external documentation
	MapOfAnythingValues map[string]interface{} `json:"-"`                      // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *Tag) UnmarshalJSON(data []byte) error {
	type p Tag

	ii := p(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"name", "description", "externalDocs"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = Tag(ii)

	return err
}

// MarshalJSON encodes JSON
func (i Tag) MarshalJSON() ([]byte, error) {
	type p Tag

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// ExternalDocs structure is generated from #/definitions/externalDocs
// information about external documentation
type ExternalDocs struct {
	Description         string                 `json:"description,omitempty"`
	URL                 string                 `json:"url,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"` // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *ExternalDocs) UnmarshalJSON(data []byte) error {
	type p ExternalDocs

	ii := p(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"description", "url"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = ExternalDocs(ii)

	return err
}

// MarshalJSON encodes JSON
func (i ExternalDocs) MarshalJSON() ([]byte, error) {
	type p ExternalDocs

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// Operation structure is generated from #/definitions/operation
type Operation struct {
	Message         *Message         `json:"-"`
	OperationOneOf1 *OperationOneOf1 `json:"-"`
}

// UnmarshalJSON decodes JSON
func (i *Operation) UnmarshalJSON(data []byte) error {

	err := unmarshalUnion(
		nil,
		[]interface{}{&i.Message, &i.OperationOneOf1},
		nil,
		nil,
		data,
	)

	return err
}

// MarshalJSON encodes JSON
func (i Operation) MarshalJSON() ([]byte, error) {
	type p Operation

	return marshalUnion(p(i), i.Message, i.OperationOneOf1)
}

// OperationOneOf1 structure is generated from #/definitions/operation/oneOf/1
type OperationOneOf1 struct {
	OneOf               []Message              `json:"oneOf,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"` // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *OperationOneOf1) UnmarshalJSON(data []byte) error {
	type p OperationOneOf1

	ii := p(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"oneOf"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = OperationOneOf1(ii)

	return err
}

// MarshalJSON encodes JSON
func (i OperationOneOf1) MarshalJSON() ([]byte, error) {
	type p OperationOneOf1

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// Stream structure is generated from #/definitions/stream
// Stream Object
type Stream struct {
	Framing             *StreamFraming         `json:"framing,omitempty"` // Stream Framing Object
	Read                []Message              `json:"read,omitempty"`    // Stream Read Object
	Write               []Message              `json:"write,omitempty"`   // Stream Write Object
	MapOfAnythingValues map[string]interface{} `json:"-"`                 // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *Stream) UnmarshalJSON(data []byte) error {
	type p Stream

	ii := p(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"framing", "read", "write"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = Stream(ii)

	return err
}

// MarshalJSON encodes JSON
func (i Stream) MarshalJSON() ([]byte, error) {
	type p Stream

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// StreamFramingOneOf0 structure is generated from #/definitions/stream->framing/oneOf/0
type StreamFramingOneOf0 struct {
	Type      StreamFramingOneOf0Type      `json:"type,omitempty"`
	Delimiter StreamFramingOneOf0Delimiter `json:"delimiter,omitempty"`
}

// StreamFraming structure is generated from #/definitions/stream->framing
// Stream Framing Object
type StreamFraming struct {
	StreamFramingOneOf0 *StreamFramingOneOf0   `json:"-"`
	StreamFramingOneOf1 *StreamFramingOneOf1   `json:"-"`
	MapOfAnythingValues map[string]interface{} `json:"-"` // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *StreamFraming) UnmarshalJSON(data []byte) error {

	err := unmarshalUnion(
		nil,
		[]interface{}{&i.StreamFramingOneOf0, &i.StreamFramingOneOf1},
		nil,
		map[string]interface{}{
			`^x-`: &i.MapOfAnythingValues,
		},
		data,
	)

	return err
}

// MarshalJSON encodes JSON
func (i StreamFraming) MarshalJSON() ([]byte, error) {
	type p StreamFraming

	return marshalUnion(p(i), i.MapOfAnythingValues, i.StreamFramingOneOf0, i.StreamFramingOneOf1)
}

// StreamFramingOneOf1 structure is generated from #/definitions/stream->framing/oneOf/1
type StreamFramingOneOf1 struct {
	Type      StreamFramingOneOf1Type      `json:"type,omitempty"`
	Delimiter StreamFramingOneOf1Delimiter `json:"delimiter,omitempty"`
}

// Events structure is generated from #/definitions/events
// Events Object
type Events struct {
	Receive             []Message              `json:"receive,omitempty"` // Events Receive Object
	Send                []Message              `json:"send,omitempty"`    // Events Send Object
	MapOfAnythingValues map[string]interface{} `json:"-"`                 // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *Events) UnmarshalJSON(data []byte) error {
	type p Events

	ii := p(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"receive", "send"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = Events(ii)

	return err
}

// MarshalJSON encodes JSON
func (i Events) MarshalJSON() ([]byte, error) {
	type p Events

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// Components structure is generated from #/definitions/components
// An object to hold a set of reusable objects for different aspects of the AsyncAPI Specification.
type Components struct {
	Schemas         map[string]map[string]interface{} `json:"schemas,omitempty"`  // JSON objects describing schemas the API uses.
	Messages        map[string]Message                `json:"messages,omitempty"` // JSON objects describing the messages being consumed and produced by the API.
	SecuritySchemes *ComponentsSecuritySchemes        `json:"securitySchemes,omitempty"`
	Parameters      map[string]Parameter              `json:"parameters,omitempty"` // JSON objects describing re-usable topic parameters.
}

// Reference structure is generated from #/definitions/Reference
type Reference struct {
	Ref string `json:"$ref,omitempty"`
}

// ComponentsSecuritySchemesAZAZ09 structure is generated from #/definitions/components->securitySchemes->^[a-zA-Z0-9\.\-_]+$
type ComponentsSecuritySchemesAZAZ09 struct {
	Reference      *Reference      `json:"-"`
	SecurityScheme *SecurityScheme `json:"-"`
}

// UnmarshalJSON decodes JSON
func (i *ComponentsSecuritySchemesAZAZ09) UnmarshalJSON(data []byte) error {

	err := unmarshalUnion(
		nil,
		[]interface{}{&i.Reference, &i.SecurityScheme},
		nil,
		nil,
		data,
	)

	return err
}

// MarshalJSON encodes JSON
func (i ComponentsSecuritySchemesAZAZ09) MarshalJSON() ([]byte, error) {
	type p ComponentsSecuritySchemesAZAZ09

	return marshalUnion(p(i), i.Reference, i.SecurityScheme)
}

// UserPassword structure is generated from #/definitions/userPassword
type UserPassword struct {
	Type                UserPasswordType       `json:"type,omitempty"`
	Description         string                 `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"` // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *UserPassword) UnmarshalJSON(data []byte) error {
	type p UserPassword

	ii := p(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"type", "description"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = UserPassword(ii)

	return err
}

// MarshalJSON encodes JSON
func (i UserPassword) MarshalJSON() ([]byte, error) {
	type p UserPassword

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// SecurityScheme structure is generated from #/definitions/SecurityScheme
type SecurityScheme struct {
	UserPassword         *UserPassword         `json:"-"`
	APIKey               *APIKey               `json:"-"`
	X509                 *X509                 `json:"-"`
	SymmetricEncryption  *SymmetricEncryption  `json:"-"`
	AsymmetricEncryption *AsymmetricEncryption `json:"-"`
	HTTPSecurityScheme   *HTTPSecurityScheme   `json:"-"`
}

// UnmarshalJSON decodes JSON
func (i *SecurityScheme) UnmarshalJSON(data []byte) error {

	err := unmarshalUnion(
		nil,
		[]interface{}{&i.UserPassword, &i.APIKey, &i.X509, &i.SymmetricEncryption, &i.AsymmetricEncryption, &i.HTTPSecurityScheme},
		nil,
		nil,
		data,
	)

	return err
}

// MarshalJSON encodes JSON
func (i SecurityScheme) MarshalJSON() ([]byte, error) {
	type p SecurityScheme

	return marshalUnion(p(i), i.UserPassword, i.APIKey, i.X509, i.SymmetricEncryption, i.AsymmetricEncryption, i.HTTPSecurityScheme)
}

// APIKey structure is generated from #/definitions/apiKey
type APIKey struct {
	Type                APIKeyType             `json:"type,omitempty"`
	In                  APIKeyIn               `json:"in,omitempty"`
	Description         string                 `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"` // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *APIKey) UnmarshalJSON(data []byte) error {
	type p APIKey

	ii := p(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"type", "in", "description"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = APIKey(ii)

	return err
}

// MarshalJSON encodes JSON
func (i APIKey) MarshalJSON() ([]byte, error) {
	type p APIKey

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// X509 structure is generated from #/definitions/X509
type X509 struct {
	Type                X509Type               `json:"type,omitempty"`
	Description         string                 `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"` // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *X509) UnmarshalJSON(data []byte) error {
	type p X509

	ii := p(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"type", "description"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = X509(ii)

	return err
}

// MarshalJSON encodes JSON
func (i X509) MarshalJSON() ([]byte, error) {
	type p X509

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// SymmetricEncryption structure is generated from #/definitions/symmetricEncryption
type SymmetricEncryption struct {
	Type                SymmetricEncryptionType `json:"type,omitempty"`
	Description         string                  `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{}  `json:"-"` // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *SymmetricEncryption) UnmarshalJSON(data []byte) error {
	type p SymmetricEncryption

	ii := p(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"type", "description"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = SymmetricEncryption(ii)

	return err
}

// MarshalJSON encodes JSON
func (i SymmetricEncryption) MarshalJSON() ([]byte, error) {
	type p SymmetricEncryption

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// AsymmetricEncryption structure is generated from #/definitions/asymmetricEncryption
type AsymmetricEncryption struct {
	Type                AsymmetricEncryptionType `json:"type,omitempty"`
	Description         string                   `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{}   `json:"-"` // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *AsymmetricEncryption) UnmarshalJSON(data []byte) error {
	type p AsymmetricEncryption

	ii := p(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"type", "description"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = AsymmetricEncryption(ii)

	return err
}

// MarshalJSON encodes JSON
func (i AsymmetricEncryption) MarshalJSON() ([]byte, error) {
	type p AsymmetricEncryption

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// NonBearerHTTPSecurityScheme structure is generated from #/definitions/NonBearerHTTPSecurityScheme
type NonBearerHTTPSecurityScheme struct {
	Scheme              string                          `json:"scheme,omitempty"`
	Description         string                          `json:"description,omitempty"`
	Type                NonBearerHTTPSecuritySchemeType `json:"type,omitempty"`
	MapOfAnythingValues map[string]interface{}          `json:"-"` // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *NonBearerHTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	type p NonBearerHTTPSecurityScheme

	ii := p(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"scheme", "description", "type"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = NonBearerHTTPSecurityScheme(ii)

	return err
}

// MarshalJSON encodes JSON
func (i NonBearerHTTPSecurityScheme) MarshalJSON() ([]byte, error) {
	type p NonBearerHTTPSecurityScheme

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// HTTPSecurityScheme structure is generated from #/definitions/HTTPSecurityScheme
type HTTPSecurityScheme struct {
	NonBearerHTTPSecurityScheme *NonBearerHTTPSecurityScheme `json:"-"`
	BearerHTTPSecurityScheme    *BearerHTTPSecurityScheme    `json:"-"`
	APIKeyHTTPSecurityScheme    *APIKeyHTTPSecurityScheme    `json:"-"`
}

// UnmarshalJSON decodes JSON
func (i *HTTPSecurityScheme) UnmarshalJSON(data []byte) error {

	err := unmarshalUnion(
		nil,
		[]interface{}{&i.NonBearerHTTPSecurityScheme, &i.BearerHTTPSecurityScheme, &i.APIKeyHTTPSecurityScheme},
		nil,
		nil,
		data,
	)

	return err
}

// MarshalJSON encodes JSON
func (i HTTPSecurityScheme) MarshalJSON() ([]byte, error) {
	type p HTTPSecurityScheme

	return marshalUnion(p(i), i.NonBearerHTTPSecurityScheme, i.BearerHTTPSecurityScheme, i.APIKeyHTTPSecurityScheme)
}

// BearerHTTPSecurityScheme structure is generated from #/definitions/BearerHTTPSecurityScheme
type BearerHTTPSecurityScheme struct {
	Scheme              BearerHTTPSecuritySchemeScheme `json:"scheme,omitempty"`
	BearerFormat        string                         `json:"bearerFormat,omitempty"`
	Type                BearerHTTPSecuritySchemeType   `json:"type,omitempty"`
	Description         string                         `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{}         `json:"-"` // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *BearerHTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	type p BearerHTTPSecurityScheme

	ii := p(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"scheme", "bearerFormat", "type", "description"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = BearerHTTPSecurityScheme(ii)

	return err
}

// MarshalJSON encodes JSON
func (i BearerHTTPSecurityScheme) MarshalJSON() ([]byte, error) {
	type p BearerHTTPSecurityScheme

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// APIKeyHTTPSecurityScheme structure is generated from #/definitions/APIKeyHTTPSecurityScheme
type APIKeyHTTPSecurityScheme struct {
	Type                APIKeyHTTPSecuritySchemeType `json:"type,omitempty"`
	Name                string                       `json:"name,omitempty"`
	In                  APIKeyHTTPSecuritySchemeIn   `json:"in,omitempty"`
	Description         string                       `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{}       `json:"-"` // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *APIKeyHTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	type p APIKeyHTTPSecurityScheme

	ii := p(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"type", "name", "in", "description"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = APIKeyHTTPSecurityScheme(ii)

	return err
}

// MarshalJSON encodes JSON
func (i APIKeyHTTPSecurityScheme) MarshalJSON() ([]byte, error) {
	type p APIKeyHTTPSecurityScheme

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// ComponentsSecuritySchemes structure is generated from #/definitions/components->securitySchemes
type ComponentsSecuritySchemes struct {
	MapOfComponentsSecuritySchemesAZAZ09Values map[string]ComponentsSecuritySchemesAZAZ09 `json:"-"` // Key must match pattern: ^[a-zA-Z0-9\.\-_]+$
}

// UnmarshalJSON decodes JSON
func (i *ComponentsSecuritySchemes) UnmarshalJSON(data []byte) error {

	err := unmarshalUnion(
		nil,
		nil,
		nil,
		map[string]interface{}{
			`^[a-zA-Z0-9\.\-_]+$`: &i.MapOfComponentsSecuritySchemesAZAZ09Values,
		},
		data,
	)

	return err
}

// MarshalJSON encodes JSON
func (i ComponentsSecuritySchemes) MarshalJSON() ([]byte, error) {
	type p ComponentsSecuritySchemes

	return marshalUnion(p(i), i.MapOfComponentsSecuritySchemesAZAZ09Values)
}

type Asyncapi string

// Asyncapi values enumeration
const (
	Asyncapi100 = Asyncapi(`1.0.0`)
	Asyncapi110 = Asyncapi(`1.1.0`)
	Asyncapi120 = Asyncapi(`1.2.0`)
)

type ServerScheme string

// ServerScheme values enumeration
const (
	ServerSchemeKafka       = ServerScheme(`kafka`)
	ServerSchemeKafkaSecure = ServerScheme(`kafka-secure`)
	ServerSchemeAmqp        = ServerScheme(`amqp`)
	ServerSchemeAmqps       = ServerScheme(`amqps`)
	ServerSchemeMqtt        = ServerScheme(`mqtt`)
	ServerSchemeMqtts       = ServerScheme(`mqtts`)
	ServerSchemeSecureMqtt  = ServerScheme(`secure-mqtt`)
	ServerSchemeWs          = ServerScheme(`ws`)
	ServerSchemeWss         = ServerScheme(`wss`)
	ServerSchemeStomp       = ServerScheme(`stomp`)
	ServerSchemeStomps      = ServerScheme(`stomps`)
	ServerSchemeJms         = ServerScheme(`jms`)
	ServerSchemeHTTP        = ServerScheme(`http`)
	ServerSchemeHTTPS       = ServerScheme(`https`)
)

type StreamFramingOneOf0Type string

// StreamFramingOneOf0Type values enumeration
const (
	StreamFramingOneOf0TypeChunked = StreamFramingOneOf0Type(`chunked`)
)

type StreamFramingOneOf0Delimiter string

// StreamFramingOneOf0Delimiter values enumeration
const (
	StreamFramingOneOf0DelimiterRN = StreamFramingOneOf0Delimiter(`\r\n`)
	StreamFramingOneOf0DelimiterN  = StreamFramingOneOf0Delimiter(`\n`)
)

type StreamFramingOneOf1Type string

// StreamFramingOneOf1Type values enumeration
const (
	StreamFramingOneOf1TypeSse = StreamFramingOneOf1Type(`sse`)
)

type StreamFramingOneOf1Delimiter string

// StreamFramingOneOf1Delimiter values enumeration
const (
	StreamFramingOneOf1DelimiterNN = StreamFramingOneOf1Delimiter(`\n\n`)
)

type UserPasswordType string

// UserPasswordType values enumeration
const (
	UserPasswordTypeUserPassword = UserPasswordType(`userPassword`)
)

type APIKeyType string

// APIKeyType values enumeration
const (
	APIKeyTypeAPIKey = APIKeyType(`apiKey`)
)

type APIKeyIn string

// APIKeyIn values enumeration
const (
	APIKeyInUser     = APIKeyIn(`user`)
	APIKeyInPassword = APIKeyIn(`password`)
)

type X509Type string

// X509Type values enumeration
const (
	X509TypeX509 = X509Type(`X509`)
)

type SymmetricEncryptionType string

// SymmetricEncryptionType values enumeration
const (
	SymmetricEncryptionTypeSymmetricEncryption = SymmetricEncryptionType(`symmetricEncryption`)
)

type AsymmetricEncryptionType string

// AsymmetricEncryptionType values enumeration
const (
	AsymmetricEncryptionTypeAsymmetricEncryption = AsymmetricEncryptionType(`asymmetricEncryption`)
)

type NonBearerHTTPSecuritySchemeType string

// NonBearerHTTPSecuritySchemeType values enumeration
const (
	NonBearerHTTPSecuritySchemeTypeHTTP = NonBearerHTTPSecuritySchemeType(`http`)
)

type BearerHTTPSecuritySchemeScheme string

// BearerHTTPSecuritySchemeScheme values enumeration
const (
	BearerHTTPSecuritySchemeSchemeBearer = BearerHTTPSecuritySchemeScheme(`bearer`)
)

type BearerHTTPSecuritySchemeType string

// BearerHTTPSecuritySchemeType values enumeration
const (
	BearerHTTPSecuritySchemeTypeHTTP = BearerHTTPSecuritySchemeType(`http`)
)

type APIKeyHTTPSecuritySchemeType string

// APIKeyHTTPSecuritySchemeType values enumeration
const (
	APIKeyHTTPSecuritySchemeTypeHTTPAPIKey = APIKeyHTTPSecuritySchemeType(`httpApiKey`)
)

type APIKeyHTTPSecuritySchemeIn string

// APIKeyHTTPSecuritySchemeIn values enumeration
const (
	APIKeyHTTPSecuritySchemeInHeader = APIKeyHTTPSecuritySchemeIn(`header`)
	APIKeyHTTPSecuritySchemeInQuery  = APIKeyHTTPSecuritySchemeIn(`query`)
	APIKeyHTTPSecuritySchemeInCookie = APIKeyHTTPSecuritySchemeIn(`cookie`)
)

func unmarshalUnion(mustUnmarshal []interface{}, mayUnmarshal []interface{}, ignoreKeys []string, regexMaps map[string]interface{}, j []byte) error {
	for _, item := range mustUnmarshal {
		// unmarshal to struct
		err := json.Unmarshal(j, item)
		if err != nil {
			return err
		}
	}

	for _, item := range mayUnmarshal {
		// unmarshal to struct
		_ = json.Unmarshal(j, item)
	}

	// unmarshal to a generic map
	var m map[string]*json.RawMessage
	err := json.Unmarshal(j, &m)
	if err != nil {
		return err
	}

	// removing ignored keys (defined in struct)
	for _, i := range ignoreKeys {
		delete(m, i)
	}

	// returning early on empty map
	if len(m) == 0 {
		return nil
	}

	// preparing regexp matchers
	var reg = make(map[string]*regexp.Regexp, len(regexMaps))
	for regex, _ := range regexMaps {
		if regex != "" {
			reg[regex], err = regexp.Compile(regex)
			if err != nil {
				return err //todo use errors.Wrap
			}
		}
	}

	subMapsRaw := make(map[string][]byte, len(regexMaps))
	// iterating map and feeding subMaps
	for key, val := range m {
		matched := false
		var ok bool
		keyEscaped := `"` + strings.Replace(key, `"`, `\"`, -1) + `":`

		for regex, exp := range reg {
			if exp.Match([]byte(key)) {
				matched = true
				var subMap []byte
				if subMap, ok = subMapsRaw[regex]; !ok {
					subMap = make([]byte, 1, 100)
					subMap[0] = '{'
				} else {
					subMap = append(subMap[:len(subMap)-1], ',')
				}

				subMap = append(subMap, []byte(keyEscaped)...)
				subMap = append(subMap, []byte(*val)...)
				subMap = append(subMap, '}')

				subMapsRaw[regex] = subMap
			}
		}

		// empty regex for additionalProperties
		if !matched {
			var subMap []byte
			if subMap, ok = subMapsRaw[""]; !ok {
				subMap = make([]byte, 1, 100)
				subMap[0] = '{'
			} else {
				subMap = append(subMap[:len(subMap)-1], ',')
			}
			subMap = append(subMap, []byte(keyEscaped)...)
			subMap = append(subMap, []byte(*val)...)
			subMap = append(subMap, '}')

			subMapsRaw[""] = subMap
		}
	}

	for regex, _ := range regexMaps {
		if subMap, ok := subMapsRaw[regex]; !ok {
			continue
		} else {
			err = json.Unmarshal(subMap, regexMaps[regex])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func marshalUnion(maps ...interface{}) ([]byte, error) {
	result := make([]byte, 1, 100)
	result[0] = '{'
	for _, m := range maps {
		if m == nil {
			continue
		}
		j, err := json.Marshal(m)
		if err != nil {
			return nil, err
		}
		if string(j) == "{}" {
			continue
		}
		if string(j) == "null" {
			continue
		}
		if j[0] != '{' {
			return nil, errors.New("failed to union map: object expected, " + string(j) + " received")
		}

		if len(result) > 1 {
			result = append(result[:len(result)-1], ',')
		}
		result = append(result, j[1:]...)
	}
	// closing empty result
	if len(result) == 1 {
		result = append(result, '}')
	}

	return result, nil
}
