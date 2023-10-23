package dom

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestAttr(t *testing.T) {
	input := `<root><el a1="val"/></root>`
	dec := xml.NewDecoder(strings.NewReader(input))
	doc, err := Parse(dec)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	el := doc.GetDocumentElement().GetFirstChild().(Element)

	s, ok := el.GetAttribute("a1")
	if s != "val" || !ok {
		t.Errorf("%s %v", s, ok)
	}

	if _, ok := el.GetAttribute("a2"); ok {
		t.Errorf("Attribute found")
	}

	names := el.GetAttributeNames()
	if len(names) != 1 || names[0] != "a1" {
		t.Errorf("names %v", names)
	}

	node := el.GetAttributeNode("a1")
	if node.GetLocalName() != "a1" {
		t.Errorf("Local name %s", node.GetLocalName())
	}
	if node.GetName() != "a1" {
		t.Errorf("name %s", node.GetName())
	}
	n := Name{Name: xml.Name{
		Space: "",
		Local: "a1",
	},
		Prefix: "",
	}
	if node.GetQName() != n {
		t.Errorf("name %s", node.GetName())
	}
	if node.GetOwnerElement() != el {
		t.Errorf("Wrong owner")
	}
	if node.GetValue() != "val" {
		t.Errorf("Wrong value")
	}
}

func TestAttrNS(t *testing.T) {
	input := `<root xmlns:ns1="http://example.org"><el ns1:a1="val" /></root>`
	dec := xml.NewDecoder(strings.NewReader(input))
	doc, err := Parse(dec)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	el := doc.GetDocumentElement().GetFirstChild().(Element)
	if el.GetQName().Prefix != "" || el.GetQName().Space != "" {
		t.Errorf("Wrong element ns")
	}

	s, ok := el.GetAttribute("a1")
	if s != "" || ok {
		t.Errorf("%s %v", s, ok)
	}

	if s, ok := el.GetAttributeNS("http://example.org", "a1"); !ok || s != "val" {
		t.Errorf("AttributeNS  not found")
	}

	names := el.GetAttributeNames()
	if len(names) != 1 || names[0] != "ns1:a1" {
		t.Errorf("names %v", names)
	}

	node := el.GetAttributeNodeNS("http://example.org", "a1")
	if node.GetLocalName() != "a1" {
		t.Errorf("Local name %s", node.GetLocalName())
	}
	if node.GetName() != "ns1:a1" {
		t.Errorf("name %s", node.GetName())
	}
	n := Name{Name: xml.Name{
		Space: "http://example.org",
		Local: "a1",
	},
		Prefix: "ns1",
	}
	if node.GetQName() != n {
		t.Errorf("name %s", node.GetName())
	}
	if node.GetOwnerElement() != el {
		t.Errorf("Wrong owner")
	}
	if node.GetValue() != "val" {
		t.Errorf("Wrong value")
	}
}

func TestAttrNS2(t *testing.T) {
	// XML with default namespaces, el in namespace, attr not
	input := `<root xmlns="http://example.org"><el a1="val" /></root>`
	dec := xml.NewDecoder(strings.NewReader(input))
	doc, err := Parse(dec)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	el := doc.GetDocumentElement().GetFirstChild().(Element)
	if el.GetQName().Prefix != "" || el.GetQName().Space != "http://example.org" {
		t.Errorf("Wrong element ns")
	}

	s, ok := el.GetAttribute("a1")
	if s != "val" || !ok {
		t.Errorf("%s %v", s, ok)
	}

	if s, ok := el.GetAttributeNS("", "a1"); !ok || s != "val" {
		t.Errorf("AttributeNS  not found")
	}

	names := el.GetAttributeNames()
	if len(names) != 1 || names[0] != "a1" {
		t.Errorf("names %v", names)
	}

	node := el.GetAttributeNode("a1")
	if node.GetLocalName() != "a1" {
		t.Errorf("Local name %s", node.GetLocalName())
	}
	if node.GetName() != "a1" {
		t.Errorf("name %s", node.GetName())
	}
	n := Name{Name: xml.Name{
		Space: "",
		Local: "a1",
	},
		Prefix: "",
	}
	if node.GetQName() != n {
		t.Errorf("name %s", node.GetName())
	}
	if node.GetOwnerElement() != el {
		t.Errorf("Wrong owner")
	}
	if node.GetValue() != "val" {
		t.Errorf("Wrong value")
	}
}

func TestEqual(t *testing.T) {
	input := `<root><el a1="val" />   <!--c--> </root>`

	dec := xml.NewDecoder(strings.NewReader(input))
	doc, err := Parse(dec)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	dec = xml.NewDecoder(strings.NewReader(input))
	doc2, err := Parse(dec)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if !doc.IsEqualNode(doc2) {
		t.Errorf("Not equal")
	}
}

func TestAttrMod(t *testing.T) {
	input := `<root><el a1="val"/></root>`
	dec := xml.NewDecoder(strings.NewReader(input))
	doc, err := Parse(dec)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	el := doc.GetDocumentElement().GetFirstChild().(Element)
	if s, ok := el.GetAttribute("a1"); s != "val" || !ok {
		t.Errorf("a1 not here")
	}

	el.SetAttribute("a1", "updated")
	if s, ok := el.GetAttribute("a1"); s != "updated" || !ok {
		t.Errorf("a1 wrong")
	}

	el.SetAttribute("a2", "new")
	if s, ok := el.GetAttribute("a1"); s != "updated" || !ok {
		t.Errorf("a1 wrong")
	}
	if s, ok := el.GetAttribute("a2"); s != "new" || !ok {
		t.Errorf("a2 wrong")
	}

	el.RemoveAttribute("a1")
	if s, ok := el.GetAttribute("a1"); s != "" || ok {
		t.Errorf("a1 still here")
	}
	if s, ok := el.GetAttribute("a2"); s != "new" || !ok {
		t.Errorf("a2 wrong")
	}
}
