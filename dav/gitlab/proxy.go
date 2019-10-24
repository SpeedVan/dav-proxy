package gitlab

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/astaxie/beego/cache"
	"github.com/gorilla/mux"

	"github.com/SpeedVan/dav-proxy/dav"
	"github.com/SpeedVan/go-common/client/httpclient/gitlab"
	"github.com/SpeedVan/go-common/config"
)

var (
	// EmptyHeader empty header
	EmptyHeader = http.Header{}
)

// DAVProxy todo
type DAVProxy struct {
	dav.DAV
	GitlabHTTPClient *gitlab.Client
	Cache            *cache.Cache
	FullFileInfo     bool
}

// New todo
func New(config config.Config) (*DAVProxy, error) {

	gitlabHTTPClient, err := gitlab.New(config)
	if err != nil {
		return nil, err
	}
	fullFileInfo, err := strconv.ParseBool(config.Get("FULL_FILEINFO"))
	if err != nil {
		fullFileInfo = false
	}
	return &DAVProxy{
		GitlabHTTPClient: gitlabHTTPClient,
		FullFileInfo:     fullFileInfo,
	}, nil
}

// NewHandleFunc todo
func NewHandleFunc(path string, config config.Config) (string, func(http.ResponseWriter, *http.Request)) {
	o, err := New(config)
	if err != nil {
		log.Fatal(err)
	}

	//localhost:8887/{protocol:(http|https)}/{domain}/{group}/{project}/{sha}/{path:.*} liunx挂载proxy服务地址

	return path, func(w http.ResponseWriter, r *http.Request) {
		// url := r.URL.Path
		println("method:" + r.Method + ",url:" + r.URL.Path)
		switch r.Method {
		case "HEAD":
			o.Head(w, r)
		case "GET":
			o.Get(w, r)
		case "PROPFIND":
			o.Propfind(w, r)
		case "OPTIONS":
			o.Options(w, r)
		default:
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		}
	}
}

// Head 根据gitlab的restApi /api/v4/projects/:id/repository/tree 效果可知，文件夹存不存在都无法直接得知，我们需要向上找一层，然后再判断是否存在列表中
func (s *DAVProxy) Head(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("head")
	fmt.Println(vars["protocol"])
	fmt.Println(vars["domain"])
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
func (s *DAVProxy) Get(w http.ResponseWriter, r *http.Request) {
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
func (s *DAVProxy) Propfind(w http.ResponseWriter, r *http.Request) {
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
func (s *DAVProxy) Options(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Set("Allow", "OPTIONS, PROPFIND")
	header.Set("Dav", "1, 2")
	header.Set("Ms-Author-Via", "DAV")
	w.Write(dataDav)
}
