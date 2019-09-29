package structure

import "encoding/xml"

// Displayname todo
type Displayname struct {
	XMLName  xml.Name `xml:"D:displayname"`
	Innerxml string   `xml:",innerxml"`
}
