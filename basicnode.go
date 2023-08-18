package dom

// basicNode implements the common functionality for DOM nodes
type basicNode struct {
	tnode
	ownerDocument    *BasicDocument
	defaultNamespace string
}

// Returns true or false value indicating whether or not a node is a
// descendant of the calling node.
func contains(parent, childNode Node) bool {
	trc := childNode
	for {
		if trc == parent {
			return true
		}
		trc = trc.GetParentNode()
	}
	return false
}

// Returns the object's root
func getRootNode(node Node) Node {
	trc := node
	for {
		parent := trc.GetParentNode()
		if parent == nil {
			return trc
		}
		trc = parent
	}
}

func (node *basicNode) treeNode() *tnode { return &node.tnode }

func (node *basicNode) GetNodeName() string {
	panic("basicNode.GetNodeName should not have been called")
}

func (node *basicNode) GetNodeType() NodeType {
	panic("basicNode.GetNodeType should not have been called")
}

func (node *basicNode) IsEqualNode(Node) bool {
	panic("basicNode.IsEqualNode should not have been called")
}

// Returns a boolean value indicating whether or not the element has
// any child nodes.
func (node *basicNode) HasChildNodes() bool {
	return node.child != nil
}

// Clean up all the text nodes under this element (merge adjacent,
// remove empty).
func (node *basicNode) Normalize() {
	for child := node.GetFirstChild(); child != nil; child = child.GetNextSibling() {
		child.Normalize()
	}
}

// Returns a Node representing the first direct child node of the
// node, or null if the node has no child.
func (node *basicNode) GetFirstChild() Node {
	return node.firstChild()
}

// Returns a Node representing the last direct child node of the node,
// or null if the node has no child.
func (node *basicNode) GetLastChild() Node {
	return node.lastChild()
}

// Returns a Node representing the next node in the tree, or null if
// there isn't such node.
func (node *basicNode) GetNextSibling() Node {
	return node.nextSibling()
}

// Returns the Document that this node belongs to. If the node is
// itself a document, returns null.
func (node *basicNode) GetOwnerDocument() Document {
	return node.ownerDocument
}

// Returns a Node that is the parent of this node. If there is no such
// node, like if this node is the top of the tree or if doesn't
// participate in a tree, this property returns null.
func (node *basicNode) GetParentNode() Node {
	return node.parent
}

// Returns an Element that is the parent of this node. If the node has
// no parent, or if that parent is not an Element, this property
// returns null.
func (node *basicNode) GetParentElement() Element {
	for trc := node.parent; trc != nil; trc = trc.GetParentNode() {
		if doc, ok := trc.(Element); ok {
			return doc
		}
	}
	return nil
}

// Returns a Node representing the previous node in the tree, or null
// if there isn't such node.
func (node *basicNode) GetPreviousSibling() Node {
	return node.prevSibling()
}

// Returns a live NodeList containing all the children of this node
// (including elements, text and comments). NodeList being live means
// that if the children of the Node change, the NodeList object is
// automatically updated.
func (node *basicNode) GetChildNodes() NodeList {
	return newBasicNodeList(node)
}

// Inserts a Node before the reference node as a child of a
// specified parent node. Returns the added child (unless newNode is
// a DocumentFragment, in which case the empty DocumentFragment is
// returned).
//
// Algorithm: To ensure pre-insertion validity of a node into a parent
// before a child, run these steps:
//
// If parent is not a Document, DocumentFragment, or Element node,
// then throw a "HierarchyRequestError" DOMException.
//
// If node is a host-including inclusive ancestor of parent, then
// throw a "HierarchyRequestError" DOMException.
//
// If child is non-null and its parent is not parent, then throw a
// "NotFoundError" DOMException.
//
// If node is not a DocumentFragment, DocumentType, Element, or
// CharacterData node, then throw a "HierarchyRequestError"
// DOMException.
//
// If either node is a Text node and parent is a document, or node is
// a doctype and parent is not a document, then throw a
// "HierarchyRequestError" DOMException.
//
// If parent is a document, and any of the statements below, switched
// on the interface node implements, are true, then throw a
// "HierarchyRequestError" DOMException.
//
// DocumentFragment: If node has more than one element child or has a
// Text node child. Otherwise, if node has one element child and
// either parent has an element child, child is a doctype, or child is
// non-null and a doctype is following child.
//
// Element: parent has an element child, child is a doctype, or child
// is non-null and a doctype is following child.
//
// DocumentType: parent has a doctype child, child is non-null and an
// element is preceding child, or child is null and parent has an
// element child.
//
// To pre-insert a node into a parent before a child, run these steps:

// Ensure pre-insertion validity of node into parent before child.

// Let referenceChild be child.

// If referenceChild is node, then set referenceChild to node’s next sibling.

// Insert newNode into node before referenceChild.

// Return node.
func (node *basicNode) InsertBefore(newNode, referenceNode Node) (Node, error) {
	if err := validatePreInsertion(newNode, node, referenceNode, "InsertBefore"); err != nil {
		return nil, err
	}
	return insertBefore(node, newNode, referenceNode), nil
}

// Append newNode as a child of node
func (node *basicNode) AppendChild(newNode Node) (Node, error) {
	if err := validatePreInsertion(newNode, node, nil, "AppendChild"); err != nil {
		return nil, err
	}
	return insertBefore(node, newNode, nil), nil
}

// Remove child from node
func (node *basicNode) RemoveChild(child Node) error {
	if child.GetParentNode() != node {
		return ErrDOM{
			Typ: NOT_FOUND_ERR,
			Msg: "Wrong parent",
			Op:  "RemoveChild",
		}
	}
	detachChild(node, child)
	return nil
}

