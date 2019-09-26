package main

import (
	"log"
	"path/filepath"

	"github.com/SpeedVan/go-common/config/env"
	"github.com/SpeedVan/proxy-in-dav/dav"
)

func main() {
	path := "/userfunc/"
	maybeEnvConfig := env.LoadAll()

	if config, ok := maybeEnvConfig.(*env.EnvConfig); ok {
		if s, err := filepath.Abs(path); err == nil {
			log.Println("start")
			dav.New(s).Run(config)
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(maybeEnvConfig.(error))
	}

}
