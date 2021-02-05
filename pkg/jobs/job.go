package jobs

import (
	"fmt"
	"github.com/kylie-sre/wrappers/pkg/dataset"
	"github.com/kylie-sre/wrappers/pkg/storage"
)

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

type job struct {
	storage storage.Creator
	dataset dataset.Dataset
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
	var record Record
	
	query := a.dataset.Query(a.queryString)
	if iterator, err = query.Read(); err != nil {
		return fmt.Errorf("error reading query results: %v", err)
	}
	if err = iterator.Next(&record); err != nil {
		return fmt.Errorf("error fetching next result: %v", err)
	}
	if _, err = a.storage.Create(RBytes{}); err != nil {
	    return fmt.Errorf("error storing result: %v", err)
	}
	return
}

type RBytes []byte

type Record struct {

}

func (r Record) B() RBytes {
	return []byte{}
}