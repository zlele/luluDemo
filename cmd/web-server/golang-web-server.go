package main

import (
	"log"
	"wanglu/golang-web-server/configs"
	"wanglu/golang-web-server/pkg/mhttp"
)

var useHandlerFunctions bool

func main() {
	server := mhttp.NewServer(configs.SERVER_STATIC_DIRECTORY, configs.SERVER_URL, configs.SERVER_PORT)
	var err error

	err = server.InitializeHandlerFunctions()

	if err != nil {
		log.Fatal(err)
	}
	server.ListenAndServe()

}
