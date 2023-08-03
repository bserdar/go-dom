package dom

// type BasicElement struct {
// 	BasicNode

// 	attributes BasicNamedNodeMap

// 	space string
// 	name  QName
// }

// var _ Element = &BasicElement{}

// // Returns a String with the name of the tag for the given element.
// func (el *BasicElement) GetTagName() string {
// 	return el.name.String()
// }

// // Returns a string representing the namespace prefix of the element,
// // or "" if no prefix is specified.
// func (el *BasicElement) GetPrefix() string {
// 	return el.name.Prefix
// }

// // A string representing the local part of the qualified name of the
// // element.
// func (el *BasicElement) GetLocalName() string {
// 	return el.name.Prefix
// }

// // The namespace URI of the element, or "" if it is no namespace.}
// func (el *BasicElement) GetNamespaceURI() string {
// 	return el.space
// }

// func (el *BasicElement) GetFirstElementChild() Element {
// 	return nextElementSibling(el.firstChild)
// }

// func (el *BasicElement) GetLastElementChild() Element {
// 	return prevElementSibling(el.lastChild)
// }

// func (el *BasicElement) GetNextElementSibling() Element {
// 	return nextElementSibling(el.nextSibling)
// }

// func (el *BasicElement) GetPreviousElementSibling() Element {
// 	return prevElementSibling(el.prevSibling)
// }

// func nextElementSibling(start treeNode) Element {
// 	for trc := start; trc != nil; trc = trc.getNextSibling() {
// 		if el, ok := trc.(Element); ok {
// 			return el
// 		}
// 	}
// 	return nil
// }

// func prevElementSibling(start treeNode) Element {
// 	for trc := start; trc != nil; trc = trc.getPrevSibling() {
// 		if el, ok := trc.(Element); ok {
// 			return el
// 		}
// 	}
// 	return nil
// }

// // Removes the element from the children list of its parent.
// func (el *BasicElement) Remove() {
// 	el.detach()
// }

// func (el *BasicElement) GetAttributes() NamedNodeMap {
// 	return &el.attributes
// }

// // 	GetID() string

// // 	// Retrieves the value of the named attribute from the current node
// // 	// and returns it as a string.
// // 	GetAttribute(string) (string, bool)

// // 	// Returns an array of attribute names from the current element.
// // 	GetAttributeNames() []string

// // 	// Retrieves the node representation of the named attribute from the
// // 	// current node and returns it as an Attr.
// // 	GetAttributeNode(string) Attr

// // 	// Retrieves the node representation of the attribute with the
// // 	// specified name and namespace, from the current node and returns
// // 	// it as an Attr.
// // 	GetAttributeNodeNS(uri string, name string) Attr

// // 	//	Retrieves the value of the attribute with the specified
// // 	//	namespace and name from the current node and returns it as a
// // 	//	string.
// // 	GetAttributeNS(uri string, name string) (string, bool)

// // 	// Returns a boolean value indicating if the element has the
// // 	// specified attribute or not.
// // 	HasAttribute(string) bool

// // 	// Returns a boolean value indicating if the element has the
// // 	// specified attribute, in the specified namespace, or not.
// // 	HasAttributeNS(uri string, name string) bool

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
