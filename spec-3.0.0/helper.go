package spec

// AddServer adds named server.
func (i *AsyncAPI) AddServer(name string, prepare func(srv *Server)) {
	if i.Servers == nil {
		i.Servers = make(map[string]ServerOrRef)
	}

	srv := Server{}
	prepare(&srv)

	i.Servers[name] = ServerOrRef{
		Server: &srv,
	}
}

// AddVariable adds named server variable.
func (s *Server) AddVariable(key string, prepare func(v *ServerVariable)) {
	v := ServerVariable{}
	prepare(&v)

	s.WithVariablesItem(key, ServerVariableOrRef{
		ServerVariable: &v,
	})
}

// UpdateInfo updates API info.
func (i *AsyncAPI) UpdateInfo(prepare func(inf *Info)) {
	prepare(i.Info.InfoEns())
}
