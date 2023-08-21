package spec

// AddServer adds named server.
func (i *AsyncAPI) AddServer(name string, srv Server) {
	if i.Servers == nil {
		i.Servers = make(map[string]ServerOrRef)
	}

	i.Servers[name] = ServerOrRef{
		Server: &srv,
	}
}

// WithVariable adds server variable.
func (s Server) WithVariable(name string, v ServerVariable) Server {
	s.WithVariablesItem(name, ServerVariableOrRef{
		ServerVariable: &v,
	})

	return s
}
