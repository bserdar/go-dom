package dom

type BasicAttr struct {
	BasicNode

	name  QName
	space string
	value string
}

var _ Attr = &BasicAttr{}

// A String representing the local part of the qualified name of the
// attribute.
func (attr *BasicAttr) GetLocalName() string {
	return attr.name.Local
}

// Returns the qualified name of an attribute, that is the name of
// the attribute, with the namespace prefix, if any, in front of
// it. For example, if the local name is lang and the namespace
// prefix is xml, the returned qualified name is xml:lang.
func (attr *BasicAttr) GetName() string {
	return attr.name.String()
}

// A String representing the URI of the namespace of the attribute, or
// null if there is no namespace.
func (attr *BasicAttr) GetNamespaceURI() string {
	return attr.space
}

// The Element the attribute belongs to.
func (attr *BasicAttr) GetOwnerElement() Element {
	if el, ok := attr.parent.(Element); ok {
		return el
	}
	return nil
}

// A String representing the namespace prefix of the attribute, or
// null if a namespace without prefix or no namespace are specified.
func (attr *BasicAttr) GetPrefix() string {
	return attr.name.Prefix
}

// The attribute's value, a string that can be set and get using this
// property.
func (attr *BasicAttr) GetValue() string {
	return attr.value
}

func (attr *BasicAttr) SetValue(v string) {
	attr.value = v
}

func (attr *BasicAttr) CloneNode(bool) Node {
	ret := BasicAttr{
		name:  attr.name,
		space: attr.space,
		value: attr.value,
	}
	return &ret
}

func (attr *BasicAttr) GetNodeName() string {
	return attr.name.String()
}

func (attr *BasicAttr) GetNodeType() NodeType { return ATTRIBUTE_NODE }
