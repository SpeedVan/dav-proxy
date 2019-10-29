package gitlab

import (
	"encoding/xml"
	"net/http"
	"strings"

	"github.com/SpeedVan/go-common/cache"
	"github.com/gorilla/mux"

	"github.com/SpeedVan/dav-proxy/dav"
	"github.com/SpeedVan/dav-proxy/dav/common"
	st "github.com/SpeedVan/dav-proxy/dav/structure"
	"github.com/SpeedVan/go-common/app/web"
	"github.com/SpeedVan/go-common/client/httpclient/gitlab"
)

// ProjectProxy todo
type ProjectProxy struct {
	dav.ReadonlyMethod
	web.Controller
	Name             string
	Group            string
	Project          string
	GitlabHTTPClient *gitlab.Client
	Cache            cache.StreamClient
}

// NewProjectProxy todo
func NewProjectProxy(cl *gitlab.Client) *ProjectProxy {
	return &ProjectProxy{
		GitlabHTTPClient: cl,
	}
}

// GetRoute todo
func (s *ProjectProxy) GetRoute() web.RouteMap {
	return common.DefaultDavReadonlyMethodsRouteMapBuilder(
		"/"+s.Name+"/{gpid}/",
		s.Head,
		s.Get,
		s.Propfind,
		s.Options,
	)
}

// Head todo
func (s *ProjectProxy) Head(w http.ResponseWriter, r *http.Request) {

}

// Get todo
func (s *ProjectProxy) Get(w http.ResponseWriter, r *http.Request) {

}

// Propfind todo
func (s *ProjectProxy) Propfind(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	gpid := vars["gpid"]
	gp := strings.Split(gpid, "%2F")

	commits, err := s.GitlabHTTPClient.GetCommits("http", gp[0], gp[1], "", "")

	responses := []*st.Response{}

	for _, item := range commits {
		responses = append(responses, st.ToDir(item.ID, "Fri, 27 Sep 2019 11:42:40 GMT"))
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

// Options todo
func (s *ProjectProxy) Options(w http.ResponseWriter, r *http.Request) {

}
