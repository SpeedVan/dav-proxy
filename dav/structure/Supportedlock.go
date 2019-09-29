package structure

import "encoding/xml"

// Supportedlock todo
type Supportedlock struct {
	XMLName   xml.Name `xml:"D:supportedlock"`
	Lockentry *Lockentry
}
