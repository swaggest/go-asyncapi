// Package amqp implements helper to extend AsyncAPI spec
package amqp

import (
	"strings"

	"github.com/swaggest/go-asyncapi/swgen/asyncapi"
)

const (
	// Exchange defines spec key
	Exchange = "x-amqp-exchange"
	// VHost defines spec key
	VHost = "x-amqp-vhost"
	// RoutingKey defines spec key
	RoutingKey = "x-amqp-routing-key"
)

// Info keeps amqp topic information
type Info struct {
	VHost      string `json:"vhost,omitempty"`
	Exchange   string `json:"exchange,omitempty"`
	RoutingKey string `json:"routingKey"`
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
	if amqpInfo.Exchange != "" {
		msg.MapOfAnythingValues[Exchange] = amqpInfo.Exchange
		msg.Description += ", AMQP Exchange: " + amqpInfo.Exchange
	}
	if amqpInfo.RoutingKey != "" {
		msg.MapOfAnythingValues[RoutingKey] = amqpInfo.RoutingKey
		msg.Description += ", AMQP RoutingKey: " + amqpInfo.RoutingKey
	}

	msg.Description = strings.TrimLeft(msg.Description, ", ")

	return msg
}
