// Package amqp implements helper to extend AsyncAPI spec
package amqp

import (
	"strings"

	"github.com/swaggest/go-asyncapi/swgen/asyncapi" // nolint:staticcheck
)

const (
	// Exchange defines spec key
	Exchange = "x-amqp-exchange"
	// Exchanges defines spec key
	Exchanges = "x-amqp-exchanges"
	// VHost defines spec key
	VHost = "x-amqp-vhost"
	// Queue defines spec key
	Queue = "x-amqp-queue"
)

// Info keeps amqp topic information
type Info struct {
	VHost     string   `json:"vhost,omitempty"`
	Exchange  string   `json:"exchange,omitempty"`
	Exchanges []string `json:"exchanges,omitempty"`
	Queue     string   `json:"queue,omitempty"`
}

// MessageWithInfo injects amqp info into topic info
func MessageWithInfo(msg *asyncapi.Message, amqpInfo Info) *asyncapi.Message {
	if msg == nil {
		msg = &asyncapi.Message{}
	}

	if msg.MapOfAnything == nil {
		msg.MapOfAnything = map[string]interface{}{}
	}

	desc := msg.Description

	if amqpInfo.VHost != "" {
		msg.MapOfAnything[VHost] = amqpInfo.VHost
		desc += "\n\nAMQP VHost: " + amqpInfo.VHost + "."
	}

	if len(amqpInfo.Exchanges) > 0 {
		if amqpInfo.Exchange != "" {
			amqpInfo.Exchanges = append(amqpInfo.Exchanges, amqpInfo.Exchange)
			amqpInfo.Exchange = ""
		}

		msg.MapOfAnything[Exchanges] = amqpInfo.Exchanges
		desc += "\n\nAMQP Exchanges: " + strings.Join(amqpInfo.Exchanges, ", ") + "."
	}

	if amqpInfo.Exchange != "" {
		msg.MapOfAnything[Exchange] = amqpInfo.Exchange
		desc += "\n\nAMQP Exchange: " + amqpInfo.Exchange + "."
	}

	if amqpInfo.Queue != "" {
		msg.MapOfAnything[Queue] = amqpInfo.Queue
		desc += "\n\nAMQP Queue: " + amqpInfo.Queue + "."
	}

	if desc != "" {
		desc = strings.TrimLeft(desc, "\n")
		msg.Description = desc
	}

	return msg
}
