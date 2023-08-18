package dom

type BasicElement struct {
	basicNode

	attributes BasicNamedNodeMap

	name Name
}

var _ Element = &BasicElement{}

// Returns HTML uppercased qualified name
func (el *BasicElement) GetNodeName() string { return el.getQName() }

// Returns ELEMENT_NODE
func (el *BasicElement) GetNodeType() NodeType { return ELEMENT_NODE }

func (el *BasicElement) getQName() string {
	return el.name.QName()
}

// Returns a String with the name of the tag for the given element.
func (el *BasicElement) GetTagName() string {
	return el.name.local
}

// Returns a string representing the namespace prefix of the element,
// or "" if no prefix is specified.
func (el *BasicElement) GetPrefix() string {
	return el.name.prefix
}

// A string representing the local part of the qualified name of the
// element.
func (el *BasicElement) GetLocalName() string {
	return el.name.local
}

// // The namespace URI of the element, or "" if it is no namespace.}
// func (el *BasicElement) GetNamespaceURI() string {
// 	return el.space
// }

func (el *BasicElement) GetFirstElementChild() Element {
	return nextElementSibling(el.GetFirstChild())
}

func (el *BasicElement) GetLastElementChild() Element {
	return prevElementSibling(el.GetLastChild())
}

func (el *BasicElement) GetNextElementSibling() Element {
	return nextElementSibling(el.GetNextSibling())
}

func (el *BasicElement) GetPreviousElementSibling() Element {
	return prevElementSibling(el.GetPreviousSibling())
}

func nextElementSibling(start Node) Element {
	if start == nil {
		return nil
	}
	for trc := start; trc != nil; trc = trc.GetNextSibling() {
		if el, ok := trc.(Element); ok {
			return el
		}
	}
	return nil
}

func prevElementSibling(start Node) Element {
	if start == nil {
		return nil
	}
	for trc := start; trc != nil; trc = trc.GetPreviousSibling() {
		if el, ok := trc.(Element); ok {
			return el
		}
	}
	return nil
}

// // Removes the element from the children list of its parent.
// func (el *BasicElement) Remove() {
// 	el.detach()
// }

func (el *BasicElement) GetAttributes() NamedNodeMap {
	return &el.attributes
}

// // 	GetID() string

// Retrieves the value of the named attribute from the current node
// and returns it as a string.
func (el *BasicElement) GetAttribute(name string) (string, bool) {
	attr := el.GetAttributeNode(name)
	if attr == nil {
		return "", false
	}
	return attr.GetValue(), true
}

// Returns an array of attribute names from the current element.
func (el *BasicElement) GetAttributeNames() []string {
	ret := make([]string, el.attributes.GetLength())
	for i := 0; i < len(ret); i++ {
		ret[i] = el.attributes.attrs[i].(*BasicAttr).name.QName()
	}
	return ret
}

// Retrieves the node representation of the named attribute from the
// current node and returns it as an Attr.
func (el *BasicElement) GetAttributeNode(name string) Attr {
	return el.attributes.GetNamedItem(name)
}

// Retrieves the node representation of the attribute with the
// specified name and namespace, from the current node and returns
// it as an Attr.
func (el *BasicElement) GetAttributeNodeNS(uri string, name string) Attr {
	return el.attributes.GetNamedItemNS(uri, name)
}

// Retrieves the value of the attribute with the specified
// namespace and name from the current node and returns it as a
// string.
func (el *BasicElement) GetAttributeNS(uri string, name string) (string, bool) {
	attr := el.GetAttributeNodeNS(uri, name)
	if attr == nil {
		return "", false
	}
	return attr.GetValue(), true
}

// Returns a boolean value indicating if the element has the
// specified attribute or not.
func (el *BasicElement) HasAttribute(name string) bool {
	return el.HasAttribute(name)
}

// Returns a boolean value indicating if the element has the
// specified attribute, in the specified namespace, or not.
func (el *BasicElement) HasAttributeNS(uri string, name string) bool {
	return el.GetAttributeNodeNS(uri, name) != nil
}

// // 	// Removes the named attribute from the current node.
// // 	RemoveAttribute(string)

// // 	// Removes the node representation of the named attribute from the
// // 	// current node.
// // 	RemoveAttributeNode(Attr)

// // 	// Removes the attribute with the specified name and namespace, from
// // 	// the current node.
// // 	RemoveAttributeNS(uri string, name string)

// // 	// Replaces the existing children of a Node with a specified new set
// // 	// of children.
// // 	ReplaceChildren(...Node)

// // 	// Replaces the element in the children list of its parent with a
// // 	// set of Node or DOMString objects.
// // 	ReplaceWith(...Node)

// // 	// Sets the value of a named attribute of the current node.
// // 	SetAttribute(name string, value string)

// // 	// Sets the node representation of the named attribute from the
// // 	// current node.
// // 	SetAttributeNode(attr Attr)

// // 	// Sets the node representation of the attribute with the specified
// // 	// name and namespace, from the current node.
// // 	SetAttributeNodeNS(attr Attr)

// // 	// Sets the value of the attribute with the specified name and
// // 	// namespace, from the current node.
// // 	SetAttributeNS(uri string, name string, value string)
// // }
