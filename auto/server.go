package auto

import (
	"context"
	"strings"

	dolk "github.com/dark-enstein/dolk/api/v1"
	"github.com/dark-enstein/dolk/internal"
)

const ( // context-specific K/Vs
	ContextClientRequest = "client request"
)

type Server struct {
	Ctx context.Context

	dolk.UnimplementedDolkServer
}

// Create receives a create request via grpc and processes it accordingly.
// It returns a standard pointer to dolk.CreateResponse, and an error
func (s *Server) Create(ctx context.Context,
	req *dolk.CreateRequest) (resp *dolk.CreateResponse, err error) {
	stack := internal.NewContextStack(s.Ctx, ctx)
	trace, log := stack.LogInit()

	trace.Info().Msg("received create request")

	trace.Info().Msg("saved create request into client context")

	// functional validations
	val, isValid, err := DetentionDirector(stack, req)
	log.Printf("provider valid: %v", isValid)
	if !isValid || err != nil {
		log.Info().Msgf("provider invalid: %v\n", err)
		return nil, err
	}
	trace.Info().Msgf("provider valid: %v\n", isValid)

	// internal state
	engineRequest := val.NewEngineRequest()

	// run engine
	trace.Info().Msgf("runnning engine request: %v", engineRequest)
	engineResponse := engineRequest.Run()

	trace.Info().Msgf("grpc response: %v", engineResponse)
	return &dolk.CreateResponse{Created: engineResponse.Created,
			Code: int32(engineResponse.Code), State: engineResponse.Shape.String(),
			AccessConfig: engineResponse.AccessConfig.String(),
			Error:        engineResponse.Error,
			CreatedTime:  engineResponse.CreatedTime.String()},
		nil
}

func getTagsInCsv(r string) []string {
	return strings.Split(r, ",")
}
