package main

import (
	"fmt"

	"github.com/micro/go-micro/service/grpc"
)

func main() {
	service := grpc.NewService()

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
