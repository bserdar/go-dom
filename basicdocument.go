package dom

// BasicDocument implements DOM document
//
// Implementation is guided by https://dom.spec.whatwg.org/
type BasicDocument struct {
	basicNode
	encoding    string
	contentType string
	url         string
	origin      string
	typ         string
	mode        string
}

var _ Document = &BasicDocument{}

// Returns "#document"
func (doc *BasicDocument) GetNodeName() string { return "#document" }

// Returns DOCUMENT_NODE
func (doc *BasicDocument) GetNodeType() NodeType { return DOCUMENT_NODE }

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

// // Returns true or false value indicating whether or not a node is a
// // descendant of the calling node.
// func (doc *BasicDocument) Contains(node Node) bool {
// 	return contains(doc, node)
// }

// // Returns the object's root
// func (doc *BasicDocument) GetRootNode() Node {
// 	return doc
// }

// // Inserts a Node before the reference node as a child of a
// // specified parent node. Returns the added child (unless newNode is
// // a DocumentFragment, in which case the empty DocumentFragment is
// // returned).
// func (doc *BasicDocument) InsertBefore(newNode, referenceNode Node) Node {
// 	panic(ErrHierarchyRequest)
// }

// // Accepts a namespace URI as an argument and returns a boolean value
// // with a value of true if the namespace is the default namespace on
// // the given node or false if not.
// func (doc *BasicDocument) IsDefaultNamespace(uri string) bool {
// 	el := doc.GetDocumentElement()
// 	if el == nil {
// 		return false
// 	}
// 	return el.IsDefaultNamespace(uri)
// }

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
		nodeDoc.mode != doc.mode {
		return false
	}
	return isEqualNode(doc, node)
}

// // Returns a boolean value indicating whether or not the two nodes are
// // the same (that is, they reference the same object).
// func (doc *BasicDocument) IsSameNode(node Node) bool {
// 	return node == doc
// }

// // Returns a string  containing the prefix for a given namespace
// // URI, if present, and "" if not. When multiple prefixes are
// // possible, the result is implementation-dependent.
// func (doc *BasicDocument) LookupPrefix(prefix string) string {
// 	el := doc.GetDocumentElement()
// 	if el == nil {
// 		return ""
// 	}
// 	return el.LookupPrefix(prefix)
// }

// // Accepts a prefix and returns the namespace URI associated with it
// // on the given node if found (and "" if not). Supplying "" for
// // the prefix will return the default namespace.
// func (doc *BasicDocument) LookupNamespaceURI(uri string) string {
// 	el := doc.GetDocumentElement()
// 	if el == nil {
// 		return ""
// 	}
// 	return el.LookupNamespaceURI(uri)
// }

// // Clean up all the text nodes under this element (merge adjacent,
// // remove empty).
// func (doc *BasicDocument) Normalize() {
// 	el := doc.GetDocumentElement()
// 	if el == nil {
// 		return
// 	}
// 	el.Normalize()
// }

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

// // Returns the Element that is a direct child of the document.
// func (doc *BasicDocument) GetDocumentElement() Element {
// 	first := doc.GetFirstChild()
// 	if first == nil {
// 		return nil
// 	}
// 	return first.(Element)
// }
