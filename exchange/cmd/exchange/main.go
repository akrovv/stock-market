package main

import (
	"log"
	"net"

	exchangereader "github.com/akrovv/exchange/internal/adapter/exchange-reader"
	v1 "github.com/akrovv/exchange/internal/handler/grpc/v1"
	"github.com/akrovv/exchange/internal/models/protos/gen/go/exchange"
	"github.com/akrovv/exchange/internal/service/statistics"
	"github.com/akrovv/exchange/internal/service/stocks"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := zap.NewProduction()

	if err != nil {
		log.Fatal(err)
	}

	storage := exchangereader.NewStorage()
	stockService := stocks.NewStockService(storage)
	statService := statistics.NewStatService()

	handler := v1.NewHandler(*logger, stockService, statService)
	server := grpc.NewServer()

	exchange.RegisterExchangeServer(server, handler)

	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		logger.Warn("can't start server" + err.Error())
		return
	}

	// startGRPCServer(server, lis)
	logger.Info("starting on :8080")
	err = server.Serve(lis)

	if err != nil {
		logger.Warn(err.Error())
	}
}
