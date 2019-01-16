package amqp_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/swaggest/go-asyncapi/amqp"
)

func TestMessageWithInfo(t *testing.T) {
	m := amqp.MessageWithInfo(nil, amqp.Info{
		Exchange:   "some-exchange",
		RoutingKey: "some-key",
		VHost:      "some-vhost",
	})

	assert.Equal(t, "AMQP VHost: some-vhost, AMQP Exchange: some-exchange, AMQP RoutingKey: some-key", m.Description)
	assert.Equal(t, "some-exchange", m.MapOfAnythingValues[amqp.Exchange])
	assert.Equal(t, "some-key", m.MapOfAnythingValues[amqp.RoutingKey])
	assert.Equal(t, "some-vhost", m.MapOfAnythingValues[amqp.VHost])
}
