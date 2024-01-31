package stocks

type BrokerData struct {
	BrokerID int64
}

type DealData struct {
	ID       int64
	Type     int64
	BrokerID int64
	ClientID int64
	Ticker   string
	Amount   int32
	Time     int32
	Price    float64
}

type CancelData struct {
	ID int64
}
