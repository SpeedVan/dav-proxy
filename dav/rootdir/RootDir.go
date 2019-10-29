package rootdir

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/SpeedVan/dav-proxy/dav"
	"github.com/SpeedVan/dav-proxy/dav/common"
	st "github.com/SpeedVan/dav-proxy/dav/structure"
	"github.com/SpeedVan/go-common/app/web"
	"github.com/SpeedVan/go-common/config"
	"github.com/SpeedVan/go-common/type/collection/omap"
)

// RootDir 根目录，内部为以每个配置的Name项作为名字的文件夹
type RootDir struct {
	dav.ReadonlyMethod
	web.Controller
	davResponse *st.Multistatus
	path        string
}

// New todo
func New(config config.Config) *RootDir {

	responses := []*st.Response{}

	config.GetMap("DIR").ForEach(func(name string, v interface{}) {
		item := v.(omap.Map)
		responses = append(responses, st.ToDir(fmt.Sprint(item.Get("NAME")), "Fri, 27 Sep 2019 11:42:40 GMT"))
	})

	ms := &st.Multistatus{
		D:         "DAV:",
		Responses: responses,
	}

	return &RootDir{
		davResponse: ms,
		path:        "/{name}/",
	}
}

// GetRoute todo
func (s *RootDir) GetRoute() web.RouteMap {
	return common.DefaultDavReadonlyMethodsRouteMapBuilder(
		"/",
		s.Head,
		s.Get,
		s.Propfind,
		s.Options,
	)
}

func (s *RootDir) Head(w http.ResponseWriter, r *http.Request) {

}

func (s *RootDir) Get(w http.ResponseWriter, r *http.Request) {

}

func (s *RootDir) Propfind(w http.ResponseWriter, r *http.Request) {

	bytes, err := xml.Marshal(s.davResponse)
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

func (s *RootDir) Options(w http.ResponseWriter, r *http.Request) {
	println("root Options")
	header := w.Header()
	header.Set("Allow", "OPTIONS, PROPFIND")
	header.Set("Dav", "1, 2")
	header.Set("Ms-Author-Via", "DAV")
}
