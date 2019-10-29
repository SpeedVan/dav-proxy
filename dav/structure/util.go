package structure

func ToDir(path, name, now string) *Response {
	return &Response{
		Href: &Href{
			Innerxml: path + name + "/",
		},
		Propstat: &Propstat{
			Prop: &Prop{
				Resourcetype: &Resourcetype{
					Collection: &Collection{
						D: "DAV:",
					},
				},
				Displayname: &Displayname{
					Innerxml: name,
				},
				Getlastmodified: &Getlastmodified{
					Innerxml: now,
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
	}
}
