package engine

import (
	"time"

	"github.com/dark-enstein/dolk/config"
	"github.com/dark-enstein/dolk/cred"
	"github.com/dark-enstein/dolk/provider"
	"github.com/dark-enstein/dolk/shape"
)

type EngineRequest struct {
	UUID     string
	Provider string
	Config   config.Config
}

func (er *EngineRequest) WithUUID(u string) *EngineRequest {
	er.UUID = u
	return er
}

func (er *EngineRequest) WithProvider(prov string) *EngineRequest {

	er.Provider = prov
	return er
}

func (er *EngineRequest) WithConfig(c config.Config) *EngineRequest {
	er.Config = c
	return er
}

func (er *EngineRequest) Run() EngineResponse {
	worker := provider.Init(er.Provider)

	shape, err := worker.Deploy()
	return EngineResponse{AccessConfig: &cred.
	Config{Config: []byte("access your new deployment here"), Metadata: "no metadata"},
		Shape: shape,
		Error: err}
}

type EngineResponse struct {
	Code         int
	Created      bool
	Error        error
	Shape        *shape.Shape // have its own package
	AccessConfig *cred.Config // its own package
	CreatedTime  time.Time
}
