package shape

var (
	ErrMarshallingState = "error with marshalling state"
)

type Shape struct {
	state []byte
}

func (s *Shape) String() string {
	return string(s.state)
}
