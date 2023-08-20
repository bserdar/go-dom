package dom

type BasicElement struct {
	basicNode

	attributes       BasicNamedNodeMap
	name             Name
	defaultNamespace string
}

var _ Element = &BasicElement{}

func (el *BasicElement) getDefaultNamespace() string { return el.defaultNamespace }

// Returns HTML uppercased qualified name
func (el *BasicElement) GetNodeName() string { return el.getQName() }

// Returns ELEMENT_NODE
func (el *BasicElement) GetNodeType() NodeType { return ELEMENT_NODE }

func (el *BasicElement) getQName() string {
	return el.name.QName()
}

// Returns a String with the name of the tag for the given element.
func (el *BasicElement) GetTagName() string {
	return el.name.QName()
}

// Returns a string representing the namespace prefix of the element,
// or "" if no prefix is specified.
func (el *BasicElement) GetPrefix() string {
	return el.name.Prefix
}

// A string representing the local part of the qualified name of the
// element.
func (el *BasicElement) GetLocalName() string {
	return el.name.Local
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

// Removes the element from the children list of its parent.
func (el *BasicElement) Remove() {
	detachChild(el.GetParentNode(), el)
}

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

// Removes the named attribute from the current node.
func (el *BasicElement) RemoveAttribute(name string) {
	el.attributes.RemoveNamedItem(name)
}

// Removes the node representation of the named attribute from the
// current node.
func (el *BasicElement) RemoveAttributeNode(attr Attr) {
	if attr.GetParentElement() == el {
		el.attributes.removeAttr(attr)
	}
}

// Removes the attribute with the specified name and namespace, from
// the current node.
func (el *BasicElement) RemoveAttributeNS(uri string, name string) {
	el.attributes.RemoveNamedItemNS(uri, name)
}

// // 	// Replaces the existing children of a Node with a specified new set
// // 	// of children.
// // 	ReplaceChildren(...Node)

// // 	// Replaces the element in the children list of its parent with a
// // 	// set of Node or DOMString objects.
// // 	ReplaceWith(...Node)

// Sets the value of a named attribute of the current node.
func (el *BasicElement) SetAttribute(name string, value string) {
	attr := el.ownerDocument.CreateAttribute(name)
	attr.SetValue(value)
	el.attributes.SetNamedItem(attr)
}

// Sets the value of the attribute with the specified name and
// namespace, from the current node.
func (el *BasicElement) SetAttributeNS(uri string, name string, value string) {
	attr := el.ownerDocument.CreateAttributeNS(uri, name)
	attr.SetValue(value)
	el.attributes.SetNamedItemNS(attr)
}

// Sets the node representation of the named attribute from the
// current node.
func (el *BasicElement) SetAttributeNode(attr Attr) {
	el.attributes.SetNamedItem(attr)
}

// Sets the node representation of the attribute with the specified
// name and namespace, from the current node.
func (el *BasicElement) SetAttributeNodeNS(attr Attr) {
	el.attributes.SetNamedItemNS(attr)
}

func (el *BasicElement) InsertBefore(newNode, referenceNode Node) Node {
	if err := validatePreInsertion(newNode, el, referenceNode, "InsertBefore"); err != nil {
		panic(err)
	}
	return insertBefore(el, newNode, referenceNode)
}

// Append newNode as a child of node
func (el *BasicElement) AppendChild(newNode Node) Node {
	if err := validatePreInsertion(newNode, el, nil, "AppendChild"); err != nil {
		panic(err)
	}
	return insertBefore(el, newNode, nil)
}

// Remove child from node
func (el *BasicElement) RemoveChild(child Node) {
	if child.GetParentNode() != el {
		panic(ErrDOM{
			Typ: NOT_FOUND_ERR,
			Msg: "Wrong parent",
			Op:  "RemoveChild",
		})
	}
	detachChild(el, child)
}
