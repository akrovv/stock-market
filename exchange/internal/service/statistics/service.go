package statistics

type service struct {
}

func NewStatService() *service {
	return &service{}
}

func (s *service) Statistic(dto *StatisticlData) (<-chan OHLCVData, error) {
	panic("not implemented")
}
