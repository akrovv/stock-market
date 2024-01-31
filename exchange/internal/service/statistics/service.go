package statistics

type service struct {
}

func NewStatService() *service {
	return &service{}
}

func (s *service) Statistic(StatisticlData) <-chan OHLCVData {
	panic("not implemented")
}
