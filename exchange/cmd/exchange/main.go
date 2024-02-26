package main

import (
	"net"

	exchangereader "github.com/akrovv/exchange/internal/adapter/exchange-reader"
	v1 "github.com/akrovv/exchange/internal/handler/grpc/v1"
	"github.com/akrovv/exchange/internal/handler/grpc/v1/interceptors"
	"github.com/akrovv/exchange/internal/models/protos/gen/go/exchange"
	"github.com/akrovv/exchange/internal/service/statistics"
	"github.com/akrovv/exchange/internal/service/stocks"
	"github.com/akrovv/exchange/pkg/logger"
	"google.golang.org/grpc"
)

func main() {
	logger := logger.NewLogger()

	storage := exchangereader.NewStorage()
	stockService := stocks.NewStockService(storage)
	statService := statistics.NewStatService()

	handler := v1.NewHandler(logger, stockService, statService)

	server := grpc.NewServer(grpc.UnaryInterceptor(interceptors.Logger(logger)))
	exchange.RegisterExchangeServer(server, handler)

	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		logger.Fatalf("can't start server: %w", err.Error())
		return
	}

	logger.Info("starting on :8080")

	if err = server.Serve(lis); err != nil {
		logger.Fatal(err.Error())
	}
}
