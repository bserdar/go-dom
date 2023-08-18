package dom

type Element interface {
	Node

	GetAttributes() NamedNodeMap

	// Returns a String with the name of the tag for the given element.
	GetTagName() string

	// Returns a string representing the namespace prefix of the element,
	// or "" if no prefix is specified.
	GetPrefix() string

	// A string representing the local part of the qualified name of the
	// element.
	GetLocalName() string

	// The namespace URI of the element, or "" if it is no namespace.}
	//	GetNamespaceURI() string

	GetFirstElementChild() Element
	GetLastElementChild() Element
	GetNextElementSibling() Element
	GetPreviousElementSibling() Element

	// Retrieves the value of the named attribute from the current node
	// and returns it as a string.
	GetAttribute(string) (string, bool)

	// Returns an array of attribute names from the current element.
	GetAttributeNames() []string

	// Retrieves the node representation of the named attribute from the
	// current node and returns it as an Attr.
	GetAttributeNode(string) Attr

	// Retrieves the node representation of the attribute with the
	// specified name and namespace, from the current node and returns
	// it as an Attr.
	GetAttributeNodeNS(uri string, name string) Attr

	//	Retrieves the value of the attribute with the specified
	//	namespace and name from the current node and returns it as a
	//	string.
	GetAttributeNS(uri string, name string) (string, bool)

	// Returns a boolean value indicating if the element has the
	// specified attribute or not.
	HasAttribute(string) bool

	// Returns a boolean value indicating if the element has the
	// specified attribute, in the specified namespace, or not.
	HasAttributeNS(uri string, name string) bool

	// Removes the element from the children list of its parent.
	Remove()

	// Removes the named attribute from the current node.
	RemoveAttribute(string)

	// Removes the node representation of the named attribute from the
	// current node.
	RemoveAttributeNode(Attr)

	// Removes the attribute with the specified name and namespace, from
	// the current node.
	RemoveAttributeNS(uri string, name string)

	// Replaces the existing children of a Node with a specified new set
	// of children.
	//	ReplaceChildren(...Node)

	// Replaces the element in the children list of its parent with a
	// set of Node or DOMString objects.
	//	ReplaceWith(...Node)

	// Sets the value of a named attribute of the current node.
	SetAttribute(name string, value string)

	// Sets the node representation of the named attribute from the
	// current node.
	SetAttributeNode(attr Attr)

	// Sets the node representation of the attribute with the specified
	// name and namespace, from the current node.
	SetAttributeNodeNS(attr Attr)

	// Sets the value of the attribute with the specified name and
	// namespace, from the current node.
	SetAttributeNS(uri string, name string, value string)
}
