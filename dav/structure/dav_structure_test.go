package structure

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	bytes, err := xml.Marshal(&Multistatus{
		D: "DAV:",
		Responses: []*Response{
			&Response{
				Href: &Href{
					Innerxml: "/dav",
				},
				Propstat: &Propstat{
					Prop: &Prop{
						Resourcetype: &Resourcetype{
							Collection: &Collection{
								D: "DAV:",
							},
						},
						Displayname: &Displayname{},
						Getlastmodified: &Getlastmodified{
							Innerxml: "Fri, 27 Sep 2019 11:42:40 GMT",
						},
						Supportedlock: &Supportedlock{
							Lockentry: &Lockentry{
								D: "DAV:",
								Lockscope: &Lockscope{
									Exclusive: &Exclusive{},
								},
								Locktype: &Locktype{
									Write: &Write{},
								},
							},
						},
					},
					Status: &Status{
						Innerxml: "HTTP/1.1 200 OK",
					},
				},
			},
			&Response{
				Href: &Href{
					Innerxml: "/dav/.bash_logout",
				},
				Propstat: &Propstat{
					Prop: &Prop{
						Getcontentlength: &Getcontentlength{
							Innerxml: "18",
						},
						Getcontenttype: &Getcontenttype{
							Innerxml: "text/plain; charset=utf-8",
						},
						Resourcetype: &Resourcetype{},
						Displayname: &Displayname{
							Innerxml: ".bash_logout",
						},
						Getlastmodified: &Getlastmodified{
							Innerxml: "Sun, 29 Dec 2013 10:26:31 GMT",
						},
						Getetag: &Getetag{
							Innerxml: "\"13442cef32eaa60012\"",
						},
						Supportedlock: &Supportedlock{
							Lockentry: &Lockentry{
								D: "DAV:",
								Lockscope: &Lockscope{
									Exclusive: &Exclusive{},
								},
								Locktype: &Locktype{
									Write: &Write{},
								},
							},
						},
					},
					Status: &Status{
						Innerxml: "HTTP/1.1 200 OK",
					},
				},
			},
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))

}
