package control

type Control interface {
	HandleJobControlStatus(controlID string, jc []byte, count int) error
	MarkJobControlStatus(controlID, status string, jc []byte) error
	GetOrCreateJobControlInstance(team string) ([]byte, error)
	ErrorWaiting() error
}
