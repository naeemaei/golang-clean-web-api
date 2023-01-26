package main

import (
	"log"

	"github.com/naeemaei/golang-clean-web-api/api"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/data/cache"
	"github.com/naeemaei/golang-clean-web-api/data/db"
)

func main() {
	cfg := config.GetConfig()
	
	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		log.Fatal(err)
	}

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		log.Fatal(err)
	}

	api.InitServer(cfg)
}
