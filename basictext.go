package dom

type basicChardata struct {
	basicNode
	text string
}

func (cd *basicChardata) GetValue() string     { return cd.text }
func (cd *basicChardata) SetValue(text string) { cd.text = text }

func (cd *basicChardata) AppendChild(Node) Node {
	panic(ErrHierarchyRequest("AppendChild", "Invalid node type: character data node"))
}

func (cd *basicChardata) HasChildNodes() bool { return false }

func (cd *basicChardata) InsertBefore(newNode, referenceNode Node) Node {
	panic(ErrHierarchyRequest("InsertBefore", "Invalid node type: character data node"))
}

func (cs *basicChardata) RemoveChild(Node) {
	panic(ErrHierarchyRequest("RemoveChild", "Invalid node type: character data node"))
}

func (cs *basicChardata) Normalize() {}

type BasicCDataSection struct {
	basicChardata
}

var _ CDATASection = &BasicCDataSection{}

// Returns a boolean value indicating whether or not the two nodes are
// the same (that is, they reference the same object).
func (cd *BasicCDataSection) IsSameNode(node Node) bool { return node == cd }

// Returns "#cdata-section"
func (cd *BasicCDataSection) GetNodeName() string { return "#cdata-section" }

// Returns CDATA_SECTION_NODE
func (cd *BasicCDataSection) GetNodeType() NodeType { return CDATA_SECTION_NODE }

func (cd *BasicCDataSection) IsEqualNode(node Node) bool {
	n, ok := node.(*BasicCDataSection)
	if !ok {
		return false
	}
	return n.text == cd.text
}

type BasicText struct {
	basicChardata
}

var _ Text = &BasicText{}

// Returns "#text"
func (cd *BasicText) GetNodeName() string { return "#text" }

// Returns TEXT_NODE
func (cd *BasicText) GetNodeType() NodeType { return TEXT_NODE }

func (cd *BasicText) IsEqualNode(node Node) bool {
	n, ok := node.(*BasicText)
	if !ok {
		return false
	}
	return n.text == cd.text
}

// Returns a boolean value indicating whether or not the two nodes are
// the same (that is, they reference the same object).
func (cd *BasicText) IsSameNode(node Node) bool { return node == cd }
