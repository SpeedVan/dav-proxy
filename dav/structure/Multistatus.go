package structure

import "encoding/xml"

// Multistatus todo
type Multistatus struct {
	XMLName   xml.Name `xml:"D:multistatus"`
	D         string   `xml:"xmlns:D,attr"`
	Responses []*Response
}
