package gitlab

import (
	"github.com/SpeedVan/dav-proxy/dav/structure"
	st "github.com/SpeedVan/dav-proxy/dav/structure"
	"github.com/SpeedVan/go-common/client/httpclient/gitlab"
	"github.com/SpeedVan/go-common/client/httpclient/gitlab/graphql"
)

// treeNodes2DAVStructure
func treeNodes2DAVStructure2(nodes []*gitlab.TreeNode, url string, now string, fileInfoFunc func(string) string) *structure.Multistatus {
	return &st.Multistatus{
		D:         "DAV:",
		Responses: treeNodes2DAVResponses2(nodes, url, now, fileInfoFunc),
	}
}

// treeNodes2DAVResponses mean map
func treeNodes2DAVResponses2(nodes []*gitlab.TreeNode, url string, now string, fileInfoFunc func(string) string) []*structure.Response {
	nodeLen := len(nodes)
	responses := make([]*structure.Response, nodeLen)

	treeLast := 0
	brChan := make(chan *structure.Response)
	for i := 0; i < nodeLen; i++ {
		node := nodes[i]
		if node.Type == "tree" {
			responses[treeLast] = treeNode2DAVResponse2(node, url, now)
			treeLast++
		} else {
			go blobNode2DAVResponse2WithChan(node, url, now, fileInfoFunc, brChan)
		}
	}
	defer close(brChan)
	for i := treeLast; i < nodeLen; i++ {
		response := <-brChan
		responses[i] = response
	}
	return responses
}

