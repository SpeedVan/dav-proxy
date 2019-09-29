package structure

import "encoding/xml"

// Response todo
type Response struct {
	XMLName  xml.Name `xml:"D:response"`
	Href     *Href
	Propstat *Propstat
}
