package exchangereader

type storage struct{}

func NewStorage() *storage {
	return &storage{}
}

func (s *storage) GetStockInfo(string) <-chan DealData {
	panic("not implemented")
}
