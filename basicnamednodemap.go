package dom

type BasicNamedNodeMap struct {
	owner *BasicElement
	attrs []Attr
	// Map of qname -> Attr
	mapAttrs map[string]Attr
}

var _ NamedNodeMap = &BasicNamedNodeMap{}

func (m *BasicNamedNodeMap) GetLength() int {
	return len(m.attrs)
}

// Returns a Attr, corresponding to the given name.
func (m *BasicNamedNodeMap) GetNamedItem(name string) Attr {
	return m.mapAttrs[name]
}

// Returns a Attr identified by a namespace and related local name.
func (m *BasicNamedNodeMap) GetNamedItemNS(uri string, name string) Attr {
	qname := Name{
		ns:    uri,
		local: name,
	}
	return m.mapAttrs[qname.QName()]
}

// Returns the Attr at the given index, or null if the index is higher or equal to the number of nodes
func (m *BasicNamedNodeMap) Item(index int) Attr {
	if index < 0 || index >= len(m.attrs) {
		return nil
	}
	return m.attrs[index]
}

// Removes the Attr identified by the given name
func (m *BasicNamedNodeMap) RemoveNamedItem(name string) {
	attr, exists := m.mapAttrs[name]
	if !exists {
		return
	}
	ba := attr.(*BasicAttr)
	m.RemoveNamedItemNS(ba.name.ns, ba.name.local)
}

// RemoveNamedItemNS removes the Attr identified by the given name
func (m *BasicNamedNodeMap) RemoveNamedItemNS(uri string, name string) {
	nm := Name{
		ns:    uri,
		local: name,
	}
	qname := nm.QName()
	attr, exists := m.mapAttrs[qname]
	if !exists {
		return
	}
	m.removeAttr(attr)
}

func (m *BasicNamedNodeMap) removeAttr(attr Attr) {
	qname := attr.(*BasicAttr).name.QName()
	delete(m.mapAttrs, qname)
	w := 0
	for k := range m.attrs {
		if m.attrs[k] != attr {
			m.attrs[w] = m.attrs[k]
			w++
		}
	}
	m.attrs = m.attrs[:w]
}

// Replaces, or adds, the Attr identified in the map by the given name.
func (m *BasicNamedNodeMap) SetNamedItem(a Attr) {
	m.SetNamedItemNS(a)
}

// Replaces, or adds, the Attr identified in the map by the given namespace and related local name.
func (m *BasicNamedNodeMap) SetNamedItemNS(attr Attr) {
	if attr.GetOwnerElement() != nil && attr.GetOwnerElement() != m.owner {
		panic(ErrDOM{
			Typ: INUSE_ATTRIBUTE_ERR,
			Msg: "Attribute already in use",
			Op:  "SetNamedItem",
		})
	}
	ba := attr.(*BasicAttr)
	qname := ba.name.QName()
	existing := m.mapAttrs[qname]
	if existing != nil {
		if existing == attr {
			return
		}
		delete(m.mapAttrs, qname)
		for k := range m.attrs {
			if m.attrs[k] == existing {
				m.attrs[k] = attr
				ba.parent = m.owner
				return
			}
		}
	}
	m.mapAttrs[qname] = attr
	m.attrs = append(m.attrs, attr)
	ba.parent = m.owner
}
