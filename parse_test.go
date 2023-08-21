package dom

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestNoNS1(t *testing.T) {
	input := `<note>
<to>Tove</to>
<from>Jani</from>
<heading>Reminder</heading>
<body>weekend!</body>
</note>`
	dec := xml.NewDecoder(strings.NewReader(input))
	doc, err := Parse(dec)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	root := doc.GetDocumentElement()
	if root.GetNodeName() != "note" {
		t.Errorf("Bad root name")
	}
	el := root.GetFirstChild().GetNextSibling()
	if el.GetNodeName() != "to" && el.GetFirstChild().(Text).GetValue() != "Tove" {
		t.Errorf("Bad to")
	}
	el = el.GetNextSibling().GetNextSibling()
	if el.GetNodeName() != "from" && el.GetFirstChild().(Text).GetValue() != "Jani" {
		t.Errorf("Bad from")
	}
	el = el.GetNextSibling().GetNextSibling()
	if el.GetNodeName() != "heading" && el.GetFirstChild().(Text).GetValue() != "Reminder" {
		t.Errorf("Bad heading")
	}
	el = el.GetNextSibling().GetNextSibling()
	if el.GetNodeName() != "body" && el.GetFirstChild().(Text).GetValue() != "weekend" {
		t.Errorf("Bad body")
	}
	if el.GetNextSibling().GetNextSibling() != nil {
		t.Errorf("Extra data")
	}
}

func TestCDATA(t *testing.T) {
	input := `<note>
<!CDATA[
characters
]]>
<to>Tove</to>
<from>Jani</from>
</note>`
	dec := xml.NewDecoder(strings.NewReader(input))
	doc, err := Parse(dec)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	root := doc.GetDocumentElement()
	if root.GetNodeName() != "note" {
		t.Errorf("Bad root name")
	}
	el := root.GetFirstChild().GetNextSibling()
	if el.GetNodeType() != CDATA_SECTION_NODE {
		t.Errorf("Not cdata: %s", el.GetNodeName())
	}
	if el.(CDATASection).GetValue() != "\ncharacters\n" {
		t.Errorf("Bad text: %s", el.(CDATASection).GetValue())
	}
	el = el.GetNextSibling().GetNextSibling()
	if el.GetNodeName() != "to" && el.GetFirstChild().(Text).GetValue() != "Tove" {
		t.Errorf("Bad to")
	}
	el = el.GetNextSibling().GetNextSibling()
	if el.GetNodeName() != "from" && el.GetFirstChild().(Text).GetValue() != "Jani" {
		t.Errorf("Bad from")
	}
	if el.GetNextSibling().GetNextSibling() != nil {
		t.Errorf("Extra data")
	}
}

func TestNS1(t *testing.T) {
	input := `

<root xmlns:h="http://www.w3.org/TR/html4/"
xmlns:f="https://test.com/furniture">

<h:table>
  <h:tr>
    <h:td>Apples</h:td>
    <h:td>Bananas</h:td>
  </h:tr>
</h:table>

<f:table>
  <f:name>African Coffee Table</f:name>
  <f:width>80</f:width>
  <f:length>120</f:length>
</f:table>

</root>`
	dec := xml.NewDecoder(strings.NewReader(input))
	doc, err := Parse(dec)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	t.Log(doc)
}
