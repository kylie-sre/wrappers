package discount

type Discounter interface {
	Discount(uom string, quantities []float64) (original, discounted float64, err error)
}

type discounter int

func (d discounter) Discount(uom string, quantities []float64) (original, discounted float64, err error) {
	panic("implement me")
}

func New() Discounter {
	return discounter(1)
}
