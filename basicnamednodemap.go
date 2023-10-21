package dom

import (
	"encoding/xml"
)

type basicNamedNodeMap struct {
	attrs    []Attr
	mapAttrs map[xml.Name]Attr
}

func (m *basicNamedNodeMap) GetLength() int {
	return len(m.attrs)
}

// Returns a Attr, corresponding to the given name.
func (m *basicNamedNodeMap) GetNamedItem(name string) Attr {
	if m.mapAttrs == nil {
		return nil
	}
	return m.mapAttrs[xml.Name{Local: name}]
}

// Returns a Attr identified by a namespace and related local name.
func (m *basicNamedNodeMap) GetNamedItemNS(uri string, name string) Attr {
	if m.mapAttrs == nil {
		return nil
	}
	return m.mapAttrs[xml.Name{Local: name, Space: uri}]
}

// Returns the Attr at the given index, or null if the index is higher or equal to the number of nodes
func (m *basicNamedNodeMap) Item(index int) Attr {
	if index < 0 || index >= len(m.attrs) {
		return nil
	}
	return m.attrs[index]
}

// Removes the Attr identified by the given name
func (m *basicNamedNodeMap) RemoveNamedItem(name string) {
	m.RemoveNamedItemNS("", name)
}

// RemoveNamedItemNS removes the Attr identified by the given name
func (m *basicNamedNodeMap) RemoveNamedItemNS(uri string, name string) {
	if m.mapAttrs == nil {
		return
	}
	attr, exists := m.mapAttrs[xml.Name{Local: name, Space: uri}]
	if !exists {
		return
	}
	m.removeAttr(attr)
}

func (m *basicNamedNodeMap) removeAttr(attr Attr) {
	qname := attr.(*BasicAttr).name.Name
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
func (m *basicNamedNodeMap) setNamedItem(owner Node, a Attr) {
	m.setNamedItemNS(owner, a)
}

// Replaces, or adds, the Attr identified in the map by the given namespace and related local name.
func (m *basicNamedNodeMap) setNamedItemNS(owner Node, attr Attr) {
	if attr.GetOwnerElement() != nil && attr.GetOwnerElement() != owner {
		panic(ErrDOM{
			Typ: INUSE_ATTRIBUTE_ERR,
			Msg: "Attribute already in use",
			Op:  "SetNamedItem",
		})
	}
	if m.mapAttrs == nil {
		m.mapAttrs = make(map[xml.Name]Attr)
	}
	ba := attr.(*BasicAttr)
	qname := ba.name.Name
	existing := m.mapAttrs[qname]
	if existing != nil {
		if existing == attr {
			return
		}
		delete(m.mapAttrs, qname)
		for k := range m.attrs {
			if m.attrs[k] == existing {
				m.attrs[k] = attr
				ba.parent = owner
				return
			}
		}
	}
	m.mapAttrs[qname] = attr
	m.attrs = append(m.attrs, attr)
	ba.parent = owner
}

type BasicNamedNodeMap struct {
	owner *BasicElement
}

var _ NamedNodeMap = &BasicNamedNodeMap{}

func (b *BasicNamedNodeMap) GetLength() int { return b.owner.attributes.GetLength() }

// Returns a Attr, corresponding to the given name.
func (b *BasicNamedNodeMap) GetNamedItem(name string) Attr {
	return b.owner.attributes.GetNamedItem(name)
}

// Replaces, or adds, the Attr identified in the map by the given name.
func (b *BasicNamedNodeMap) SetNamedItem(a Attr) {
	b.owner.attributes.setNamedItem(b.owner, a)
}

// Removes the Attr identified by the given name
func (b *BasicNamedNodeMap) RemoveNamedItem(name string) {
	b.owner.attributes.RemoveNamedItem(name)
}

// Returns the Attr at the given index, or null if the index is higher or equal to the number of nodes
func (b *BasicNamedNodeMap) Item(i int) Attr {
	return b.owner.attributes.Item(i)
}

// Returns a Attr identified by a namespace and related local name.
func (b *BasicNamedNodeMap) GetNamedItemNS(uri string, name string) Attr {
	return b.owner.attributes.GetNamedItemNS(uri, name)
}

// Replaces, or adds, the Attr identified in the map by the given namespace and related local name.
func (b *BasicNamedNodeMap) SetNamedItemNS(a Attr) {
	b.owner.attributes.setNamedItemNS(b.owner, a)
}

func (b *BasicNamedNodeMap) RemoveNamedItemNS(uri string, name string) {
	b.owner.attributes.RemoveNamedItemNS(uri, name)
}
