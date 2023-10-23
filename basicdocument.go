package dom

import (
	"encoding/xml"
)

// BasicDocument implements DOM document
//
// Implementation is guided by https://dom.spec.whatwg.org/
type BasicDocument struct {
	basicNode
}

var _ Document = &BasicDocument{}

func NewDocument() Document {
	ret := &BasicDocument{}
	ret.ownerDocument = ret
	return ret
}

// Returns "#document"
func (doc *BasicDocument) GetNodeName() string { return "#document" }

// Returns DOCUMENT_NODE
func (doc *BasicDocument) GetNodeType() NodeType { return DOCUMENT_NODE }

// Returns a boolean value indicating whether or not the two nodes are
// the same (that is, they reference the same object).
func (doc *BasicDocument) IsSameNode(node Node) bool { return node == doc }

// Creates a new Attr object and returns it.
func (doc *BasicDocument) CreateAttribute(name string) Attr {
	return &BasicAttr{
		basicNode: basicNode{
			ownerDocument: doc,
		},
		name: Name{
			Name: xml.Name{
				Local: name,
			},
		},
	}
}

// Creates a new attribute node in a given namespace and returns it.
func (doc *BasicDocument) CreateAttributeNS(ns string, name string) Attr {
	return &BasicAttr{
		basicNode: basicNode{
			ownerDocument: doc,
		},
		name: Name{
			Name: xml.Name{
				Local: name,
				Space: ns,
			},
		},
	}
}

// Creates a new element with the given tag name.
func (doc *BasicDocument) CreateElement(tag string) Element {
	el := &BasicElement{
		basicNode: basicNode{
			ownerDocument: doc,
		},
		name: Name{
			Name: xml.Name{
				Local: tag,
			},
		},
	}
	return el
}

// Creates a new element with the given tag name and namespace URI.
func (doc *BasicDocument) CreateElementNS(ns string, tag string) Element {
	el := &BasicElement{
		basicNode: basicNode{
			ownerDocument: doc,
		},
		name: Name{
			Name: xml.Name{
				Local: tag,
				Space: ns,
			},
		},
	}
	return el
}

// Creates a new CDATA node and returns it.
func (doc *BasicDocument) CreateCDATASection(text string) CDATASection {
	return &BasicCDataSection{
		basicChardata: basicChardata{
			basicNode: basicNode{
				ownerDocument: doc,
			},
			text: text,
		},
	}
}

// Creates a text node.
func (doc *BasicDocument) CreateTextNode(text string) Text {
	return &BasicText{
		basicChardata: basicChardata{
			basicNode: basicNode{
				ownerDocument: doc,
			},
			text: text,
		},
	}
}

// Creates a comment node.
func (doc *BasicDocument) CreateComment(text string) Comment {
	return &BasicComment{
		basicChardata: basicChardata{
			basicNode: basicNode{
				ownerDocument: doc,
			},
			text: text,
		},
	}
}

// Creates a processing instruction node.
func (doc *BasicDocument) CreateProcessingInstruction(target, data string) ProcessingInstruction {
	return &BasicProcessingInstruction{
		basicChardata: basicChardata{
			basicNode: basicNode{
				ownerDocument: doc,
			},
			text: data,
		},
		target: target,
	}
}

// Clone a Node, and optionally, all of its contents.
//
// Returns the new Node cloned. The cloned node has no parent and is not
// part of the document, until it is added to another node that is
// part of the document, using Node.appendChild() or a similar
// method.
func (doc *BasicDocument) CloneNode(deep bool) Node {
	return doc.cloneNode(nil, deep)
}

func (doc *BasicDocument) cloneNode(_ Document, deep bool) Node {
	ret := NewDocument().(*BasicDocument)
	if deep {
		for child := doc.GetFirstChild(); child != nil; child = child.GetNextSibling() {
			newNode := child.cloneNode(ret, deep)
			ret.AppendChild(newNode)
		}
	}
	return ret
}

// Accepts a namespace URI as an argument and returns a boolean value
// with a value of true if the namespace is the default namespace on
// the given node or false if not.
func (doc *BasicDocument) IsDefaultNamespace(uri string) bool {
	root := doc.GetDocumentElement()
	if root == nil {
		return false
	}

	return root.IsDefaultNamespace(uri)
}

// Returns a boolean value which indicates whether or not two nodes
// are of the same type and all their defining data points match.
func (doc *BasicDocument) IsEqualNode(node Node) bool {
	return isEqualNode(doc, node)
}

// Returns a string  containing the prefix for a given namespace
// URI, if present, and "" if not. When multiple prefixes are
// possible, the result is implementation-dependent.
func (doc *BasicDocument) LookupPrefix(uri string) string {
	el := doc.GetDocumentElement()
	if el == nil {
		return ""
	}
	return el.LookupPrefix(uri)
}

// Accepts a prefix and returns the namespace URI associated with it
// on the given node if found (and "" if not).
func (doc *BasicDocument) LookupNamespaceURI(prefix string) string {
	if prefix == "" {
		return ""
	}
	el, _ := doc.GetDocumentElement().(Element)
	if el == nil {
		return ""
	}
	return el.LookupNamespaceURI(prefix)
}

// Clean up all the text nodes under this element (merge adjacent,
// remove empty).
func (doc *BasicDocument) Normalize() {
	el := doc.GetDocumentElement()
	if el == nil {
		return
	}
	el.Normalize()
}

// Returns the Element that is a direct child of the document.
func (doc *BasicDocument) GetDocumentElement() Element {
	for child := doc.GetFirstChild(); child != nil; child = child.GetNextSibling() {
		el, ok := child.(Element)
		if ok {
			return el
		}
	}
	return nil
}

// GetDocumentType returns the document type node
func (doc *BasicDocument) GetDocumentType() DocumentType {
	for ch := doc.GetFirstChild(); ch != nil; ch = ch.GetNextSibling() {
		if dt, ok := ch.(DocumentType); ok {
			return dt
		}
	}
	return nil
}

func (doc *BasicDocument) InsertBefore(newNode, referenceNode Node) Node {
	panic(ErrHierarchyRequest("InsertBefore", "Cannot insert before a document"))
}

// Append newNode as a child of node
func (doc *BasicDocument) AppendChild(newNode Node) Node {
	if err := validatePreInsertion(newNode, doc, nil, "AppendChild"); err != nil {
		panic(err)
	}
	return insertBefore(doc, newNode, nil)
}

// Remove child from node
func (doc *BasicDocument) RemoveChild(child Node) {
	if child.GetParentNode() != doc {
		panic(ErrDOM{
			Typ: NOT_FOUND_ERR,
			Msg: "Wrong parent",
			Op:  "RemoveChild",
		})
	}
	detachChild(doc, child)
}

// Adopt node from an external document.
func (doc *BasicDocument) AdoptNode(node Node) Node {
	if _, ok := node.(Document); ok {
		panic(ErrDOM{
			Typ: NOT_SUPPORTED_ERR,
			Msg: "Cannot adopt a document",
			Op:  "AdoptNode",
		})
	}
	if node.GetOwnerDocument() == doc {
		return node
	}
	if node.GetParentNode() != nil {
		detachChild(node.GetParentNode(), node)
	}
	type setOwnerSupport interface {
		setOwner(*BasicDocument)
	}
	var setOwner func(Node)
	setOwner = func(nd Node) {
		nd.(setOwnerSupport).setOwner(doc)
		for ch := nd.GetFirstChild(); ch != nil; ch = ch.GetNextSibling() {
			setOwner(ch)
		}
	}
	setOwner(node)
	return node
}
