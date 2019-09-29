package structure

import "encoding/xml"

// Locktype todo
type Locktype struct {
	XMLName xml.Name `xml:"D:locktype"`
	Write   *Write
}
