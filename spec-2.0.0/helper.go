package spec

func (i *AsyncAPI) WithChannelsItem(name string, value ChannelItem) {
	if i.Channels == nil {
		i.Channels = make(map[string]ChannelItem, 1)
	}

	i.Channels[name] = value
}

func (i *AsyncAPI) ComponentsEns() *Components {
	if i.Components == nil {
		i.Components = &Components{}
	}

	return i.Components
}

func (c *Components) WithMessagesItem(name string, value Message) *Components {
	if c.Messages == nil {
		c.Messages = make(map[string]Message, 1)
	}

	c.Messages[name] = value

	return c
}
