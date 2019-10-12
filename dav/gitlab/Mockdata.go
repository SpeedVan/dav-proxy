package gitlab

import (
	"encoding/xml"

	st "github.com/SpeedVan/dav-proxy/dav/structure"
)

var (
	dataDav, _ = xml.Marshal(&st.Multistatus{
		D: "DAV:",
		Responses: []*st.Response{
			&st.Response{
				Href: &st.Href{
					Innerxml: "/dav",
				},
				Propstat: &st.Propstat{
					Prop: &st.Prop{
						Resourcetype: &st.Resourcetype{
							Collection: &st.Collection{
								D: "DAV:",
							},
						},
						Displayname: &st.Displayname{},
						Getlastmodified: &st.Getlastmodified{
							Innerxml: "Fri, 27 Sep 2019 11:42:40 GMT",
						},
						Supportedlock: &st.Supportedlock{
							Lockentry: &st.Lockentry{
								D: "DAV:",
								Lockscope: &st.Lockscope{
									Exclusive: &st.Exclusive{},
								},
								Locktype: &st.Locktype{
									Write: &st.Write{},
								},
							},
						},
					},
					Status: &st.Status{
						Innerxml: "HTTP/1.1 200 OK",
					},
				},
			},
			&st.Response{
				Href: &st.Href{
					Innerxml: "/dav/.bash_logout",
				},
				Propstat: &st.Propstat{
					Prop: &st.Prop{
						// Getcontentlength: &st.Getcontentlength{
						// 	Innerxml: "18",
						// },
						// Getcontenttype: &st.Getcontenttype{
						// 	Innerxml: "text/plain; charset=utf-8",
						// },
						Resourcetype: &st.Resourcetype{},
						Displayname: &st.Displayname{
							Innerxml: ".bash_logout",
						},
						Getlastmodified: &st.Getlastmodified{
							Innerxml: "Sun, 29 Dec 2013 10:26:31 GMT",
						},
						Getetag: &st.Getetag{
							Innerxml: "\"13442cef32eaa60012\"",
						},
						Supportedlock: &st.Supportedlock{
							Lockentry: &st.Lockentry{
								D: "DAV:",
								Lockscope: &st.Lockscope{
									Exclusive: &st.Exclusive{},
								},
								Locktype: &st.Locktype{
									Write: &st.Write{},
								},
							},
						},
					},
					Status: &st.Status{
						Innerxml: "HTTP/1.1 200 OK",
					},
				},
			},
		},
	})

	dataDavBashLogout, _ = xml.Marshal(&st.Multistatus{
		D: "DAV:",
		Responses: []*st.Response{
			&st.Response{
				Href: &st.Href{
					Innerxml: "/dav/.bash_logout",
				},
				Propstat: &st.Propstat{
					Prop: &st.Prop{
						Getcontentlength: &st.Getcontentlength{
							Innerxml: "18",
						},
						Getcontenttype: &st.Getcontenttype{
							Innerxml: "text/plain; charset=utf-8",
						},
						Resourcetype: &st.Resourcetype{},
						Displayname: &st.Displayname{
							Innerxml: ".bash_logout",
						},
						Getlastmodified: &st.Getlastmodified{
							Innerxml: "Sun, 29 Dec 2013 10:26:31 GMT",
						},
						Getetag: &st.Getetag{
							Innerxml: "\"13442cef32eaa60012\"",
						},
						Supportedlock: &st.Supportedlock{
							Lockentry: &st.Lockentry{
								D: "DAV:",
								Lockscope: &st.Lockscope{
									Exclusive: &st.Exclusive{},
								},
								Locktype: &st.Locktype{
									Write: &st.Write{},
								},
							},
						},
					},
					Status: &st.Status{
						Innerxml: "HTTP/1.1 200 OK",
					},
				},
			},
		},
	})
)
