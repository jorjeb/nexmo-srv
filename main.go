package main

import (
	"log"

	"github.com/jorjeb/nexmo-srv/handler"
	"github.com/jorjeb/nexmo-srv/nexmo"
	sms "github.com/jorjeb/nexmo-srv/proto/sms"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.nexmo"),
		micro.Flags(
			cli.StringFlag{
				Name:   "nexmo_api_key",
				EnvVar: "NEXMO_API_KEY",
				Usage:  "Nexmo API key",
			},
			cli.StringFlag{
				Name:   "nexmo_api_secret",
				EnvVar: "NEXMO_API_SECRET",
				Usage:  "Nexmo API secret",
			},
		),
		micro.Action(func(c *cli.Context) {
			nexmo.ApiKey = c.String("nexmo_api_key")
			nexmo.ApiSecret = c.String("nexmo_api_secret")
		}),
	)

	service.Init()
	nexmo.Init()

	sms.RegisterSMSHandler(service.Server(), new(handler.SMS))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
