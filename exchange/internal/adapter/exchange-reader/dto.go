package exchangereader

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
