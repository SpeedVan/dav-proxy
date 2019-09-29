package structure

import "encoding/xml"

// Lockscope todo
type Lockscope struct {
	XMLName   xml.Name `xml:"D:lockscope"`
	Exclusive *Exclusive
}
