package dom

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestNormalize(t *testing.T) {
	input := `<note><to></to></note>`
	dec := xml.NewDecoder(strings.NewReader(input))
	doc, err := Parse(dec)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	toEl := doc.GetDocumentElement().GetFirstChild().(Element)
	// Add three texts
	toEl.AppendChild(doc.CreateTextNode("1"))
	toEl.AppendChild(doc.CreateTextNode("2"))
	toEl.AppendChild(doc.CreateTextNode(""))
	// Add an element
	toEl.AppendChild(doc.CreateElement("el"))
	// Add three more texts
	toEl.AppendChild(doc.CreateTextNode("3"))
	toEl.AppendChild(doc.CreateTextNode("4"))
	toEl.AppendChild(doc.CreateTextNode("5"))

	doc.Normalize()

	toEl = doc.GetDocumentElement().GetFirstChild().(Element)
	t1 := toEl.GetFirstChild().(Text)
	if t1.GetValue() != "12" {
		t.Errorf("Wrong text: %v", t1)
	}
	t1 = t1.GetNextSibling().GetNextSibling().(Text)
	if t1.GetValue() != "345" {
		t.Errorf("Wrong text: %v", t1)
	}
	if t1.GetNextSibling() != nil {
		t.Errorf("Extra nodes")
	}
}
