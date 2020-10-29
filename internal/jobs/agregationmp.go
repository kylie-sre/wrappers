package jobs

import (
	"github.com/kylie-sre/wrappers/internal/dataset"
	"github.com/kylie-sre/wrappers/internal/storage"
	"os"
)

type Option func(*aggregationmp)

type aggregationmp struct {
	storage storage.Storage
	dataset dataset.Dataset
	queryString string
}

func Storage(storage storage.Storage) Option {
	return func(a *aggregationmp) {
		a.storage = storage
	}
}

func Dataset(dataset dataset.Dataset) Option {
	return func(a *aggregationmp) {
		a.dataset = dataset
	}
}

func QueryString(q string) Option {
	return func(a *aggregationmp) {
		a.queryString = q
	}
}

func NewAggregationMP(options ...Option) *aggregationmp {
	a := &aggregationmp{}
	for _, option := range options {
		option(a)
	}
	return a
}

func (a *aggregationmp) Run() error {
	var err error
	var iterator dataset.Iterator
	var record *storage.Record
	

	query := a.dataset.Query(a.queryString)
	if iterator, err = query.Read(); err != nil {
		os.Exit(1)
	}
	if err = iterator.Next(record); err != nil {
		os.Exit(2)
	}
	if record, err = a.storage.Create(record); err != nil {
	    os.Exit(3)
	}
	return nil
}

