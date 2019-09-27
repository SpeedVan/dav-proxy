package main

import (
	"log"

	"github.com/SpeedVan/go-common/config/env"
	"github.com/SpeedVan/proxy-in-dav/dav"
)

func main() {
	maybeEnvConfig := env.LoadAll()

	if config, ok := maybeEnvConfig.(*env.EnvConfig); ok {
		log.Println("start")
		log.Fatal(dav.New(config).Run())
	} else {
		log.Fatal(maybeEnvConfig.(error))
	}

}
