package dom

type BasicNodeList struct {
	parentNode treeNode

	list    []treeNode
	listVer int
}

func newBasicNodeList(parent treeNode) *BasicNodeList {
	ret := &BasicNodeList{
		parentNode: parent,
		list:       make([]treeNode, 0),
		listVer:    parent.getChildListVer(),
	}
	ret.buildList()
	return ret
}

func (list *BasicNodeList) buildList() {
	list.list = make([]treeNode, 0)
	for itr := list.parentNode.getFirstChild(); itr != nil; itr = itr.getNextSibling() {
		list.list = append(list.list, itr)
	}
	list.listVer = list.parentNode.getChildListVer()
}

func (list *BasicNodeList) GetLength() int {
	if list.listVer != list.parentNode.getChildListVer() {
		list.buildList()
	}
	return len(list.list)
}

func (list *BasicNodeList) Item(i int) Node {
	if list.listVer != list.parentNode.getChildListVer() {
		list.buildList()
	}
	return list.list[i].(Node)
}
