package dom

// BasicNode implements the common functionality for DOM nodes
type BasicNode struct {
	parent      treeNode
	firstChild  treeNode
	lastChild   treeNode
	nextSibling treeNode
	prevSibling treeNode

	// childListVer is incremented every time the list of children is
	// modified
	childListVer int
}

var _ treeNode = &BasicNode{}

// treeNode is an internal interface implemented by all Basic node
// types.
type treeNode interface {
	insertAfter(newChild, after treeNode)
	insertBefore(newChild, before treeNode)
	detach()

	getChildListVer() int
	setFirstChild(treeNode)
	setLastChild(treeNode)
	setNextSibling(treeNode)
	setPrevSibling(treeNode)
	setParent(treeNode)

	incChildListVer()
	getFirstChild() treeNode
	getLastChild() treeNode
	getNextSibling() treeNode
	getPrevSibling() treeNode
	getParent() treeNode
}

func (node *BasicNode) getChildListVer() int      { return node.childListVer }
func (node *BasicNode) getFirstChild() treeNode   { return node.firstChild }
func (node *BasicNode) getLastChild() treeNode    { return node.lastChild }
func (node *BasicNode) getPrevSibling() treeNode  { return node.prevSibling }
func (node *BasicNode) getNextSibling() treeNode  { return node.nextSibling }
func (node *BasicNode) getParent() treeNode       { return node.parent }
func (node *BasicNode) incChildListVer()          { node.childListVer++ }
func (node *BasicNode) setFirstChild(n treeNode)  { node.firstChild = n }
func (node *BasicNode) setLastChild(n treeNode)   { node.lastChild = n }
func (node *BasicNode) setPrevSibling(n treeNode) { node.prevSibling = n }
func (node *BasicNode) setNextSibling(n treeNode) { node.nextSibling = n }

func (node *BasicNode) setParent(n treeNode) {
	if node.parent != nil {
		node.parent.incChildListVer()
	}
	node.parent = n
	if node.parent != nil {
		node.parent.incChildListVer()
	}
}

// inserts the new child after a node. If after is nil, inserts at the
// beginning. The child must be detached
func (node *BasicNode) insertAfter(newChild, after treeNode) {
	newChild.setParent(node)
	newChild.setPrevSibling(after)
	if after == nil {
		newChild.setNextSibling(node.getFirstChild())
		node.setFirstChild(newChild)
	} else {
		newChild.setNextSibling(after.getNextSibling())
		after.setNextSibling(newChild)
	}
	if newChild.getNextSibling() != nil {
		newChild.getNextSibling().setPrevSibling(newChild)
	}
	if node.getLastChild() == nil {
		node.setLastChild(node.getFirstChild())
	}
}

// inserts the new child before a node. If before is nil, inserts at the
// end. The child must be detached
func (node *BasicNode) insertBefore(newChild, before treeNode) {
	newChild.setParent(node)
	newChild.setNextSibling(before)
	if before == nil {
		newChild.setPrevSibling(node.getLastChild())
		node.setLastChild(newChild)
	} else {
		newChild.setPrevSibling(before.getPrevSibling())
		before.setPrevSibling(newChild)
	}
	if newChild.getPrevSibling() != nil {
		newChild.getPrevSibling().setNextSibling(newChild)
	}
	if node.getFirstChild() == nil {
		node.setFirstChild(node.getLastChild())
	}
}

// detach a node from the tree it is in
func (node *BasicNode) detach() {
	if node.getNextSibling() != nil {
		node.getNextSibling().setPrevSibling(node.getPrevSibling())
		node.setNextSibling(nil)
	} else if node.getParent() != nil {
		node.getParent().setLastChild(node.getPrevSibling())
	}
	if node.getPrevSibling() != nil {
		node.getPrevSibling().setNextSibling(node.getNextSibling())
		node.setPrevSibling(nil)
	} else if node.getParent() != nil {
		node.getParent().setFirstChild(node.getNextSibling())
	}
	node.setParent(nil)
}

// Adds the specified childNode argument as the last child to the
// current node. If the argument referenced an existing node on the
// DOM tree, the node will be detached from its current position and
// attached at the new position.
//
// Returns a Node that is the appended child (aChild), except when aChild
// is a DocumentFragment, in which case the empty DocumentFragment
// is returned.
func (node *BasicNode) AppendChild(childNode Node) Node {
	chnode := childNode.(treeNode)
	chnode.detach()
	node.insertAfter(chnode, node.lastChild)
	if _, frag := childNode.(DocumentFragment); frag {
		return &BasicDocumentFragment{}
	}
	return childNode
}

