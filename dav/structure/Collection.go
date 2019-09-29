package structure

import "encoding/xml"

// Collection todo
type Collection struct {
	XMLName xml.Name `xml:"D:collection"`
	D       string   `xml:"xmlns:D,attr"`
}
