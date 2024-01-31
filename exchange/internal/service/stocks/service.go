package stocks

import exchangereader "github.com/akrovv/exchange/internal/adapter/exchange-reader"

type service struct {
	storage exchangereader.ExchangeReader
}

func NewStockService(storage exchangereader.ExchangeReader) *service {
	return &service{storage: storage}
}

func (s *service) Create(DealData) (int64, int64) {
	panic("not implemented")
}

func (s *service) Cancel(CancelData) bool {
	panic("not implemented")
}

func (s *service) Results(BrokerData) <-chan DealData {
	panic("not implemented")
}
