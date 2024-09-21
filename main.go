package main

import (
	"template/src/core/config"
	log "template/src/utils/logs"
)

func main() {
	log.InitLogger()
	config.Setup()
}
