package storage

type storage struct {}

func New() *storage {
	return &storage{}
}

func (s *storage) Create(aggUsage *Record) (*Record, error) {
	panic("implement me")
}

func (s *storage) Get(id string) (*Record, error) {
	panic("implement me")
}

func (s *storage) GetFirst(controlID string) (*Record, error) {
	panic("implement me")
}

func (s *storage) MarkSubmission(id string) error {
	panic("implement me")
}

func (s *storage) Delete(controlID string) error {
	panic("implement me")
}

type mockstorage struct {}

func NewMock() *mockstorage {
	return &mockstorage{}
}

func (m *mockstorage) Create(aggUsage *Record) (*Record, error) {
	panic("implement me")
}

func (m *mockstorage) Get(id string) (*Record, error) {
	panic("implement me")
}

func (m *mockstorage) GetFirst(controlID string) (*Record, error) {
	panic("implement me")
}

func (m *mockstorage) MarkSubmission(id string) error {
	panic("implement me")
}

func (m *mockstorage) Delete(controlID string) error {
	panic("implement me")
}

