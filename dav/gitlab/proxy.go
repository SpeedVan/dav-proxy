package gitlab

import (
	"net/http"

	"github.com/SpeedVan/go-common/config"
)

// DAVProxy todo
type DAVProxy struct {
	DAV
	GitlabAPIUrl string
	PrimaryToken string
}

// New todo
func New(config config.Config) {

}

// Run todo
func (s *DAVProxy) Run() error {

	return http.ListenAndServe(s.Address, s.Handler)
}
