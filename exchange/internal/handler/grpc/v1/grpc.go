package v1

import (
	"github.com/akrovv/exchange/internal/models/protos/gen/go/exchange"
	"github.com/akrovv/exchange/internal/service/statistics"
	"github.com/akrovv/exchange/internal/service/stocks"
	"go.uber.org/zap"
)

type handler struct {
	exchange.UnimplementedExchangeServer

	logger       zap.Logger
	stockService stocks.StockService
	statService  statistics.StatisticService
}

func NewHandler(
	logger zap.Logger,
	stockService stocks.StockService,
	statService statistics.StatisticService,
) *handler {
	return &handler{
		logger:       logger,
		stockService: stockService,
		statService:  statService,
	}
}
