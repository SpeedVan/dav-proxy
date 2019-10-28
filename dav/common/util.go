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
	options func(w http.ResponseWriter, r *http.Request),
) web.RouteMap {
	items := []*web.RouteItem{
		&web.RouteItem{Path: path, Methods: web.RouteMethods{"Head"}, HandleFunc: head},
		&web.RouteItem{Path: path, Methods: web.RouteMethods{"Get"}, HandleFunc: get},
		&web.RouteItem{Path: path, Methods: web.RouteMethods{"Propfind"}, HandleFunc: propfind},
		&web.RouteItem{Path: path, Methods: web.RouteMethods{"Options"}, HandleFunc: options},
	}

	return web.NewRouteMap(items...)
}
