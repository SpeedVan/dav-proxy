package structure

import "encoding/xml"

// Propstat todo
type Propstat struct {
	XMLName xml.Name `xml:"D:propstat"`
	Prop    *Prop
	Status  *Status
}
