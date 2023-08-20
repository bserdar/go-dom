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
<body>Don't forget me this weekend!</body>
</note>`
	dec := xml.NewDecoder(strings.NewReader(input))
	doc, err := Parse(dec)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	t.Log(doc)
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
	t.Log(doc)
}

func TestNS1(t *testing.T) {
	input := `

<root xmlns:h="http://www.w3.org/TR/html4/"
xmlns:f="https://www.w3schools.com/furniture">

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
