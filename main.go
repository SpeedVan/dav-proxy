package main

import (
	"log"

	"github.com/SpeedVan/dav-proxy/dav/filesystem"
	"github.com/SpeedVan/dav-proxy/dav/gitlab"
	"github.com/SpeedVan/go-common/app/web"
	"github.com/SpeedVan/go-common/config/env"
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