// treeNode2DAVResponse2
func treeNode2DAVResponse2(node *gitlab.TreeNode, url string, now string) *structure.Response {
	return &st.Response{
		Href: &st.Href{
			Innerxml: url + node.Name,
		},
		Propstat: &st.Propstat{
			Prop: &st.Prop{
				Resourcetype: &st.Resourcetype{
					Collection: &st.Collection{
						D: "DAV:",
					},
				},
				Displayname: &st.Displayname{
					Innerxml: node.Name,
				},
				Getlastmodified: &st.Getlastmodified{
					Innerxml: now,
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
	}
}

func blobNode2DAVResponse2(node *gitlab.TreeNode, url string, now string, fileInfoFunc func(string) string) *structure.Response {
	var getcontentlength *st.Getcontentlength
	if size := fileInfoFunc(node.ID); size == "" {
		getcontentlength = nil
	} else {
		getcontentlength = &st.Getcontentlength{
			Innerxml: size,
		}
	}
	return &st.Response{
		Href: &st.Href{
			Innerxml: url + node.Name,
		},
		Propstat: &st.Propstat{
			Prop: &st.Prop{
				Getcontentlength: getcontentlength,
				// Getcontenttype: &st.Getcontenttype{
				// 	Innerxml: "text/plain; charset=utf-8",
				// },
				Resourcetype: &st.Resourcetype{},
				Displayname: &st.Displayname{
					Innerxml: node.Name,
				},
				Getlastmodified: &st.Getlastmodified{
					Innerxml: now,
				},
				Getetag: &st.Getetag{
					Innerxml: "\"" + node.ID + "\"",
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
	}
}

func blobNode2DAVResponse2WithChan(node *gitlab.TreeNode, url string, now string, fileInfoFunc func(string) string, brChan chan<- *structure.Response) {
	var getcontentlength *st.Getcontentlength
	if size := fileInfoFunc(node.ID); size == "" {
		getcontentlength = nil
	} else {
		getcontentlength = &st.Getcontentlength{
			Innerxml: size,
		}
	}
	brChan <- &st.Response{
		Href: &st.Href{
			Innerxml: url + node.Name,
		},
		Propstat: &st.Propstat{
			Prop: &st.Prop{
				Getcontentlength: getcontentlength,
				// Getcontenttype: &st.Getcontenttype{
				// 	Innerxml: "text/plain; charset=utf-8",
				// },
				Resourcetype: &st.Resourcetype{},
				Displayname: &st.Displayname{
					Innerxml: node.Name,
				},
				Getlastmodified: &st.Getlastmodified{
					Innerxml: now,
				},
				Getetag: &st.Getetag{
					Innerxml: "\"" + node.ID + "\"",
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
	}
}

func graphql2DAVStructure2(graphql *graphql.Graphql, url string, now string, fileInfoFunc func(string) string) *structure.Multistatus {
	tree := graphql.Data.Project.Repository.Tree
	return &st.Multistatus{
		D:         "DAV:",
		Responses: treesAndBlobs2DAVResponses2(tree.Trees.Nodes, tree.Blobs.Nodes, url, now, fileInfoFunc),
	}
}

func treesAndBlobs2DAVResponses2(trees []*graphql.Node, blobs []*graphql.Node, url string, now string, fileInfoFunc func(string) string) []*structure.Response {
	lenTree := len(trees)
	responses := make([]*structure.Response, lenTree+len(blobs))
	for index, tree := range trees {
		responses[index] = tree2DAVResponse(tree, url, now)
	}
	brChan := make(chan *structure.Response)
	for _, blob := range blobs {
		go blob2DAVResponse2WithChan(blob, url, now, fileInfoFunc(blob.ID), brChan)
	}
	defer close(brChan)
	for index := range blobs {
		response := <-brChan
		responses[lenTree+index] = response
	}
	return responses
}

func tree2DAVResponse2(tree *graphql.Node, url string, now string) *structure.Response {
	return &st.Response{
		Href: &st.Href{
			Innerxml: url + tree.Name,
		},
		Propstat: &st.Propstat{
			Prop: &st.Prop{
				Resourcetype: &st.Resourcetype{
					Collection: &st.Collection{
						D: "DAV:",
					},
				},
				Displayname: &st.Displayname{
					Innerxml: tree.Name,
				},
				Getlastmodified: &st.Getlastmodified{
					Innerxml: now,
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
	}
}

func blob2DAVResponse2(blob *graphql.Node, url string, now string, size string) *structure.Response {
	var getcontentlength *st.Getcontentlength
	if size == "" {
		getcontentlength = nil
	} else {
		getcontentlength = &st.Getcontentlength{
			Innerxml: size,
		}
	}
	return &st.Response{
		Href: &st.Href{
			Innerxml: url + blob.Name,
		},
		Propstat: &st.Propstat{
			Prop: &st.Prop{
				Getcontentlength: getcontentlength,
				// Getcontenttype: &st.Getcontenttype{
				// 	Innerxml: "text/plain; charset=utf-8",
				// },
				Resourcetype: &st.Resourcetype{},
				Displayname: &st.Displayname{
					Innerxml: blob.Name,
				},
				Getlastmodified: &st.Getlastmodified{
					Innerxml: now,
				},
				Getetag: &st.Getetag{
					Innerxml: "\"" + blob.ID + "\"",
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
	}
}

func blob2DAVResponse2WithChan(blob *graphql.Node, url string, now string, size string, brChan chan<- *structure.Response) {
	var getcontentlength *st.Getcontentlength
	if size == "" {
		getcontentlength = nil
	} else {
		getcontentlength = &st.Getcontentlength{
			Innerxml: size,
		}
	}
	brChan <- &st.Response{
		Href: &st.Href{
			Innerxml: url + blob.Name,
		},
		Propstat: &st.Propstat{
			Prop: &st.Prop{
				Getcontentlength: getcontentlength,
				// Getcontenttype: &st.Getcontenttype{
				// 	Innerxml: "text/plain; charset=utf-8",
				// },
				Resourcetype: &st.Resourcetype{},
				Displayname: &st.Displayname{
					Innerxml: blob.Name,
				},
				Getlastmodified: &st.Getlastmodified{
					Innerxml: now,
				},
				Getetag: &st.Getetag{
					Innerxml: "\"" + blob.ID + "\"",
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
	}
}
