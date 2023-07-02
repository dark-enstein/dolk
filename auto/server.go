package auto

import (
	"context"
	"log"
	"strings"

	dolk "github.com/dark-enstein/dolk/api/v1"
	"github.com/dark-enstein/dolk/dlog"
)

type Server struct {
	Logger *dlog.Logger
	dolk.UnimplementedDolkServer
}

func (s *Server) Create(ctx context.Context,
	req *dolk.CreateRequest) (resp *dolk.CreateResponse, err error) {
	log.Println("received create request")

	// functional validations
	val, isValid, err := DetentionDirector(ctx, req)
	log.Printf("provider valid: %v", isValid)
	if !isValid || err != nil {
		return &dolk.CreateResponse{}, err
	}

	// internal state
	engineRequest := val.NewEngineRequest()

	// run engine
	engineResponse := engineRequest.Run()

	return &dolk.CreateResponse{Created: engineResponse.Created,
			Code: int32(engineResponse.Code), State: engineResponse.Shape.String(),
			AccessConfig: engineResponse.AccessConfig.String(),
			Error:        engineResponse.Error.Error(),
			CreatedTime:  engineResponse.CreatedTime.String()},
		nil
}

func getTagsInCsv(r string) []string {
	return strings.Split(r, ",")
}
