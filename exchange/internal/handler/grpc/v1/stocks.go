package v1

import (
	"context"

	"github.com/akrovv/exchange/internal/models/protos/gen/go/exchange"
	"github.com/akrovv/exchange/internal/service/statistics"
	"github.com/akrovv/exchange/internal/service/stocks"
	"github.com/akrovv/exchange/pkg/logger"
)

type handler struct {
	exchange.UnimplementedExchangeServer

	logger       logger.Logger
	stockService stocks.StockService
	statService  statistics.StatisticService
}

func NewHandler(
	logger logger.Logger,
	stockService stocks.StockService,
	statService statistics.StatisticService,
) *handler {
	return &handler{
		logger:       logger,
		stockService: stockService,
		statService:  statService,
	}
}

func (h *handler) Create(ctx context.Context, in *exchange.Deal) (*exchange.DealID, error) {
	createDTO := stocks.DealData{
		ID:       in.ID,
		Type:     in.Type,
		BrokerID: in.BrokerID,
		ClientID: in.ClientID,
		Ticker:   in.Ticker,
		Amount:   in.Amount,
		Time:     in.Time,
		Price:    float64(in.Price),
	}

	id, brokerID, err := h.stockService.Create(&createDTO)

	if err != nil {
		h.logger.Fatalf("stock service with method create returned: %w", err)
		return nil, err
	}

	return &exchange.DealID{ID: id, BrokerID: brokerID}, nil
}

func (h *handler) Cancel(ctx context.Context, in *exchange.DealID) (*exchange.CancelResult, error) {
	cancelDTO := stocks.CancelData{ID: in.ID}

	success, err := h.stockService.Cancel(&cancelDTO)

	if err != nil {
		h.logger.Fatalf("stock service with method cancel returned: %w", err)
		return nil, err
	}

	return &exchange.CancelResult{Success: success}, nil
}

func (h *handler) Results(bi *exchange.BrokerID, stream exchange.Exchange_ResultsServer) error {
	resultsDTO := stocks.BrokerData{
		BrokerID: bi.ID,
	}

	resultsChan, err := h.stockService.Results(&resultsDTO)

	if err != nil {
		h.logger.Fatalf("stock service with method results returned: %w", err)
		return err
	}

	for data := range resultsChan {
		err = stream.Send(&exchange.Deal{ID: data.ID})

		if err != nil {
			h.logger.Fatalf("stream sends error: %w", err)
			return err
		}
	}

	return nil
}

func (h *handler) Statistic(bi *exchange.BrokerID, stream exchange.Exchange_StatisticServer) error {
	statisticDTO := statistics.StatisticlData{
		ID: bi.ID,
	}

	statisticChan, err := h.statService.Statistic(&statisticDTO)

	if err != nil {
		h.logger.Fatalf("statistic service with method results returned: %w", err)
		return err
	}

	for data := range statisticChan {
		err = stream.Send(&exchange.OHLCV{ID: data.ID})

		if err != nil {
			h.logger.Fatalf("stream sends error: %w", err)
			return err
		}
	}

	return nil
}
