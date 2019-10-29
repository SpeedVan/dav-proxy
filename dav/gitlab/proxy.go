package gitlab

import (
	"net/http"
	"strconv"

	"github.com/SpeedVan/go-common/app/web"
	"github.com/SpeedVan/go-common/client/httpclient/gitlab"
	"github.com/SpeedVan/go-common/config"
)

var (
	// EmptyHeader empty header
	EmptyHeader = http.Header{}
)

// DAVProxy todo
type DAVProxy struct {
	web.Controller
	Domain       string
	NameProxy    web.Controller
	GroupProxy   web.Controller
	ProjectProxy web.Controller
	ShaProxy     web.Controller
	PathProxy    web.Controller
}

// New todo
func New(config config.Config) (*DAVProxy, error) {
	gitlabHTTPClient, err := gitlab.New(config.WithPrefix("GITLAB_"))
	if err != nil {
		return nil, err
	}
	fullFileInfo, err := strconv.ParseBool(config.Get("GITLAB_FULL_FILEINFO"))
	if err != nil {
		fullFileInfo = false
	}
	name := config.Get("NAME")
	return &DAVProxy{
		Domain:    name,
		NameProxy: &NameProxy{Name: name, GitlabHTTPClient: gitlabHTTPClient},
		// GroupProxy:   &GroupProxy{Name: name, GitlabHTTPClient: gitlabHTTPClient},
		ProjectProxy: &ProjectProxy{Name: name, GitlabHTTPClient: gitlabHTTPClient},
		// ShaProxy:     &ShaProxy{Name: name, GitlabHTTPClient: gitlabHTTPClient},
		PathProxy: &PathProxy{Name: name, GitlabHTTPClient: gitlabHTTPClient, FullFileInfo: fullFileInfo},
	}, nil
}

// GetRoute todo
func (s *DAVProxy) GetRoute() web.RouteMap {
	return web.MergeRouteMap(
		s.NameProxy.GetRoute(),
		// s.GroupProxy.GetRoute(),
		s.ProjectProxy.GetRoute(),
		// s.ShaProxy.GetRoute(),
		s.PathProxy.GetRoute(),
	)
}
