package main

import (
	orglog "log"

	"github.com/SpeedVan/dav-proxy/dav/filesystem"
	"github.com/SpeedVan/dav-proxy/dav/gitlab"
	"github.com/SpeedVan/dav-proxy/dav/rootdir"
	"github.com/SpeedVan/go-common/app/web"
	"github.com/SpeedVan/go-common/config"
	"github.com/SpeedVan/go-common/config/env"
	"github.com/SpeedVan/go-common/log"
	lc "github.com/SpeedVan/go-common/log/common"
)

func main() {
	if cfg, err := env.LoadAllWithoutPrefix("MOUNT_"); err == nil {
		logger := lc.NewCommon(log.Debug) // this level control webapp init log level

		app := web.New(cfg, logger)

		app.HandleController(rootdir.New(cfg, logger))
		cfg.ForEachArrayConfig("DIR", func(c config.Config) {
			switch c.Get("TYPE") {
			case "file":
				app.HandleController(filesystem.New(c, logger))
			case "http":
				if controller, err := gitlab.New(c, logger); err == nil {
					app.HandleController(controller)
				}
			}
		})
		orglog.Fatal(app.Run(log.Debug)) // this level control webapp runtime log level
	} else {
		orglog.Fatal(err)
	}
}
