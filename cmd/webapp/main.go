package main

import (
	"log"

	"github.com/MakMoinee/gobrankas/cmd/webapp/config"
	"github.com/MakMoinee/gobrankas/cmd/webapp/routes"
	"github.com/MakMoinee/gobrankas/internal/gobrankas/common"
	"github.com/MakMoinee/gobrankas/internal/pkg/localhttp"
)

func main() {
	config.Set()

	httpService := localhttp.NewService(common.SERVER_PORT)
	routes.Set(httpService)
	log.Println("Server started at port " + common.SERVER_PORT)
	if err := httpService.Start(); err != nil {
		panic(err)
	}
}
