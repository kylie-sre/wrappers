package storage

type Creator interface {
	Create(in []byte) ([]byte, error)
}

type Getter interface {
	Get(id string) ([]byte, error)
	GetFirst(controlID string) ([]byte, error)
}

type Marker interface {
	MarkSubmission(id string) error
}

type Deleter interface {
	Delete(controlID string) error
}

type Storage interface {
	Creator
	Getter
	Marker
	Deleter
}
