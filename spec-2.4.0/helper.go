package spec

// AddServer adds named server.
func (i *AsyncAPI) AddServer(name string, srv Server) {
	if i.Servers == nil {
		i.Servers = make(map[string]ServersAdditionalProperties)
	}

	i.Servers[name] = ServersAdditionalProperties{
		Server: &srv,
	}
}
