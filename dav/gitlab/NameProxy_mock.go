package gitlab

import st "github.com/SpeedVan/dav-proxy/dav/structure"

// MockData just for mock
func MockData(path string) *st.Multistatus {
	return &st.Multistatus{
		D: "DAV:",
		Responses: []*st.Response{
			&st.Response{
				Href: &st.Href{
					Innerxml: path + "bash_logout",
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
							Innerxml: "bash_logout",
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
	}
}
