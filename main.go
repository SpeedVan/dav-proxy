package main

import (
	"log"

	"github.com/SpeedVan/go-common/app/web"
	"github.com/SpeedVan/go-common/config/env"
	"github.com/SpeedVan/proxy-in-dav/dav/filesystem"
	"github.com/SpeedVan/proxy-in-dav/dav/gitlab"
)

func main() {
	if config, err := env.LoadAll(); err == nil {
		app := web.New(config)
		app.
			HandleFunc(filesystem.NewHandleFunc("/file/", config.WithPrefix("WEBDAV_"))).
			HandleFunc(gitlab.NewHandleFunc("/{protocol}/{domain}/{group}/{project}/{sha}/{path:.*}", config.WithPrefix("GITLAB_")))
		log.Println("start")
		log.Fatal(app.Run())
	} else {
		log.Fatal(err)
	}
}
