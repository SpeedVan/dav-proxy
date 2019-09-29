package structure

import "encoding/xml"

// Status todo
type Status struct {
	XMLName  xml.Name `xml:"D:status"`
	Innerxml string   `xml:",innerxml"`
}
