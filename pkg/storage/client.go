package storage

type storage struct {}

func (s storage) Create(in []byte) ([]byte, error) {
	return in, nil
}

func (s storage) Get(id string) ([]byte, error) {
	return []byte(id), nil
}

func (s storage) GetFirst(controlID string) ([]byte, error) {
	return []byte(controlID), nil
}

func (s storage) MarkSubmission(id string) error {
	return nil
}

func (s storage) Delete(controlID string) error {
	return nil
}

func New() *storage {
	return &storage{}
}

type mockstorage struct {}

func (m mockstorage) Create(in []byte) ([]byte, error) {
	return in, nil
}

func (m mockstorage) Get(id string) ([]byte, error) {
	return []byte(id), nil
}

func (m mockstorage) GetFirst(controlID string) ([]byte, error) {
	return []byte(controlID), nil
}

func (m mockstorage) MarkSubmission(id string) error {
	return nil
}

func (m mockstorage) Delete(controlID string) error {
	return nil
}

func NewMock() *mockstorage {
	return &mockstorage{}
}

