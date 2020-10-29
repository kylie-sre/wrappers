package storage

type Record struct{}

type Storage interface {
	Create(aggUsage *Record) (*Record, error)
	Get(id string) (*Record, error)
	GetFirst(controlID string) (*Record, error)
	MarkSubmission(id string) error
	Delete(controlID string) error
}
