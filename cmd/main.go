package main

import (
	"home/leonid/Git/Pract/network/pkg/database"
	"home/leonid/Git/Pract/network/pkg/service"
	"home/leonid/Git/Pract/network/pkg/transport"
	"log"
)

func main() {
	db := database.NewDB()
	svc := service.NewService(db)

	c := transport.NewTransport("8080", svc)
	c.InitEndpoints()

	err := c.Start()
	if err != nil {
		log.Println(err)
	}
}
