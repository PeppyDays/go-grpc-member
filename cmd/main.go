package main

import (
	"context"
	"fmt"

	"github.com/peppydays/go-grpc-member/configs"
)

func main() {
	config, err := configs.LoadFromPath(context.Background(), "configs/Application.pkl")
	if err != nil {
		panic(err)
	}

	fmt.Println(config.Host, config.Port)
}
