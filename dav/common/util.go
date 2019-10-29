package common

import (
	"net/http"

	"github.com/SpeedVan/go-common/app/web"
)

// DefaultDavReadonlyMethodsRouteMapBuilder todo
func DefaultDavReadonlyMethodsRouteMapBuilder(
	path string,
	head func(w http.ResponseWriter, r *http.Request),
	get func(w http.ResponseWriter, r *http.Request),
	propfind func(w http.ResponseWriter, r *http.Request),
) web.RouteMap {
	items := []*web.RouteItem{
		&web.RouteItem{Path: path, Method: "HEAD", HandleFunc: head},
		&web.RouteItem{Path: path, Method: "GET", HandleFunc: get},
		&web.RouteItem{Path: path, Method: "PROPFIND", HandleFunc: propfind},
		&web.RouteItem{Path: path, Method: "OPTIONS", HandleFunc: options},
	}

	return web.NewRouteMap(items...)
}

func options(w http.ResponseWriter, r *http.Request) {
	println("dav Options")
	header := w.Header()
	header.Set("Allow", "OPTIONS, PROPFIND")
	header.Set("Dav", "1, 2, 3")
	header.Set("Ms-Author-Via", "DAV")
}
