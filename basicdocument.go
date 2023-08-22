package dom

import (
	"encoding/xml"
)

// BasicDocument implements DOM document
//
// Implementation is guided by https://dom.spec.whatwg.org/
type BasicDocument struct {
	basicNode

	encoding         string
	contentType      string
	url              string
	origin           string
	typ              string
	mode             string
	defaultNamespace string
}

var _ Document = &BasicDocument{}

func (doc *BasicDocument) getDefaultNamespace() string { return doc.defaultNamespace }

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
	return &BasicElement{
		basicNode: basicNode{
			ownerDocument: doc,
		},
		name: Name{
			Name: xml.Name{
				Local: tag,
			},
		},
	}
}

// Creates a new element with the given tag name and namespace URI.
func (doc *BasicDocument) CreateElementNS(ns string, tag string) Element {
	return &BasicElement{
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

// // Clone a Node, and optionally, all of its contents.
// //
// // Returns the new Node cloned. The cloned node has no parent and is not
// // part of the document, until it is added to another node that is
// // part of the document, using Node.appendChild() or a similar
// // method.
// func (doc *BasicDocument) CloneNode(deep bool) Node {
// 	ret := &BasicDocument{
// 		encoding:    doc.encoding,
// 		contentType: doc.contentType,
// 		url:         doc.url,
// 		origin:      doc.origin,
// 		typ:         doc.typ,
// 		mode:        doc.mode,
// 	}
// 	if deep {
// 		for child := doc.GetFirstChild(); child != nil; child = child.GetNextSibling() {
// 			cloneNode(ret, child, deep)
// 		}
// 	}
// 	return ret
// }

// Accepts a namespace URI as an argument and returns a boolean value
// with a value of true if the namespace is the default namespace on
// the given node or false if not.
func (doc *BasicDocument) IsDefaultNamespace(uri string) bool {
	return uri == doc.defaultNamespace
}

// Returns a boolean value which indicates whether or not two nodes
// are of the same type and all their defining data points match.
func (doc *BasicDocument) IsEqualNode(node Node) bool {
	nodeDoc, ok := node.(*BasicDocument)
	if !ok {
		return false
	}
	if nodeDoc.encoding != doc.encoding ||
		nodeDoc.contentType != doc.contentType ||
		nodeDoc.url != doc.url ||
		nodeDoc.origin != doc.origin ||
		nodeDoc.typ != doc.typ ||
		nodeDoc.mode != doc.mode ||
		nodeDoc.defaultNamespace != doc.defaultNamespace {
		return false
	}
	return isEqualNode(doc, node)
}

// Returns a string  containing the prefix for a given namespace
// URI, if present, and "" if not. When multiple prefixes are
// possible, the result is implementation-dependent.
func (doc *BasicDocument) LookupPrefix(uri string) string {
	el, _ := doc.GetDocumentElement().(*BasicElement)
	if el == nil {
		return ""
	}
	s, _ := el.getPrefix(uri)
	return s
}

// Accepts a prefix and returns the namespace URI associated with it
// on the given node if found (and "" if not). Supplying "" for
// the prefix will return the default namespace.
func (doc *BasicDocument) LookupNamespaceURI(prefix string) string {
	if prefix == "" {
		return doc.defaultNamespace
	}
	el, _ := doc.GetDocumentElement().(*BasicElement)
	if el == nil {
		return ""
	}
	s, _ := el.getNS(prefix)
	return s
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

// // Replaces one child Node of the current one with the second one
// // given in parameter.
// func (doc *BasicDocument) ReplaceChild(newChild, oldChild Node) {
// 	if newChild.GetOwnerDocument() != doc || newChild == doc {
// 		panic(ErrHierarchyRequest)
// 	}
// 	if newChild.GetParentNode() != nil {
// 		detachChild(newChild.GetParentNode(), newChild)
// 	}
// 	if frag, ok := newChild.(DocumentFragment); ok {
// 		for child := frag.GetFirstChild(); child != nil; {
// 			detachChild(frag, childNode)
// 			insertChildAfter(parent, child, oldChild)
// 		}
// 		detachChild(doc, oldChild)
// 		return frag
// 	}
// }

// // Returns a live NodeList containing all the children of this node
// // (including elements, text and comments). NodeList being live means
// // that if the children of the Node change, the NodeList object is
// // automatically updated.
// func (doc *BasicDocument) GetChildNodes() NodeList {
// 	return newBasicNodeLisr(doc)
// }

// Returns the Element that is a direct child of the document.
func (doc *BasicDocument) GetDocumentElement() Element {
	first := doc.GetFirstChild()
	if first == nil {
		return nil
	}
	return first.(Element)
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
	if node.GetParentNode() != nil {
		detachChild(node.GetParentNode(), node)
	}
	type setOwnerSupport interface {
		setOwner(*BasicDocument)
	}
	var setParent func(Node)
	setParent = func(nd Node) {
		nd.(setOwnerSupport).setOwner(doc)
		for ch := nd.GetFirstChild(); ch != nil; ch = ch.GetNextSibling() {
			setParent(ch)
		}
	}
	setParent(node)
	return node
}
