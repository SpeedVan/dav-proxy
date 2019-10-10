package gitlab

import (
	"github.com/SpeedVan/go-common/client/httpclient/gitlab"
	"github.com/SpeedVan/go-common/client/httpclient/gitlab/graphql"
	"github.com/SpeedVan/proxy-in-dav/dav/structure"
	st "github.com/SpeedVan/proxy-in-dav/dav/structure"
)

// treeNodes2DAVStructure
func treeNodes2DAVStructure(nodes []*gitlab.TreeNode, url string, now string) *structure.Multistatus {
	return &st.Multistatus{
		D:         "DAV:",
		Responses: treeNodes2DAVResponses(nodes, url, now),
	}
}

// treeNodes2DAVResponses mean map
func treeNodes2DAVResponses(nodes []*gitlab.TreeNode, url string, now string) []*structure.Response {
	responses := make([]*structure.Response, len(nodes))
	for index, node := range nodes {
		responses[index] = treeNode2DAVResponse(node, url, now)
	}
	return responses
}

// treeNode2DAVResponse
func treeNode2DAVResponse(node *gitlab.TreeNode, url string, now string) *structure.Response {
	if node.Type == "tree" {
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
	return &st.Response{
		Href: &st.Href{
			Innerxml: url + node.Name,
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

func graphql2DAVStructure(graphql *graphql.Graphql, url string, now string) *structure.Multistatus {
	tree := graphql.Data.Project.Repository.Tree
	return &st.Multistatus{
		D:         "DAV:",
		Responses: treesAndBlobs2DAVResponses(tree.Trees.Nodes, tree.Blobs.Nodes, url, now),
	}
}

func treesAndBlobs2DAVResponses(trees []*graphql.Node, blobs []*graphql.Node, url string, now string) []*structure.Response {
	lenTree := len(trees)
	responses := make([]*structure.Response, lenTree+len(blobs))
	for index, tree := range trees {
		responses[index] = tree2DAVResponse(tree, url, now)
	}
	for index, blob := range blobs {
		responses[lenTree+index] = blob2DAVResponse(blob, url, now)
	}
	return responses
}

func tree2DAVResponse(tree *graphql.Node, url string, now string) *structure.Response {
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

func blob2DAVResponse(blob *graphql.Node, url string, now string) *structure.Response {
	return &st.Response{
		Href: &st.Href{
			Innerxml: url + blob.Name,
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