// Returns the object's root
func (node *BasicNode) GetRootNode() Node {
	trc := treeNode(node)
	for trc.getParent() != nil {
		trc = trc.getParent()
	}
	return trc.(Node)
}

// Returns a boolean value indicating whether or not the element has
// any child nodes.
func (node *BasicNode) HasChildNodes() bool {
	return node.firstChild != nil
}

// Returns a boolean value indicating whether or not the two nodes are
// the same (that is, they reference the same object).
func (node *BasicNode) IsSameNode(n Node) bool {
	return node == n.(treeNode)
}

// Returns a Node representing the first direct child node of the
// node, or null if the node has no child.
func (node *BasicNode) GetFirstChild() Node {
	return node.firstChild.(Node)
}

// Returns a Node representing the last direct child node of the node,
// or null if the node has no child.
func (node *BasicNode) GetLastChild() Node {
	return node.lastChild.(Node)
}

// Returns a Node that is the parent of this node. If there is no such
// node, like if this node is the top of the tree or if doesn't
// participate in a tree, this property returns null.
func (node *BasicNode) GetParentNode() Node {
	return node.parent.(Node)
}

// Returns a Node representing the next node in the tree, or null if
// there isn't such node.
func (node *BasicNode) GetNextSibling() Node {
	return node.nextSibling.(Node)
}

// Returns a Node representing the previous node in the tree, or null
// if there isn't such node.
func (node *BasicNode) GetPreviousSibling() Node {
	return node.prevSibling.(Node)
}

// Inserts a Node before the reference node as a child of a
// specified parent node. Returns the added child (unless newNode is
// a DocumentFragment, in which case the empty DocumentFragment is
// returned).
func (node *BasicNode) InsertBefore(newNode, referenceNode Node) Node {
	newTreeNode := newNode.(treeNode)
	thisNode := treeNode(node)
	var after treeNode
	if referenceNode != nil {
		after = referenceNode.(treeNode)
	}
	newTreeNode.detach()
	thisNode.insertBefore(newTreeNode, after)
	if _, frag := newNode.(DocumentFragment); frag {
		return &BasicDocumentFragment{}
	}
	return newNode
}

// Returns true or false value indicating whether or not a node is a
// descendant of the calling node.
func (node *BasicNode) Contains(childNode Node) bool {
	// Is this node an ascendent?
	found := false
	tnode := childNode.(treeNode)
	for trc := childNode; trc != nil; trc = trc.GetParentNode() {
		if trc.(treeNode) == tnode {
			found = true
			break
		}
	}
	return found
}

// Returns a live NodeList containing all the children of this node
// (including elements, text and comments). NodeList being live means
// that if the children of the Node change, the NodeList object is
// automatically updated.
func (node *BasicNode) GetChildNodes() NodeList {
	return newBasicNodeList(node)
}

// Returns the Document that this node belongs to. If the node is
// itself a document, returns null.
func (node *BasicNode) GetOwnerDocument() Document {
	for trc := node.parent; trc != nil; trc = trc.getParent() {
		if doc, ok := trc.(Document); ok {
			return doc
		}
	}
	return nil
}

// Returns an Element that is the parent of this node. If the node has
// no parent, or if that parent is not an Element, this property
// returns null.
func (node *BasicNode) GetParentElement() Element {
	for trc := node.parent; trc != nil; trc = trc.getParent() {
		if doc, ok := trc.(Element); ok {
			return doc
		}
	}
	return nil
}

// Removes a child node from the current element, which must be a
// child of the current node.
func (node *BasicNode) RemoveChild(child Node) {
	tchild := child.(treeNode)
	if tchild.getParent() != node {
		return
	}
	tchild.detach()
}

// Replaces one child Node of the current one with the second one
// given in parameter.
func (node *BasicNode) ReplaceChild(newChild, oldChild Node) {
	oc := oldChild.(treeNode)
	nc := newChild.(treeNode)
	nc.detach()
	prev := oc.getPrevSibling()
	node.insertAfter(nc, prev)
}
