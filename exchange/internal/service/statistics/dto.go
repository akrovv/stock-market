package statistics

type StatisticlData struct {
	ID int64
}

type OHLCVData struct {
	ID       int64
	Time     int32
	Interval int32
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Ticker   string
	Volume   float64
}
