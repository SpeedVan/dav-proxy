package main

import (
	"log"

	"github.com/SpeedVan/dav-proxy/dav/gitlab"
	"github.com/SpeedVan/dav-proxy/dav/rootdir"
	"github.com/SpeedVan/go-common/app/web"
	"github.com/SpeedVan/go-common/config"
	"github.com/SpeedVan/go-common/config/env"
)

func main() {
	if cfg, err := env.LoadAllWithoutPrefix("MOUNT_"); err == nil {
		app := web.New(cfg)

		app.HandleController(rootdir.New(cfg))
		cfg.ForEachArrayConfig("DIR", func(c config.Config) {
			switch c.Get("TYPE") {
			// case "file":
			// app.HandleFunc(filesystem.NewHandleFunc(c.Get("NAME"), c.WithPrefix("WEBDAV_")))
			case "http":
				if controller, err := gitlab.New(c); err == nil {
					app.HandleController(controller)
				}
			}
		})
		log.Println("start")
		log.Fatal(app.Run())
	} else {
		log.Fatal(err)
	}
}
