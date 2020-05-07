package decorator

type ICoffe interface {
	Description() string
	Cost() float64
}

type ShortBlack struct {
	Desc string
	Cos  float64
}

func (s *ShortBlack) Description() string {
	return s.Desc
}

func (s *ShortBlack) Cost() float64 {
	return s.Cos
}
