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

func TestClone(t *testing.T) {
	input := `<note>
<to attr="val">  </to>  <!--comment-->
</note>`
	dec := xml.NewDecoder(strings.NewReader(input))
	doc, err := Parse(dec)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	newDoc := doc.CloneNode(true)
	if !doc.IsEqualNode(newDoc) {
		t.Errorf("Not equal")
	}
}

func TestProcessingInstruction(t *testing.T) {
	input := `<?xml version = "1.0" ?>
<note>
<to attr="val">  </to>  <!--comment-->
</note>`
	dec := xml.NewDecoder(strings.NewReader(input))
	doc, err := Parse(dec)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	pi := doc.GetFirstChild().(ProcessingInstruction)
	if pi.GetTarget() != "xml" {
		t.Errorf("Wrong target")
	}
	if pi.GetValue() != `version = "1.0" ` {
		t.Errorf("Wrong text: %s", pi.GetValue())
	}
}

func TestAdopt(t *testing.T) {
	input := `<note>
<to attr="val">  </to>  <!--comment-->
</note>`
	dec := xml.NewDecoder(strings.NewReader(input))
	doc, err := Parse(dec)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	newDoc := NewDocument()
	root := newDoc.CreateElement("newRoot")
	newDoc.AppendChild(root)

	toNode := doc.GetDocumentElement().GetFirstElementChild()
	newToNode := newDoc.AdoptNode(toNode)
	if newToNode != toNode {
		t.Errorf("Wrong return value")
	}
	root.AppendChild(newToNode)

	if doc.GetDocumentElement().GetFirstElementChild() != nil {
		t.Errorf("still in original doc")
	}

	if newDoc.GetDocumentElement().GetFirstChild() != newToNode {
		t.Errorf("Not in new doc")
	}
}

func TestDTD(t *testing.T) {
	input := `<?xml version="1.0"?>
<!DOCTYPE note
[
<!ELEMENT note (to,from,heading,body)>
<!ELEMENT to (#PCDATA)>
<!ELEMENT from (#PCDATA)>
<!ELEMENT heading (#PCDATA)>
<!ELEMENT body (#PCDATA)>
]>

<note>
<to attr="val">  </to>  <!--comment-->
</note>`
	dec := xml.NewDecoder(strings.NewReader(input))
	doc, err := Parse(dec)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	_ = doc
}
