package discount

type Discounter interface {
	Discount(uom string, quantities []float64) (original, discounted float64, err error)
}