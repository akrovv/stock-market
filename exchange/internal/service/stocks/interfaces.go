package stocks

type StockService interface {
	Create(*DealData) (int64, int64, error)
	Cancel(*CancelData) (bool, error)
	Results(*BrokerData) (<-chan DealData, error)
}
