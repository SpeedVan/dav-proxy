package structure

import "encoding/xml"

// Href todo
type Href struct {
	XMLName  xml.Name `xml:"D:href"`
	Innerxml string   `xml:",innerxml"`
}
