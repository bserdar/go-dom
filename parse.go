package dom

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"
	"unicode"
)

const (
	xmlURL      = "http://www.w3.org/XML/1998/namespace"
	xmlnsPrefix = "xmlns"
	xmlPrefix   = "xml"
)

func Parse(decoder *xml.Decoder) (Document, error) {
	ret := &BasicDocument{}
	ret.ownerDocument = ret

	type stackEl interface {
		Node
		getDefaultNamespace() string
	}
	type stackItem struct {
		el   stackEl
		dict *nsMap
	}
	elementStack := make([]stackItem, 0, 32)
	elementStack = append(elementStack, stackItem{
		el:   ret,
		dict: newNsMap(),
	})
	elementStack[0].dict.add(xmlPrefix, xmlURL)

	makeName := func(tokenName xml.Name, dict *nsMap, defaultNamespace string) Name {
		name := Name{}
		if len(tokenName.Space) == 0 {
			name.Space = defaultNamespace
			name.Local = tokenName.Local
		} else {
			name.Space, _ = dict.getNS(tokenName.Space)
			// Preserve Go XML decoder semantics: Unknown namespace uses the prefix as namespace
			if len(name.Space) == 0 {
				name.Space = tokenName.Space
			}
			name.Prefix = tokenName.Space
			name.Local = tokenName.Local
		}
		return name
	}

	autoClose := func(name xml.Name) bool {
		if decoder.Strict {
			return false
		}
		for _, str := range decoder.AutoClose {
			if strings.EqualFold(str, name.Local) {
				return true
			}
		}
		return false
	}

	for {
		tok, err := decoder.RawToken()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		switch token := tok.(type) {
		case xml.StartElement:

			defaultNamespace := ""
			dict := newNsMap()
			if len(elementStack) > 0 {
				defaultNamespace = elementStack[len(elementStack)-1].el.getDefaultNamespace()
				dict.parent = elementStack[len(elementStack)-1].dict
			} else {
				dict.add(xmlPrefix, xmlURL)
			}
			for _, attr := range token.Attr {
				switch {
				case attr.Name.Space == "" && attr.Name.Local == xmlnsPrefix:
					// Set the default namespace for this element
					defaultNamespace = attr.Value
				case attr.Name.Space == xmlnsPrefix:
					// Set the namespace uri for a prefix
					if err := dict.add(attr.Name.Local, attr.Value); err != nil {
						return nil, err
					}
				case attr.Name.Space == xmlPrefix:
					// Namespace is xmlURL
				}
			}
			// We processed the namespaces, now deal with the element itself
			name := makeName(token.Name, dict, defaultNamespace)
			var newElement *BasicElement
			if len(name.Space) > 0 {
				newElement = ret.CreateElementNS(name.Space, name.Local).(*BasicElement)
			} else {
				newElement = ret.CreateElement(name.Local).(*BasicElement)
			}
			newElement.defaultNamespace = defaultNamespace
			newElement.name = name
			for _, attr := range token.Attr {
				name := makeName(attr.Name, dict, defaultNamespace)
				var newAttr *BasicAttr
				if len(name.Space) > 0 {
					newAttr = ret.CreateAttributeNS(name.Space, name.Local).(*BasicAttr)
				} else {
					newAttr = ret.CreateAttribute(name.Local).(*BasicAttr)
				}
				newAttr.name = name
				newAttr.value = attr.Value
				newElement.attributes.SetNamedItem(newAttr)
			}
			elementStack[len(elementStack)-1].el.AppendChild(newElement)
			if !autoClose(name.Name) {
				elementStack = append(elementStack, stackItem{el: newElement, dict: dict})
			}

		case xml.EndElement:
			if len(elementStack) == 0 {
				return nil, &xml.SyntaxError{
					Msg: "Extra objects before document",
				}
			}
			if autoClose(token.Name) {
				break
			}
			last := elementStack[len(elementStack)-1]
			if el, ok := last.el.(*BasicElement); ok {
				name := makeName(token.Name, last.dict, el.defaultNamespace)
				if name.Space == el.name.Space && strings.EqualFold(name.Local, el.name.Local) {
					// ok
				} else {
					return nil, &xml.SyntaxError{
						Msg: fmt.Sprintf("Mismatched closing tag %s", name.Local),
					}
				}
			}
			elementStack = elementStack[:len(elementStack)-1]

		case xml.CharData:
			if len(elementStack) == 1 {
				// charData must be only spaces
				for _, x := range string(token) {
					if !unicode.IsSpace(x) {
						return nil, &xml.SyntaxError{
							Msg: "Extra characters before document",
						}
					}
				}
			} else {
				newNode := ret.CreateTextNode(string(token))
				elementStack[len(elementStack)-1].el.AppendChild(newNode)
			}

		case xml.Comment:
			if len(elementStack) == 1 {
				return nil, &xml.SyntaxError{
					Msg: "Comment before document",
				}
			}
		case xml.ProcInst:
			if len(elementStack) == 1 {
				return nil, &xml.SyntaxError{
					Msg: "Processing instruction before document",
				}
			}
		case xml.Directive:
			if len(elementStack) == 1 {
				return nil, &xml.SyntaxError{
					Msg: "XML directive before document",
				}
			}
			var newNode Node
			content := string(token)
			if strings.HasPrefix(content, "CDATA[") && strings.HasSuffix(content, "]]") {
				newNode = ret.CreateCDATASection(string(content[6 : len(content)-2]))
			}
			elementStack[len(elementStack)-1].el.AppendChild(newNode)
		}
	}
	return ret, nil
}
