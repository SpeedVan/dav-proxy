package structure

import "encoding/xml"

// Getcontentlength todo
type Getcontentlength struct {
	XMLName  xml.Name `xml:"D:getcontentlength"`
	Innerxml string   `xml:",innerxml"`
}
