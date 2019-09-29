package structure

import "encoding/xml"

// Getlastmodified todo
type Getlastmodified struct {
	XMLName  xml.Name `xml:"D:getlastmodified"`
	Innerxml string   `xml:",innerxml"`
}
