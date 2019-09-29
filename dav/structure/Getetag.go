package structure

import "encoding/xml"

// Getetag todo
type Getetag struct {
	XMLName  xml.Name `xml:"D:getetag"`
	Innerxml string   `xml:",innerxml"`
}
