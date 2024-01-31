package statistics

type StatisticService interface {
	Statistic(StatisticlData) <-chan OHLCVData
}
