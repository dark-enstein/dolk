package auto

import (
	"context"
	"strings"

	dolk "github.com/dark-enstein/dolk/api/v1"
	"github.com/dark-enstein/dolk/dlog"
	"github.com/rs/zerolog"
)

type Server struct {
	Genke *dlog.Logger
	dolk.UnimplementedDolkServer
}

func logInit(s *Server) (zerolog.Logger, zerolog.Logger) {
	return s.Genke.Trace, s.Genke.Err
}

func (s *Server) Create(ctx context.Context,
	req *dolk.CreateRequest) (resp *dolk.CreateResponse, err error) {
	trace, log := logInit(s)

	trace.Info().Msg("received create request")

	// functional validations
	val, isValid, err := DetentionDirector(ctx, req)
	if !isValid || err != nil {
		log.Error().Msgf("provider invalid: %v\n", err)
		return &dolk.CreateResponse{}, err
	}
	trace.Info().Msgf("provider valid: %v\n", isValid)

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
