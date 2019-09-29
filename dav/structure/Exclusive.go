package structure

import "encoding/xml"

// Exclusive todo
type Exclusive struct {
	XMLName xml.Name `xml:"D:exclusive"`
}
