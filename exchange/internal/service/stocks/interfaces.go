package stocks

type StockService interface {
	Create(DealData) (int64, int64)
	Cancel(CancelData) bool
	Results(BrokerData) <-chan DealData
}