// Inserts a Node before the reference node as a child of a
// specified parent node. Returns the added child (unless newNode is
// a DocumentFragment, in which case the empty DocumentFragment is
// returned).
func insertBefore(parent, newNode, referenceNode Node) Node {
	// Document fragment?
	if frag, ok := newNode.(DocumentFragment); ok {
		for child := frag.GetFirstChild(); child != nil; {
			detachChild(frag, child)
			insertChildBefore(parent, child, nil)
		}
		return frag
	}
	detachChild(parent, newNode)
	insertChildBefore(parent, newNode, nil)
	return newNode
}

// Compares the children of each node
func isEqualNode(n1, n2 Node) bool {
	n1child := n1.GetFirstChild()
	n2child := n2.GetFirstChild()
	if n1child == nil && n2child == nil {
		return true
	}
	if n1child == nil || n2child == nil {
		return false
	}
	for {
		if !n1child.IsEqualNode(n2child) {
			return false
		}
		n1child = n1child.GetNextSibling()
		n2child = n2child.GetNextSibling()
		if n1child == nil && n2child == nil {
			return true
		}
		if n1child == nil || n2child == nil {
			return false
		}
	}
}

func hasChildOfType(parent Node, typ NodeType) bool {
	for child := parent.GetFirstChild(); child != nil; child = child.GetNextSibling() {
		if child.GetNodeType() == typ {
			return true
		}
	}
	return false
}

func validatePreInsertion(node, parent, beforeChild Node, op string) error {
	if node.GetOwnerDocument() != parent.GetOwnerDocument() {
		panic("Child node does not belong to the target document. Use AdoptNode to adopt it.")
	}
	if beforeChild != nil && beforeChild.GetOwnerDocument() != parent.GetOwnerDocument() {
		panic("ReferenceNode node does not belong to the target document.")
	}
	parentType := parent.GetNodeType()
	nodeType := node.GetNodeType()
	if parentType != DOCUMENT_NODE && parentType != DOCUMENT_FRAGMENT_NODE && parentType != ELEMENT_NODE {
		return ErrHierarchyRequest(op, "Parent is not a DOCUMENT, DOCUMENT_FRAGMENT, or ELEMENT")
	}
	if nodeType != DOCUMENT_FRAGMENT_NODE ||
		nodeType != DOCUMENT_TYPE_NODE ||
		nodeType != ELEMENT_NODE ||
		nodeType != CDATA_SECTION_NODE ||
		nodeType != TEXT_NODE ||
		nodeType != COMMENT_NODE {
		return ErrHierarchyRequest(op, "Invalid node type")
	}
	if beforeChild != nil && beforeChild.GetParentNode() != node.GetParentNode() {
		return ErrDOM{
			Typ: NOT_FOUND_ERR,
			Msg: "Reference node not found in parent",
			Op:  op,
		}
	}
	if nodeType == DOCUMENT_TYPE_NODE && parentType != DOCUMENT_NODE {
		return ErrHierarchyRequest(op, "Document type node must be under document node")
	}
	if parentType == DOCUMENT_NODE {
		if beforeChild != nil && beforeChild.GetNodeType() == TEXT_NODE {
			return ErrHierarchyRequest(op, "Text reference node under document node is not allowed")
		}
		switch nodeType {
		case TEXT_NODE:
			return ErrHierarchyRequest(op, "Text under document node is not allowed")
		case DOCUMENT_FRAGMENT_NODE:
			nElementChild := 0
			for child := node.GetFirstChild(); child != nil; child = child.GetNextSibling() {
				switch child.GetNodeType() {
				case ELEMENT_NODE:
					nElementChild++
					if nElementChild > 1 {
						return ErrHierarchyRequest(op, "Attempting to add multiple elements to a document node")
					}
				case TEXT_NODE:
					return ErrHierarchyRequest(op, "Attempting to add text to a document node")
				}
			}
			if nElementChild == 1 {
				if hasChildOfType(parent, ELEMENT_NODE) && beforeChild != nil && beforeChild.GetNodeType() == DOCUMENT_TYPE_NODE {
					return ErrHierarchyRequest(op, "Invalid fragment")
				}
				if beforeChild != nil {
					for x := beforeChild; x != nil; x = x.GetNextSibling() {
						if x.GetNodeType() == DOCUMENT_TYPE_NODE {
							return ErrHierarchyRequest(op, "Invalid fragment")
						}
					}
				}
			}
		case ELEMENT_NODE:
			if hasChildOfType(parent, ELEMENT_NODE) && beforeChild != nil && beforeChild.GetNodeType() == DOCUMENT_TYPE_NODE {
				return ErrHierarchyRequest(op, "Invalid element")
			}
			if beforeChild != nil {
				for x := beforeChild; x != nil; x = x.GetNextSibling() {
					if x.GetNodeType() == DOCUMENT_TYPE_NODE {
						return ErrHierarchyRequest(op, "Invalid fragment")
					}
				}
			}
		case DOCUMENT_TYPE_NODE:
			if hasChildOfType(parent, DOCUMENT_TYPE_NODE) && beforeChild != nil {
				for x := beforeChild; x != nil; x.GetPreviousSibling() {
					if x.GetNodeType() == ELEMENT_NODE {
						return ErrHierarchyRequest(op, "Invalid document type node")
					}
				}
			} else {
				// beforeChild is nil
				for child := parent.GetFirstChild(); child != nil; child = child.GetNextSibling() {
					if child.GetNodeType() == ELEMENT_NODE {
						return ErrHierarchyRequest(op, "Invalid document type node")
					}
				}
			}
		}
	}
	return nil
}