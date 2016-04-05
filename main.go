package main

import (
	"fmt"
	"log"
	"os"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"

	"github.com/olivere/greeter/handler"
	proto "github.com/olivere/greeter/proto/greeter"
)

func main() {
	var (
		id   string
		name string
	)

	// Specify the properties and arguments of the service
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
		micro.Flags(
			cli.BoolFlag{
				Name:   "client",
				EnvVar: "MACRO_CLIENT",
				Usage:  "Run as client",
			},
			cli.StringFlag{
				Name:   "id",
				EnvVar: "MACRO_ID",
				Usage:  "ID of server",
			},
			cli.StringFlag{
				Name:  "name",
				Usage: "User name",
			},
		),
	)

	// Initialize the service and arguments
	service.Init(
		micro.Action(func(c *cli.Context) {
			if c.IsSet("id") {
				id = c.String("id")
			}
			if c.IsSet("name") {
				name = c.String("name")
			}
			if c.Bool("client") {
				runClient(service, name)
				os.Exit(0)
			}
		}),
	)

	// Register the protobuf handler
	proto.RegisterGreeterHandler(service.Server(), &handler.Greeter{ID: id})

	// Run the service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

// runClient simply runs the client-side and calls the service.
func runClient(service micro.Service, name string) {
	// Create a client
	greeter := proto.NewGreeterClient("greeter", service.Client())

	// Invoke the Say rpc call
	ctx := context.Background()
	res, err := greeter.Say(ctx, &proto.SayRequest{Name: name})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.Reply)
}
