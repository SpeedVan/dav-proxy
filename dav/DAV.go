package dav

import "github.com/SpeedVan/go-common/app"

// DAV todo
type DAV interface {
	app.App
	DAVMethod
}
