// Package amqp implements helper to extend AsyncAPI spec
package amqp

import (
	"strings"

	"github.com/swaggest/go-asyncapi/swgen/asyncapi" // nolint:staticcheck // Deprecated.
)

const (
	// Exchange defines spec key.
	Exchange = "x-amqp-exchange"
	// Exchanges defines spec key.
	Exchanges = "x-amqp-exchanges"
	// VHost defines spec key.
	VHost = "x-amqp-vhost"
	// Queue defines spec key.
	Queue = "x-amqp-queue"
)

// Info keeps amqp topic information.
type Info struct {
	VHost     string   `json:"vhost,omitempty"`
	Exchange  string   `json:"exchange,omitempty"`
	Exchanges []string `json:"exchanges,omitempty"`
	Queue     string   `json:"queue,omitempty"`
}

// MessageWithInfo injects amqp info into topic info.
func MessageWithInfo(msg *asyncapi.Message, amqpInfo Info) *asyncapi.Message {
	if msg == nil {
		msg = &asyncapi.Message{}
	}

	if msg.MapOfAnythingValues == nil {
		msg.MapOfAnythingValues = map[string]interface{}{}
	}

	if amqpInfo.VHost != "" {
		msg.MapOfAnythingValues[VHost] = amqpInfo.VHost
		msg.Description += "\n\nAMQP VHost: " + amqpInfo.VHost + "."
	}

	if len(amqpInfo.Exchanges) > 0 {
		if amqpInfo.Exchange != "" {
			amqpInfo.Exchanges = append(amqpInfo.Exchanges, amqpInfo.Exchange)
			amqpInfo.Exchange = ""
		}

		msg.MapOfAnythingValues[Exchanges] = amqpInfo.Exchanges
		msg.Description += "\n\nAMQP Exchanges: " + strings.Join(amqpInfo.Exchanges, ", ") + "."
	}

	if amqpInfo.Exchange != "" {
		msg.MapOfAnythingValues[Exchange] = amqpInfo.Exchange
		msg.Description += "\n\nAMQP Exchange: " + amqpInfo.Exchange + "."
	}

	if amqpInfo.Queue != "" {
		msg.MapOfAnythingValues[Queue] = amqpInfo.Queue
		msg.Description += "\n\nAMQP Queue: " + amqpInfo.Queue + "."
	}

	msg.Description = strings.TrimLeft(msg.Description, "\n")

	return msg
}
