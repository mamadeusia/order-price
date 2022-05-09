package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"preh/config"
	"preh/core"
	"preh/server/grpc/pricegrpc"
	"syscall"

	"google.golang.org/grpc"
)

var (
	// StartGrpcServer starts listening to gRPC requests
	StartGrpcServer = startGrpcServer
)

type pricegrpcServer struct {
	pricegrpc.UnimplementedPriceServer
}

func startGrpcServer() {
	url := config.GetGrpcConnectionString()

	lis, err := net.Listen("tcp", url)
	if err != nil {
		return
	}

	server := grpc.NewServer()
	pricegrpc.RegisterPriceServer(server, &pricegrpcServer{})

	go func() {
		// Listen to operating system's interrupt signal
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
		<-interrupt

		// Gracefully shut down the server when it happens
		server.GracefulStop()
	}()
	server.Serve(lis)

	return
}

func (d *pricegrpcServer) GetCurrentPrice(ctx context.Context, request *pricegrpc.GetCurrentPriceReq) (respone *pricegrpc.GetCurrentPriceRes, err error) {
	fmt.Println("get Deposit Address ")
	currentPrice := core.GetCurrenctPairPrice(request.Pair)
	respone = &pricegrpc.GetCurrentPriceRes{
		Price: currentPrice,
	}
	return
}
