package dom

import (
	"encoding/xml"
	"strings"
)

type Name struct {
	xml.Name
	// Namespace prefix
	Prefix string
	qname  string
}

func (name *Name) QName() string {
	if len(name.qname) == 0 {
		if len(name.Prefix) == 0 {
			name.qname = name.Local
		} else {
			name.qname = name.Prefix + ":" + name.Local
		}
	}
	return name.qname
}

// ParseName splits the input at ':'
func ParseName(in string) xml.Name {
	ix := strings.IndexRune(in, ':')
	if ix == -1 {
		return xml.Name{
			Local: in,
		}
	}
	return xml.Name{
		Space: in[:ix],
		Local: in[ix+1:],
	}
}
