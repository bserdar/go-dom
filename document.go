package dom

type Document interface {
	Node

	// Returns the Element that is a direct child of the document.
	GetDocumentElement() Element

	// Creates a new Attr object and returns it.
	CreateAttribute(string) Attr

	// Creates a new attribute node in a given namespace and returns it.
	CreateAttributeNS(ns string, name string) Attr

	// Creates a new CDATA node and returns it.
	CreateCDATASection(string) CDATASection

	// Creates a new comment node and returns it.
	CreateComment(string) Comment

	// // Creates a new document fragment.
	// CreateDocumentFragment() DocumentFragment

	// Creates a new element with the given tag name.
	CreateElement(string) Element

	// Creates a new element with the given tag name and namespace URI.
	CreateElementNS(ns string, tag string) Element

	// Creates a text node.
	CreateTextNode(string) Text

	//Creates a new ProcessingInstruction object.
	CreateProcessingInstruction(target, data string) ProcessingInstruction

	// // Replaces entities, normalizes text nodes, etc.
	// NormalizeDocument()

	//	Adopt node from an external document.
	AdoptNode(Node) Node
}
