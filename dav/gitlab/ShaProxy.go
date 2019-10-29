package gitlab

import (
	"encoding/xml"
	"net/http"

	"github.com/SpeedVan/go-common/cache"
	"github.com/gorilla/mux"

	"github.com/SpeedVan/dav-proxy/dav"
	"github.com/SpeedVan/go-common/app/web"
	"github.com/SpeedVan/go-common/client/httpclient/gitlab"
)

// ShaProxy todo
type ShaProxy struct {
	dav.ReadonlyMethod
	web.Controller
	Name             string
	Group            string
	Project          string
	Sha              string
	GitlabHTTPClient *gitlab.Client
	Cache            cache.StreamClient
	FullFileInfo     bool
}

// NewShaProxy todo
func NewShaProxy(cl *gitlab.Client) *ShaProxy {
	return &ShaProxy{
		GitlabHTTPClient: cl,
	}
}

// Head todo
func (s *ShaProxy) Head(w http.ResponseWriter, r *http.Request) {

}

// Get todo
func (s *ShaProxy) Get(w http.ResponseWriter, r *http.Request) {

}

// Propfind todo
func (s *ShaProxy) Propfind(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	treeNodes, err := s.GitlabHTTPClient.GetTree("http", vars["group"], vars["project"], vars["sha"], "/")
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
func (s *ShaProxy) Options(w http.ResponseWriter, r *http.Request) {

}
