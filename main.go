package main

import (
	"fmt"
	"log"

	"github.com/SpeedVan/dav-proxy/dav/gitlab"
	"github.com/SpeedVan/go-common/app/web"
	"github.com/SpeedVan/go-common/config"
	"github.com/SpeedVan/go-common/config/env"
)

func main() {
	if cfg, err := env.LoadAllWithoutPrefix("MOUNT_"); err == nil {
		// fmt.Println(cfg)
		app := web.New(cfg)

		cfg.ForEachArrayConfig("DIR", func(c config.Config) {
			switch c.Get("TYPE") {
			case "file":
				// app.HandleFunc(filesystem.NewHandleFunc("/"+c.Get("NAME")+"/", c.WithPrefix("WEBDAV_")))
			case "http":
				fmt.Println(c)
				app.HandleFunc(gitlab.NewHandleFunc("/http/"+c.Get("NAME")+"/{group}/{project}/{sha}/{path:.*}", c.WithPrefix("GITLAB_")))
			}
		})
		log.Println("start")
		log.Fatal(app.Run())
	} else {
		log.Fatal(err)
	}
}
