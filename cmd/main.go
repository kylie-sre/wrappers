package main

import (
	"os"

	"github.com/kylie-sre/wrappers/internal/dataset"
	"github.com/kylie-sre/wrappers/internal/jobs"
	"github.com/kylie-sre/wrappers/internal/storage"
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

func main() {
	var err error

	args := ParseArgs()
	storageInstance := getStorage(args.Env)
	datasetInstance := getDataset(args.Env)

	options := []jobs.Option{
		jobs.Storage(storageInstance),
		jobs.Dataset(datasetInstance),
	}
	job := jobs.NewAggregationMP(options...)
	if err = job.Run(); err != nil {
		os.Exit(1)
	}
}
