package auto

import (
	"context"
	"strings"

	dolk "github.com/dark-enstein/dolk/api/v1"
)

type Server struct {
	dolk.UnimplementedDolkServer
}

func (s *Server) Create(ctx context.Context,
	req *dolk.CreateRequest) (resp dolk.CreateResponse, err error) {

	// functional validations
	val, isValid, err := DetentionDirector(ctx, req)
	if !isValid || err != nil {
		return dolk.CreateResponse{}, err
	}

	// internal state
	engineRequest := val.NewEngineRequest()

	// run engine
	engineResponse := engineRequest.Run()

	return dolk.CreateResponse{Created: engineResponse.Created,
			Code: int32(engineResponse.Code), State: engineResponse.Shape.String(),
			AccessConfig: engineResponse.AccessConfig.String(),
			Error:        engineResponse.Error.Error(),
			CreatedTime:  engineResponse.CreatedTime.String()},
		nil
}

func getTagsInCsv(r string) []string {
	return strings.Split(r, ",")
}
