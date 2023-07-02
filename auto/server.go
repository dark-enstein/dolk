package auto

import (
	"context"
	dolk "dolk/api/v1"
	"fmt"
)

type Server struct {
	dolk.UnimplementedDolkServer
}

func (s *Server) Create(ctx context.Context,
	request *dolk.CreateRequest) (resp *dolk.CreateResponse, err error) {

	return &dolk.CreateResponse{Created: true, Error: fmt.Sprintf("no error")},
		nil
}
