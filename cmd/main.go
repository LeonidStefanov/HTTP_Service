package main

import (
	"fmt"

	"github.com/LeonidStefanov/HTTP_Service/pkg/database"
	"github.com/LeonidStefanov/HTTP_Service/pkg/option"
	"github.com/LeonidStefanov/HTTP_Service/pkg/service"
	"github.com/LeonidStefanov/HTTP_Service/pkg/transport"
	"github.com/kelseyhightower/envconfig"

	"log"
)

func main() {

	var cfg option.Options
	err := envconfig.Process("service", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(cfg)

	db, err := database.NewDB(cfg.DBHost, cfg.DBPort, "skillbox")
	if err != nil {
		log.Println(err)
		return
	}

	defer db.Close()

	svc := service.NewService(db)

	c := transport.NewTransport(cfg.Port, svc)
	c.InitEndpoints()

	err = c.Start()
	if err != nil {
		log.Println(err)
	}
}
