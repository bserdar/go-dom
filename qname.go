package dom

import (
	"strings"
)

type Name struct {
	// Namespace prefix
	prefix string
	// Namespace URI
	ns string
	// Local name
	local string
	// Qualified name
	qname string
}

func (name Name) Prefix() string { return name.prefix }
func (name Name) NS() string     { return name.ns }
func (name Name) Local() string  { return name.local }

func (name *Name) QName() string {
	if len(name.qname) != 0 {
		return name.qname
	}
	if len(name.ns) == 0 {
		if len(name.prefix) == 0 {
			name.qname = name.local
		} else {
			name.qname = name.prefix + ":" + name.local
		}
	} else {
		name.qname = name.ns + ":" + name.local
	}
	return name.qname
}

func (name *Name) SetNS(ns string) {
	name.ns = ns
	name.qname = ""
}

func (name *Name) IsEqualQName(n *Name) bool {
	return name.ns == n.ns && name.local == n.local
}

func NewName(ns, local string) Name {
	return Name{
		ns:    ns,
		local: local,
	}
}

// ParseName splits the input at ':'
func ParseName(in string) Name {
	ix := strings.IndexRune(in, ':')
	if ix == -1 {
		return Name{
			local: in,
		}
	}
	return Name{
		prefix: in[:ix],
		local:  in[ix+1:],
	}
}
