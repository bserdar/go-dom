package dom

import (
	"bufio"
	"strings"
	"testing"
)

func TestDTDScanner(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(`!DOCTYPE note
[
<!ELEMENT note (to,from,heading,body)>
<!ELEMENT to (#PCDATA)>
]`))
	scanner.Split(scanDTDToken)
	expected := []string{"!DOCTYPE", "note", "[", "<", "!ELEMENT", "note", "(", "to", ",", "from", ",", "heading", ",", "body", ")", ">", "<", "!ELEMENT", "to", "(", "#PCDATA", ")", ">", "]"}

	i := 0
	for scanner.Scan() {
		str := scanner.Text()
		t.Log(str)
		if str != expected[i] {
			t.Errorf("Expected %s got %s", expected[i], str)
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		t.Error(err)
	}
	if i < len(expected) {
		t.Errorf("i: %d", i)
	}
}

func TestDTDParse(t *testing.T) {
	ret, ok, err := parseDocumentType([]byte(`!DOCTYPE note
[
<!ELEMENT note (to,from,heading,body)>
<!ELEMENT to (#PCDATA)>
]`))
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Errorf("Not DTD")
	}
	dt := ret.(*BasicDocumentType)
	if dt.GetName() != "note" {
		t.Errorf("Wrong name")
	}
	if dt.defn != "<!ELEMENT note(to,from,heading,body)><!ELEMENT to(#PCDATA)>" {
		t.Errorf("Wrong defn: %s", dt.defn)
	}
}
