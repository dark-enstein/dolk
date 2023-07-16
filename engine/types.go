package engine

import (
	"time"

	"github.com/dark-enstein/dolk/config"
	"github.com/dark-enstein/dolk/cred"
	"github.com/dark-enstein/dolk/internal"
	"github.com/dark-enstein/dolk/provider"
	"github.com/dark-enstein/dolk/shape"
)

type EngineRequest struct {
	UUID     string
	Provider string
	Config   *config.Config
	Ctx      *internal.ContextStack
}

func (er *EngineRequest) WithUUID(u string) *EngineRequest {
	er.UUID = u
	return er
}

func (er *EngineRequest) WithProvider(prov string) *EngineRequest {

	er.Provider = prov
	return er
}

func (er *EngineRequest) WithConfig(c *config.Config) *EngineRequest {
	er.Config = c
	return er
}

func (er *EngineRequest) Run() EngineResponse {
	trace, log := er.Ctx.LogInit()
	trace.Info().Msg("starting worker")
	worker := er.init()
	trace.Info().Msgf("worker: %v", worker)

	trace.Info().Msg("deploying")
	shape, err := worker.Deploy()
	if err != nil {
		log.Error().Msgf("encountered error while worker deploying: %v", err)
	}
	trace.Info().Msgf("deployed\nshape: %v", shape)
	return EngineResponse{AccessConfig: &cred.
		Config{Config: []byte("access your new deployment here")},
		Shape:       shape,
		Error:       "",
		CreatedTime: time.Now()}
}

func (er *EngineRequest) init() *provider.Worker {
	return &provider.Worker{UUID: er.UUID, Provider: er.Provider,
		Version: er.Config.Version, Tags: er.Config.Tags,
		Name: er.Config.Name, Options: er.Config.Directives,
		Stack: er.Ctx}
}

type EngineResponse struct {
	Code         int
	Created      bool
	Error        string
	Shape        *shape.Shape // have its own package
	AccessConfig *cred.Config // its own package
	CreatedTime  time.Time
}
