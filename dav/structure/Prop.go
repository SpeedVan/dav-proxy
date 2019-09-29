package structure

import "encoding/xml"

// Prop todo
type Prop struct {
	XMLName          xml.Name `xml:"D:prop"`
	Getcontentlength *Getcontentlength
	Getcontenttype   *Getcontenttype
	Resourcetype     *Resourcetype
	Displayname      *Displayname
	Getlastmodified  *Getlastmodified
	Getetag          *Getetag
	Supportedlock    *Supportedlock
}
