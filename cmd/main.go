package main

import (
	"fmt"
	"os"
	
	"github.com/kylie-sre/wrappers/pkg/aggregator"
	"github.com/kylie-sre/wrappers/pkg/dataset"
	"github.com/kylie-sre/wrappers/pkg/discount"
	"github.com/kylie-sre/wrappers/pkg/jobs"
	"github.com/kylie-sre/wrappers/pkg/storage"
)

type Args struct {
	QueryString, Env string
}

func ParseArgs() Args {
	return Args{QueryString: "foo", Env: "test"}
}

func getDataset(env string) dataset.Dataset {
	switch env {
	case "prd":
		return dataset.New()
	default:
		return dataset.NewMock()
	}
}

func getStorage(env string) storage.Storage {
	switch env {
	case "prd":
		return storage.New()
	default:
		return storage.NewMock()
	}
}

func getAggregator(env string) aggregator.Aggregator {
	return aggregator.New()
}

func getDiscounter(env string) discount.Discounter {
	return discount.New()
}

func main() {
	var err error

	args := ParseArgs()
	storageInstance := getStorage(args.Env)
	datasetInstance := getDataset(args.Env)
	discounterInstance := getDiscounter(args.Env)
	aggregatorInstance := getAggregator(args.Env)

	options := []jobs.Option{
		jobs.Storage(storageInstance),
		jobs.Dataset(datasetInstance),
		jobs.Discount(discounterInstance),
		jobs.Aggregator(aggregatorInstance),
	}
	job := jobs.New(options...)
	if err = job.Run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
