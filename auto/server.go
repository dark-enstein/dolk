package auto

import (
	"context"
	"log"
	"strings"

	dolk "github.com/dark-enstein/dolk/api/v1"

	"github.com/dark-enstein/dolk/internal"
	"github.com/rs/zerolog"
)

type Server struct {
	Ctx context.Context

	dolk.UnimplementedDolkServer
}

func logInit(s *Server) (*zerolog.Logger, *zerolog.Logger) {
	config := s.Ctx.Value(internal.MainConfig).(*internal.StartUpConfig)
	return config.Logger.Trace(), config.Logger.Err()
}

// Create receives a create request via grpc and processes it accordingly.
// It returns a standard pointer to dolk.CreateResponse, and an error
func (s *Server) Create(ctx context.Context,
	req *dolk.CreateRequest) (resp *dolk.CreateResponse, err error) {
	trace, log := logInit(s)

	trace.Info().Msg("received create request")

	// functional validations
	val, isValid, err := DetentionDirector(ctx, req)
	log.Printf("provider valid: %v", isValid)
	if !isValid || err != nil {
		log.Info().Msgf("provider invalid: %v\n", err)
		return nil, err
	}
	trace.Info().Msgf("provider valid: %v\n", isValid)

	// internal state
	engineRequest := val.NewEngineRequest(&s.Ctx)

	// run engine
	engineResponse := engineRequest.Run()

	trace.Info().Msgf("grpc response: %v", engineResponse)
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
