package auto

import (
	"context"
	"fmt"

	dolk "github.com/dark-enstein/dolk/api/v1"
	"github.com/dark-enstein/dolk/config"
	"github.com/dark-enstein/dolk/engine"
	"github.com/dark-enstein/dolk/internal"
)

const (
	ongoing = "SupportInProgress"
	done    = "Supported"
	undone  = "Unsupported"
)

var (
	ErrProviderSupportInprogress = fmt.Errorf("provider support in progress\n")
	ErrProviderUnsupported       = fmt.Errorf("provider unsupported\n")
)

var supportedProviders = map[string]string{
	"AWS": done,
}

type Detention struct {
	Config   config.Config
	Provider string
	UUID     string
	Ctx      *internal.ContextStack
}

func (d *Detention) NewEngineRequest(sCtx *context.Context) *engine.
	EngineRequest {
	d.Ctx.Server = *sCtx
	return &engine.EngineRequest{UUID: d.UUID, Provider: d.Provider,
		Config: d.Config, Ctx: d.Ctx}
}

// Director
func DetentionDirector(ctx context.Context,
	req *dolk.CreateRequest) (*Detention, bool, error) {

	config, isValidConfig, err := validateConfig(req.Config)
	if !isValidConfig {
		return nil, false, err
	}

	uuid, isValid, err := validateUUID(req.UUID)
	if !isValid {
		return nil, false, err
	}
	prov, isSupported, err := validateProvider(req.Provider)
	if !isSupported {
		return nil, false, err
	}
	return &Detention{
		Config:   config,
		Provider: prov,
		UUID:     uuid,
		Ctx:      &internal.ContextStack{Client: ctx},
	}, true, nil
}

func validateConfig(cfg *dolk.Config) (config.Config, bool, error) {
	tags := getTagsInCsv(cfg.Tag)
	return config.Config{Tags: tags}, true, nil
}

func validateUUID(uuid string) (string, bool, error) {
	return uuid, true, nil
}

func validateProvider(prov string) (string, bool, error) {
	status, ok := supportedProviders[prov]
	if !ok {
		return "", false, ErrProviderUnsupported
	}
	switch status {
	case ongoing:
		return "", false, ErrProviderSupportInprogress
	case undone:
		return "", false, ErrProviderUnsupported
	}
	return prov, true, nil
}
