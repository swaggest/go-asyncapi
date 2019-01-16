// Package amqp implements helper to extend AsyncAPI spec
package amqp

import (
	"github.com/swaggest/go-asyncapi/spec"
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

// TopicWithInfo injects amqp info into topic info
func TopicWithInfo(topicInfo asyncapi.TopicInfo, amqpInfo Info) asyncapi.TopicInfo {
	if amqpInfo.VHost == "" && amqpInfo.Exchange == "" && amqpInfo.RoutingKey == "" {
		return topicInfo
	}
	if topicInfo.BaseTopicItem == nil {
		topicInfo.BaseTopicItem = &spec.TopicItem{}
	}
	if topicInfo.BaseTopicItem.MapOfAnythingValues == nil {
		topicInfo.BaseTopicItem.MapOfAnythingValues = map[string]interface{}{}
	}

	if amqpInfo.VHost != "" {
		topicInfo.BaseTopicItem.MapOfAnythingValues[VHost] = amqpInfo.VHost
	}
	if amqpInfo.Exchange != "" {
		topicInfo.BaseTopicItem.MapOfAnythingValues[Exchange] = amqpInfo.Exchange
	}
	if amqpInfo.RoutingKey != "" {
		topicInfo.BaseTopicItem.MapOfAnythingValues[RoutingKey] = amqpInfo.RoutingKey
	}

	return topicInfo
}
