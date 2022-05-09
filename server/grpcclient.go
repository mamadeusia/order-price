package server

import (
	"context"
	"preh/config"
	"preh/server/grpc/pricegrpc"

	"google.golang.org/grpc"
)

var (
	// StartUserGrpcClient starts gRPC client for user grpc
	StartUserGrpcClient = startUserGrpcClient
)

func startUserGrpcClient() (err error) {
	address := config.GetUserGRPCConnectionString()

	opts := grpc.WithInsecure()
	clientConn, err := grpc.Dial(address, opts)
	if err != nil {
		return
	}

	pricegrpcClient = pricegrpc.NewPriceClient(clientConn)

	return
}

var pricegrpcClient pricegrpc.PriceClient

func GetCurrentPrice(ctx context.Context, pair string) (price int64, err error) {
	request := &pricegrpc.GetCurrentPriceReq{Pair: pair}

	response, err := pricegrpcClient.GetCurrentPrice(ctx, request)
	if err != nil {
		return
	}

	price = response.Price

	return
}
