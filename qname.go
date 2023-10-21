package dom

import (
	"encoding/xml"
	"strings"
)

type Name struct {
	xml.Name
	// Namespace prefix
	Prefix string
}

func (name *Name) QName() string {
	if len(name.Prefix) == 0 {
		return name.Local
	}
	return name.Prefix + ":" + name.Local
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
