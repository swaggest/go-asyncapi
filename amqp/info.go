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
)

// Info keeps amqp topic information
type Info struct {
	VHost       string   `json:"vhost,omitempty"`
	Exchange    string   `json:"exchange,omitempty"`
	Exchanges   []string `json:"exchanges,omitempty"`
	RoutingKey  string   `json:"routingKey,omitempty"`
	RoutingKeys []string `json:"routingKeys,omitempty"`
}

// MessageWithInfo injects amqp info into topic info
func MessageWithInfo(msg *asyncapi.Message, amqpInfo Info) *asyncapi.Message {
	if amqpInfo.VHost == "" && amqpInfo.Exchange == "" && amqpInfo.RoutingKey == "" {
		return msg
	}

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

	msg.Description = strings.TrimLeft(msg.Description, ", ")

	return msg
}
