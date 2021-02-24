package jobs

import (
	"fmt"
	"strings"
	
	"github.com/kylie-sre/wrappers/pkg/aggregator"
	"github.com/kylie-sre/wrappers/pkg/dataset"
	"github.com/kylie-sre/wrappers/pkg/discount"
	"github.com/kylie-sre/wrappers/pkg/storage"
)

type Done int

func (d Done) Error() string {
	return ""
}

type Runner interface {
	Run() error
}

type Option func(*job)

func Storage(storage storage.Creator) Option {
	return func(a *job) {
		a.storage = storage
	}
}

func Dataset(dataset dataset.Dataset) Option {
	return func(a *job) {
		a.dataset = dataset
	}
}

func Aggregator(aggregator aggregator.Aggregator) Option {
	return func(j *job) {
		j.aggregator = aggregator
	}
}

func Discount(discount discount.Discounter) Option {
	return func(j *job) {
		j.discount = discount
	}
}


type job struct {
	storage     storage.Creator
	dataset     dataset.Dataset
	aggregator  aggregator.Aggregator
	discount    discount.Discounter
	queryString string
}

func New(options ...Option) *job {
	j := new(job)
	for _, option := range options {
		option(j)
	}
	return j
}

func (a *job) Run() (err error) {
	var iterator dataset.Iterator
	query := a.dataset.Query(a.queryString)
	if iterator, err = query.Read(); err != nil {
		return fmt.Errorf("error reading query results: %v", err)
	}
	
	var record Record
	switch err = iterator.Next(&record); err.(type) {
	case Done:
		break
	case error:
		return fmt.Errorf("error fetching next result: %v", err)
	default:
		if err = a.aggregator.Add(record.B()); err != nil {
			return
		}
	}

	aggregated := make(Aggregated)
	for key, quantities := range aggregated.S(a.aggregator.Values()) {
		var original, discounted float64
		
		k := Key(key)
		if original, discounted, err = a.discount.Discount(k.UOM(), quantities); err != nil {
			return
		}
		var b []byte
		if b, err = a.aggregator.Get(aggregator.Quantity(original), aggregator.Discounted(discounted)); err != nil {
			return
		}
		if _, err = a.storage.Create(b); err != nil {
			return
		}
	}
	return
}

type Record struct {
}

func (r Record) B() []byte {
	return []byte{}
}

func (r Record) S([]byte) Record {
	return Record{}
}

type Aggregated map[string][]float64

func (r Aggregated) B() []byte {
	return []byte{}
}

func (r Aggregated) S([]byte) Aggregated {
	return Aggregated{}
}

// composed key: ${Customer}::${Subscription}::${UOM}
type key [3]string

func (k key) Customer() string {
	return k[0]
}

func (k key) Subscription() string {
	return k[1]
}

func (k key) UOM() string {
	return k[2]
}

func (k key) Compose() string {
	return fmt.Sprintf("%s::%s::%s", k[0], k[1], k[2])
}

func Key(s ...string) key {
	switch len(s) {
	case 1:
		x := strings.Split(s[0], "::")
		return key{x[0], x[1], x[2]}
	case 3:
		return key{s[0], s[1], s[2]}
	default:
		panic("lulz")
	}
}
