package dom

// type BasicNamedNodeMap struct {
// 	owner    *BasicElement
// 	attrs    []Attr
// 	mapAttrs map[QName]Attr
// }

// var _ NamedNodeMap = &BasicNamedNodeMap{}

// func (m *BasicNamedNodeMap) GetLength() int {
// 	return len(m.attrs)
// }

// // Returns a Attr, corresponding to the given name.
// func (m *BasicNamedNodeMap) GetNamedItem(name string) Attr {
// 	qname := ParseQName(name)
// 	return m.mapAttrs[qname]
// }

// // Replaces, or adds, the Attr identified in the map by the given name.
// func (m *BasicNamedNodeMap) SetNamedItem(a Attr) {
// 	ba := a.(*BasicAttr)
// 	ba.detach()
// 	old := m.mapAttrs[ba.name]
// 	if old == nil {
// 		m.attrs = append(m.attrs, a)
// 		m.mapAttrs[ba.name] = a
// 		ba.setParent(m.owner)
// 		return
// 	}
// 	m.mapAttrs[ba.Name] = a
// 	for k := range m.attrs {
// 		if m.attrs[k].(*BasicAttr).name == ba.name {
// 			m.attrs[k] = ba
// 			ba.setParent(m.owner)
// 			break
// 		}
// 	}
// }

// // Removes the Attr identified by the given name
// func (m *BasicNamedNodeMap) RemoveNamedItem(name string) {
// 	qname := ParseQName(name)
// 	delete(m.mapAttrs, qname)
// 	w := 0
// 	for k := range m.attrs {
// 		if m.attrs[k].(*BasicAttr).name != qname {
// 			m.attrs[w] = m.attrs[k]
// 			w++
// 		}
// 	}
// 	m.attrs = m.attrs[:w]
// }

// // Returns the Attr at the given index, or null if the index is higher or equal to the number of nodes
// func (m *BasicNamedNodeMap) Item(index int) Attr {
// 	if index < 0 || index >= len(m.attrs) {
// 		return nil
// 	}
// 	return m.attrs[index]
// }

// // Returns a Attr identified by a namespace and related local name.
// func (m *BasicNamedNodeMap) GetNamedItemNS(uri string, name string) Attr {
// 	for _, k := range m.mapAttrs {
// 		if k.GetNamespaceURI() == uri && k.GetLocalName() == name {
// 			return k
// 		}
// 	}
// 	return nil
// }

// // Replaces, or adds, the Attr identified in the map by the given namespace and related local name.
// func (m *BasicNamedNodeMap) SetNamedItemNS(attr Attr) {
// 	ba := attr.(*BasicAttr)
// 	ba.detach()
// 	found := false
// 	for k, v := range m.mapAttrs {
// 		if v.GetNamespaceURI() == ba.GetNamespaceURI() && v.GetLocalName() == attr.GetLocalName() {
// 			delete(m.mapAttrs, k)
// 			found = true
// 		}
// 	}
// 	m.mapAttrs[ba.name] = ba
// 	if !found {
// 		m.attrs = append(m.attrs, attr)
// 		return
// 	}
// 	for k := range m.attrs {
// 		if m.attrs[k].GetNamespaceURI() == ba.GetNamespaceURI() && m.attrs[k].GetLocalName() == attr.GetLocalName() {
// 			m.attrs[k] = ba
// 			break
// 		}
// 	}
// }

// func (m *BasicNamedNodeMap) RemoveNamedItemNS(uri string, name string) {
// 	found := false
// 	for k, v := range m.mapAttrs {
// 		if v.GetNamespaceURI() == uri && v.GetLocalName() == name {
// 			delete(m.mapAttrs, k)
// 			found = true
// 		}
// 	}
// 	if found {
// 		w := 0
// 		for k := range m.attrs {
// 			if m.attrs[k].GetNamespaceURI() != uri || m.attrs[k].GetLocalName() != name {
// 				m.attrs[w] = m.attrs[k]
// 				w++
// 			}
// 		}
// 		m.attrs = m.attrs[:w]
// 	}
// }
