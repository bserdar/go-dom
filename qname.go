package dom

import (
	"strings"
)

type QName struct {
	Prefix string
	Local  string
}

// String returns local name, or ns:local
func (name QName) String() string {
	if len(name.Prefix) == 0 {
		return name.Local
	}
	return name.Prefix + ":" + name.Local
}

// ParseQName splits the input at ':'
func ParseQName(in string) QName {
	ix := strings.IndexRune(in, ':')
	if ix == -1 {
		return QName{Local: in}
	}
	return QName{
		Prefix: in[:ix],
		Local:  in[:ix+1],
	}
}
