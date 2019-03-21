// Package amqp implements helper to extend AsyncAPI spec
package amqp

import (
	"strings"

	"github.com/swaggest/go-asyncapi/swgen/asyncapi"
)

const (
	// Exchange defines spec key
	Exchange = "x-amqp-exchange"
	// Exchanges defines spec key
	Exchanges = "x-amqp-exchanges"
	// VHost defines spec key
	VHost = "x-amqp-vhost"
	// RoutingKey defines spec key
	RoutingKey = "x-amqp-routing-key"
	// RoutingKeys defines spec key
	RoutingKeys = "x-amqp-routing-keys"
	// Bindings defines spec key
	Bindings = "x-amqp-bindings"
)

// Info keeps amqp topic information
type Info struct {
	VHost       string    `json:"vhost,omitempty"`
	Exchange    string    `json:"exchange,omitempty"`
	Exchanges   []string  `json:"exchanges,omitempty"`
	RoutingKey  string    `json:"routingKey,omitempty"`
	RoutingKeys []string  `json:"routingKeys,omitempty"`
	Bindings    []Binding `json:"bindings,omitempty"`
}

// Binding is amqp binding
type Binding struct {
	Exchange   string `json:"exchange,omitempty" yaml:"exchange"`
	RoutingKey string `json:"routingKey,omitempty" yaml:"routing-key"`
}

// MessageWithInfo injects amqp info into topic info
func MessageWithInfo(msg *asyncapi.Message, amqpInfo Info) *asyncapi.Message {
	if msg == nil {
		msg = &asyncapi.Message{}
	}

	if msg.MapOfAnythingValues == nil {
		msg.MapOfAnythingValues = map[string]interface{}{}
	}

	if amqpInfo.VHost != "" {
		msg.MapOfAnythingValues[VHost] = amqpInfo.VHost
		msg.Description += ", AMQP VHost: " + amqpInfo.VHost
	}
	if len(amqpInfo.Exchanges) > 0 {
		if amqpInfo.Exchange != "" {
			amqpInfo.Exchanges = append(amqpInfo.Exchanges, amqpInfo.Exchange)
			amqpInfo.Exchange = ""
		}

		msg.MapOfAnythingValues[Exchanges] = amqpInfo.Exchanges
		msg.Description += ", AMQP Exchanges: " + strings.Join(amqpInfo.Exchanges, ", ")
	}
	if amqpInfo.Exchange != "" {
		msg.MapOfAnythingValues[Exchange] = amqpInfo.Exchange
		msg.Description += ", AMQP Exchange: " + amqpInfo.Exchange
	}

	if len(amqpInfo.RoutingKeys) > 0 {
		if amqpInfo.RoutingKey != "" {
			amqpInfo.RoutingKeys = append(amqpInfo.RoutingKeys, amqpInfo.RoutingKey)
			amqpInfo.RoutingKey = ""
		}

		msg.MapOfAnythingValues[RoutingKeys] = amqpInfo.RoutingKeys
		msg.Description += ", AMQP RoutingKeys: " + strings.Join(amqpInfo.RoutingKeys, ", ")
	}
	if amqpInfo.RoutingKey != "" {
		msg.MapOfAnythingValues[RoutingKey] = amqpInfo.RoutingKey
		msg.Description += ", AMQP RoutingKey: " + amqpInfo.RoutingKey
	}

	if len(amqpInfo.Bindings) > 0 {
		msg.MapOfAnythingValues[Bindings] = amqpInfo.Bindings
		msg.Description += ", AMQP bindings: \n"
		for _, b := range amqpInfo.Bindings {
			msg.Description += " * Exchange: " + b.Exchange + ", Routing Key: " + b.RoutingKey + "\n"
		}
	}

	msg.Description = strings.TrimLeft(msg.Description, ", ")

	return msg
}
