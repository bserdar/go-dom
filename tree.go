package dom

type tnode struct {
	parent Node
	next   Node
	prev   Node
	child  Node
	// ver is used by nodelists to keep track of child list changes
	ver int
}

func (node *tnode) firstChild() Node {
	return node.child
}

func (node *tnode) lastChild() Node {
	if node.child == nil {
		return nil
	}
	return node.child.treeNode().prev
}

func (node *tnode) prevSibling() Node {
	if node.parent == nil {
		return nil
	}
	if node.parent.treeNode().firstChild().treeNode() == node {
		return nil
	}
	return node.prev
}

func (node *tnode) nextSibling() Node {
	if node.parent == nil {
		return nil
	}
	if node.parent.treeNode().lastChild().treeNode() == node {
		return nil
	}
	return node.next
}

// Insert child after given node. If after is nil, insert as last node
func insertChildAfter(parent, newChild, after Node) {
	newChildtn := newChild.treeNode()
	newChildtn.parent = parent
	parenttn := parent.treeNode()
	parenttn.ver++
	if after == nil {
		after = parenttn.lastChild()
		if after == nil {
			newChildtn.next = newChild
			newChildtn.prev = newChild
			parenttn.child = newChild
			return
		}
	}
	newChildtn.prev = after
	newChildtn.next = after.treeNode().next
	after.treeNode().next.treeNode().prev = newChild
	after.treeNode().next = newChild
}

// Insert child before given node. If before is nil, insert as first node
func insertChildBefore(parent, newChild, before Node) {
	newChildtn := newChild.treeNode()
	newChildtn.parent = parent
	parenttn := parent.treeNode()
	parenttn.ver++
	if before == nil {
		before = parenttn.firstChild()
		if before == nil {
			newChildtn.next = newChild
			newChildtn.prev = newChild
			parenttn.child = newChild
			return
		}
	}
	newChildtn.next = before
	newChildtn.prev = before.treeNode().prev
	before.treeNode().prev.treeNode().next = newChild
	before.treeNode().prev = newChild
}

func detachChild(parent, child Node) {
	childtn := child.treeNode()
	if parent != nil {
		parenttn := parent.treeNode()
		parenttn.ver++
		if parenttn.child == child {
			parenttn.child = childtn.next
		}
	}
	if childtn.next != nil {
		childtn.next.treeNode().prev = childtn.prev
	}
	if childtn.prev != nil {
		childtn.prev.treeNode().next = childtn.next
	}
	childtn.next = nil
	childtn.prev = nil
}
