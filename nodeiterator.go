package dom

type WhatToShow uint32

const SHOW_ALL WhatToShow = 4294967295
const SHOW_ATTRIBUTE WhatToShow = 2
const SHOW_CDATA_SECTION WhatToShow = 8
const SHOW_COMMENT WhatToShow = 128
const SHOW_DOCUMENT WhatToShow = 256
const SHOW_DOCUMENT_FRAGMENT WhatToShow = 1024
const SHOW_DOCUMENT_TYPE WhatToShow = 512
const SHOW_ELEMENT WhatToShow = 1
const SHOW_ENTITY WhatToShow = 32
const SHOW_ENTITY_REFERENCE WhatToShow = 16
const SHOW_NOTATION WhatToShow = 2048
const SHOW_PROCESSING_INSTRUCTION WhatToShow = 64
const SHOW_TEXT WhatToShow = 4

type NodeIterator interface {
	PreviousNode() Node
	NextNode() Node
}
