package shape

import "log"

var (
	ErrMarshallingState = "error with marshalling state"
)

type Shape struct {
	State []byte
}

func (s *Shape) String() string {
	if len(s.State) == 0 {
		log.Println(s.State)
		return "successful state {}"
	}
	log.Println(s.State)
	return string(s.State)
}
