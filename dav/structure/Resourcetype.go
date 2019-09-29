package structure

import "encoding/xml"

// Resourcetype todo
type Resourcetype struct {
	XMLName    xml.Name `xml:"D:resourcetype"`
	Collection *Collection
}
