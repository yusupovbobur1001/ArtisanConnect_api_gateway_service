package main

import (
	"api_service/api"
	"api_service/config"
	"fmt"
)

func main() {
	cfg := config.Load()

	router := api.NewRouter(cfg)
	router.Run(fmt.Sprintf("localhost:%s",cfg.HTTP_PORT))
}