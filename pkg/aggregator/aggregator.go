package aggregator

type Option func(*aggregator)

func Quantity(val float64) Option {
	return func(a *aggregator) {
	
	}
}
func Discounted(val float64) Option {
	return func(a *aggregator) {
	
	}
}

type aggregator struct {
}

func New() *aggregator {
	return &aggregator{}
}

func (a *aggregator) Values() []byte {
	return []byte{}
}

func (a *aggregator) Add([]byte) error {
	return nil
}

func (a *aggregator) Get(opts ...Option) ([]byte, error) {
	return []byte{}, nil
}

type Aggregator interface {
	Values() []byte
	Add([]byte) error
	Get(opts ...Option) ([]byte, error)
}
