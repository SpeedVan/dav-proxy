package gitlab

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/astaxie/beego/cache"

	"github.com/SpeedVan/dav-proxy/dav"
	"github.com/SpeedVan/dav-proxy/dav/common"
	st "github.com/SpeedVan/dav-proxy/dav/structure"
	"github.com/SpeedVan/go-common/app/web"
	"github.com/SpeedVan/go-common/client/httpclient/gitlab"
)

// NameProxy todo
type NameProxy struct {
	dav.ReadonlyMethod
	web.Controller
	Name             string
	GitlabHTTPClient *gitlab.Client
	Cache            *cache.Cache
}

// NewNameProxy todo
func NewNameProxy(cl *gitlab.Client) *NameProxy {
	return &NameProxy{
		GitlabHTTPClient: cl,
	}
}

// GetRoute todo
func (s *NameProxy) GetRoute() web.RouteMap {
	return common.DefaultDavReadonlyMethodsRouteMapBuilder(
		"/"+s.Name+"/",
		s.Head,
		s.Get,
		s.Propfind,
		s.Options,
	)
}

// Head todo
func (s *NameProxy) Head(w http.ResponseWriter, r *http.Request) {

}

// Get todo
func (s *NameProxy) Get(w http.ResponseWriter, r *http.Request) {

}

// Propfind todo
func (s *NameProxy) Propfind(w http.ResponseWriter, r *http.Request) {

	projects, err := s.GitlabHTTPClient.GetProjects("http", group)

	responses := []*st.Response{}

	for _, item := range projects {
		responses = append(responses, st.ToDir(fmt.Sprint(item.Name), "Fri, 27 Sep 2019 11:42:40 GMT"))
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
func (s *NameProxy) Options(w http.ResponseWriter, r *http.Request) {

}
