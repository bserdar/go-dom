package dom

type nsMap struct {
	// prefix -> namespace
	namespaces map[string]string
	// namespace -> prefix
	prefixes map[string][]string
	parent   *nsMap
}

func newNsMap() *nsMap {
	return &nsMap{
		namespaces: make(map[string]string),
		prefixes:   make(map[string][]string),
	}
}

func (m nsMap) getPrefix(uri string) (string, bool) {
	s := m.prefixes[uri]
	if len(s) > 0 {
		return s[0], true
	}
	if m.parent == nil {
		return "", false
	}
	return m.parent.getPrefix(uri)
}

func (m nsMap) getNS(prefix string) (string, bool) {
	s, ok := m.namespaces[prefix]
	if ok {
		return s, ok
	}
	if m.parent == nil {
		return "", false
	}
	return m.parent.getNS(prefix)
}

func (m nsMap) add(prefix, uri string) error {
	ns, ok := m.namespaces[prefix]
	if ok {
		if ns == uri {
			return nil
		}
		return ErrDOM{
			Typ: NAMESPACE_ERR,
			Msg: "Same prefix with different namespace",
		}
	}
	// Prefix does not exist in the map
	m.namespaces[prefix] = uri
	m.prefixes[uri] = append(m.prefixes[uri], prefix)
	return nil
}
