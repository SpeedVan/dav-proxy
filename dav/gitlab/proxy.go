package gitlab

import (
	"net/http"

	"github.com/SpeedVan/go-common/client/gitlabclient"
	"github.com/SpeedVan/go-common/config"
	"github.com/SpeedVan/proxy-in-dav/dav"
)

// DAVProxy todo
type DAVProxy struct {
	dav.DAV
	Address          string
	GitlabHTTPClient *gitlabclient.GitlabClient
}

// New todo
func New(config config.Config) (dav.DAV, error) {

	gitlabHTTPClient, err := gitlabclient.New(config)
	if err != nil {
		return nil, err
	}

	return &DAVProxy{
		GitlabHTTPClient: gitlabHTTPClient,
	}, nil
}

// Head todo
func (s *DAVProxy) Head(w http.ResponseWriter, r *http.Request) error {
	r.URL.Path
	// s.GitlabHTTPClient
	return nil
}

func (s *DAVProxy) Get(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Run todo
func (s *DAVProxy) Run() error {

	return http.ListenAndServe(s.Address, nil)
}
