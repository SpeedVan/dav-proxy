package structure

import "encoding/xml"

// Lockentry todo
type Lockentry struct {
	XMLName   xml.Name `xml:"D:lockentry"`
	D         string   `xml:"xmlns:D,attr"`
	Lockscope *Lockscope
	Locktype  *Locktype
}
