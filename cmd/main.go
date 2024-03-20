package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/peppydays/go-grpc-member/configs"
	"github.com/peppydays/go-grpc-member/internal/infrastructure/repository"
	"github.com/peppydays/go-grpc-member/internal/interface/rpc"
	"github.com/peppydays/go-grpc-member/internal/service"
	"github.com/peppydays/go-grpc-member/pkg/contract"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config := configs.Application{
		Host: "0.0.0.0",
		Port: 50051,
	}

	memberRepository := repository.NewMemoryRepository()
	memberService := service.NewMemberService(memberRepository)

	baseHandler := rpc.NewBaseHandler(memberService)
	rpcServer, err := rpc.NewServer(baseHandler, rpc.WithPort(int(config.Port)))
	if err != nil {
		panic(err)
	}

	contract.RegisterHealthServer(rpcServer.Server(), rpc.NewHealthHandler())

	reflection.Register(rpcServer.Server())

	go func() {
		err := rpcServer.Serve()
		if err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	select {
	case v := <-quit:
		fmt.Printf("signal notified: %v\n", v)
	case done := <-ctx.Done():
		fmt.Printf("context done: %v\n", done)
	}
}
