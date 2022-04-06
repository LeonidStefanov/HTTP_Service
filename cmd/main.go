package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"

	"home/leonid/Git/Pract/network/pkg/database"
	"home/leonid/Git/Pract/network/pkg/option"
	"home/leonid/Git/Pract/network/pkg/service"
	"home/leonid/Git/Pract/network/pkg/transport"
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
