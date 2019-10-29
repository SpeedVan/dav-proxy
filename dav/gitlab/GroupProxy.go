package gitlab

import (
	"encoding/xml"
	"net/http"

	"github.com/SpeedVan/go-common/cache"
	"github.com/gorilla/mux"

	"github.com/SpeedVan/dav-proxy/dav"
	"github.com/SpeedVan/dav-proxy/dav/common"
	st "github.com/SpeedVan/dav-proxy/dav/structure"
	"github.com/SpeedVan/go-common/app/web"
	"github.com/SpeedVan/go-common/client/httpclient/gitlab"
)

// GroupProxy todo
type GroupProxy struct {
	dav.ReadonlyMethod
	web.Controller
	Name             string
	GitlabHTTPClient *gitlab.Client
	Cache            cache.StreamClient
}

// NewGroupProxy todo
func NewGroupProxy(cl *gitlab.Client) *GroupProxy {
	return &GroupProxy{
		GitlabHTTPClient: cl,
	}
}

// GetRoute todo
func (s *GroupProxy) GetRoute() web.RouteMap {
	return common.DefaultDavReadonlyMethodsRouteMapBuilder(
		"/"+s.Name+"/{group}/",
		s.Head,
		s.Get,
		s.Propfind,
	)
}

// Head todo
func (s *GroupProxy) Head(w http.ResponseWriter, r *http.Request) {

}

// Get todo
func (s *GroupProxy) Get(w http.ResponseWriter, r *http.Request) {

}

// Propfind todo
func (s *GroupProxy) Propfind(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	group := vars["group"]

	projects, err := s.GitlabHTTPClient.GetProjects("http", group)

	responses := []*st.Response{}

	for _, item := range projects {
		responses = append(responses, st.ToDir(r.URL.Path, item.Name, "Fri, 27 Sep 2019 11:42:40 GMT"))
	}

	ms := &st.Multistatus{
		D:         "DAV:",
		Responses: responses,
	}

	bytes, err := xml.Marshal(ms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	header := w.Header()
	header.Set("Content-Type", "text/xml; charset=utf-8")
	header.Set("Transfer-Encoding", "chunked")
	w.WriteHeader(207)
	_, err = w.Write([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(bytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
