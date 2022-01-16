package dom

type DocumentFragment interface {
	Node

	Append(...Node)
	Prepend(...Node)
}
