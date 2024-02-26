package stocks

import (
	"time"

	exchangereader "github.com/akrovv/exchange/internal/adapter/exchange-reader"
)

type service struct {
	storage exchangereader.ExchangeReader
}

func NewStockService(storage exchangereader.ExchangeReader) *service {
	return &service{storage: storage}
}

func (s *service) Create(dto *DealData) (int64, int64, error) {
	return 0, 0, nil
}

func (s *service) Cancel(dto *CancelData) (bool, error) {
	return true, nil
}

func (s *service) Results(dto *BrokerData) (<-chan DealData, error) {
	c := make(chan DealData)
	go func() {
		for i := 0; i < 10; i++ {
			c <- DealData{ID: int64(i)}
			time.Sleep(time.Second)
		}
		close(c)
	}()

	return c, nil
}
