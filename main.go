package main

import (
	"log"

	"github.com/micro/go-micro"
	proto "github.com/microhq/greeting-api/proto/hello"
	hello "github.com/microhq/greeting-srv/proto/hello"

	"context"
)

type Greeting struct {
	Client hello.SayService
}

func (g *Greeting) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	log.Print("Received Greeting.Hello API request")

	// make the request
	response, err := g.Client.Hello(ctx, &hello.Request{Name: req.Name})
	if err != nil {
		return err
	}

	// set api response
	rsp.Msg = response.Msg
	return nil
}

func main() {
	// Create service
	service := micro.NewService(
		micro.Name("go.micro.api.greeting"),
	)

	// Init to parse flags
	service.Init()

	// Register Handlers
	proto.RegisterGreetingHandler(service.Server(), &Greeting{
		// Create Service Client
		Client: hello.NewSayService("go.micro.srv.greeting", service.Client()),
	})

	// for handler use

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
