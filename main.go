package main

import (
	"log"

	"github.com/SpeedVan/go-common/config/env"
	"github.com/SpeedVan/proxy-in-dav/dav/filesystem"
)

func main() {
	if config, err := env.LoadAll(); err == nil {
		if fs, err := filesystem.New(config); err == nil {
			log.Println("start")
			log.Fatal(fs.Run())
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}
