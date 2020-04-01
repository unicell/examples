package main

import (
	"context"
	"fmt"

	"github.com/micro/cli/v2"
	proto "github.com/micro/examples/helloworld/proto"
	"github.com/micro/examples/mocking/mock"
	"github.com/micro/go-micro/v2"
)

func main() {
	var c proto.GreeterService

	service := micro.NewService(
		micro.Flags(cli.StringFlag{
			Name:  "environment",
			Value: "testing",
		}),
	)

	service.Init(
		micro.Action(func(ctx *cli.Context) {
			env := ctx.String("environment")
			// use the mock when in testing environment
			if env == "testing" {
				c = mock.NewGreeterService()
			} else {
				c = proto.NewGreeterService("helloworld", service.Client())
			}
		}),
	)

	// call hello service
	rsp, err := c.Hello(context.TODO(), &proto.HelloRequest{
		Name: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.Greeting)
}
