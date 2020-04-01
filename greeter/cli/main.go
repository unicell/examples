package main

import (
	"context"
	"fmt"
	"time"

	hello "github.com/micro/examples/greeter/srv/proto/hello"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
)

func main() {
	// create a new service
	service := micro.NewService()

	// parse command line flags
	service.Init()

	service.Client().Init(
		client.ContentType("application/json"),
	)

	// Use the generated client stub
	cl := hello.NewSayService("go.micro.srv.greeter", service.Client())

	for {
		d := time.Now()

		fmt.Println("making request")
		// Make request
		rsp, err := cl.Hello(context.Background(), &hello.Request{
			Name: "John",
		}, client.WithAddress("127.0.0.1:9091"))
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%v %s\n", time.Since(d), rsp.Msg)
		time.Sleep(time.Millisecond * 100)
	}
}
