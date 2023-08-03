package dom

type BasicNodeList struct {
	parentNode *treeNode

	list []Node
	ver  int
}

func newBasicNodeList(parent Node) *BasicNodeList {
	ret := &BasicNodeList{
		parentNode: parent.treeNode(),
		list:       nil,
		ver:        0,
	}
	return ret
}

func (list *BasicNodeList) buildList() {
	if list.list != nil && list.ver != list.parentNode.ver {
		list.list = nil
	}
	if list.list != nil && list.ver == list.parentNode.ver {
		return
	}
	list.list = make([]Node, 0)
	for itr := list.parentNode.GetFirstChild(); itr != nil; itr = itr.GetNextSibling() {
		list.list = append(list.list, itr)
	}
	list.listVer = list.parentNode.ver
}

func (list *BasicNodeList) GetLength() int {
	list.buildList()
	return len(list.list)
}

func (list *BasicNodeList) Item(i int) Node {
	list.buildList()
	if i < 0 || i >= len(list.list) {
		return nil
	}
	return list.list[i]
}
