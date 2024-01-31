package exchangereader

type ExchangeReader interface {
	GetStockInfo(string) <-chan DealData
}
