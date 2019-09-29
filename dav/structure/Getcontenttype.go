package structure

import "encoding/xml"

// Getcontenttype todo
type Getcontenttype struct {
	XMLName  xml.Name `xml:"D:getcontenttype"`
	Innerxml string   `xml:",innerxml"`
}
