package gitlab

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/astaxie/beego/cache"
	"github.com/gorilla/mux"

	"github.com/SpeedVan/dav-proxy/dav"
	"github.com/SpeedVan/dav-proxy/dav/common"
	"github.com/SpeedVan/go-common/app/web"
	"github.com/SpeedVan/go-common/client/httpclient/gitlab"
)

// PathProxy todo
type PathProxy struct {
	dav.ReadonlyMethod
	web.Controller
	Name             string
	Group            string
	Project          string
	Sha              string
	Path             string
	GitlabHTTPClient *gitlab.Client
	Cache            *cache.Cache
	FullFileInfo     bool
}

// NewPathProxy todo
func NewPathProxy(cl *gitlab.Client) *PathProxy {
	return &PathProxy{
		GitlabHTTPClient: cl,
	}
}

// GetRoute todo
func (s *PathProxy) GetRoute() web.RouteMap {
	return common.DefaultDavReadonlyMethodsRouteMapBuilder(
		"/"+s.Name+"/{group}/{project/{sha}/{path:.*}",
		s.Head,
		s.Get,
		s.Propfind,
		s.Options,
	)
}

// Head 根据gitlab的restApi /api/v4/projects/:id/repository/tree 效果可知，文件夹存不存在都无法直接得知，我们需要向上找一层，然后再判断是否存在列表中
func (s *PathProxy) Head(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// fmt.Println("head")
	// fmt.Println(vars["protocol"])
	// fmt.Println(vars["domain"])
	fmt.Println(vars["group"])
	fmt.Println(vars["project"])
	fmt.Println(vars["sha"])
	fmt.Println("Path:" + vars["path"])
	header := w.Header()
	header.Set("Accept-Ranges", "bytes")
	header.Set("Content-Length", "18")
	header.Set("Content-Type", "text/plain; charset=utf-8")
	header.Set("Etag", "\"13442cef32eaa60012\"")
	header.Set("Last-Modified", "Sun, 29 Dec 2013 02:26:31 GMT")
	header.Set("Date", "Mon, 30 Sep 2019 02:08:43 GMT")

	w.WriteHeader(200)
}

// Get todo
func (s *PathProxy) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reader, resHeader, err := s.GitlabHTTPClient.GetFile("http", vars["group"], vars["project"], vars["sha"], vars["path"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	header := w.Header()
	header.Set("Accept-Ranges", "bytes")
	header.Set("Content-Length", resHeader.Get("X-Gitlab-Size"))
	header.Set("Content-Type", resHeader.Get("Content-Type"))
	header.Set("Etag", resHeader.Get("X-Gitlab-Blob-Id"))
	header.Set("Last-Modified", resHeader.Get("Date"))
	header.Set("Date", resHeader.Get("Date"))
	io.Copy(w, reader)
}

// Propfind todo
func (s *PathProxy) Propfind(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	treeNodes, err := s.GitlabHTTPClient.GetTree("http", vars["group"], vars["project"], vars["sha"], vars["path"])
	// graphql, err := s.GitlabHTTPClient.Graphql(vars["protocol"], vars["group"], vars["project"], vars["sha"], vars["path"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fileInfoFunc := func(string) string {
		return ""
	}
	if s.FullFileInfo {
		fileInfoFunc = func(blodID string) string {
			size, err := s.GitlabHTTPClient.GetBlobSizeFromBody("http", vars["group"], vars["project"], blodID)
			if err != nil {
				return ""
			}
			return size
		}
	}
	bytes, err := xml.Marshal(treeNodes2DAVStructure2(treeNodes, r.URL.Path, "Fri, 27 Sep 2019 11:42:40 GMT", fileInfoFunc))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	header := w.Header()
	header.Set("Content-Type", "text/xml; charset=utf-8")
	header.Set("Transfer-Encoding", "chunked")
	w.WriteHeader(207)
	w.Write([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>"))

	w.Write(bytes)
}

// Options todo
func (s *PathProxy) Options(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Set("Allow", "OPTIONS, PROPFIND")
	header.Set("Dav", "1, 2")
	header.Set("Ms-Author-Via", "DAV")
	w.Write(dataDav)
}
