package internal

import (
	"context"

	"github.com/rs/zerolog"
)

const (
	MainConfig = "_mainConfig"
)

type StartUpConfig struct {
	Logger *Logger
	Port   string
}

type ContextStack struct {
	Server context.Context
	Client context.Context
}

func NewContextStack(server context.Context, client context.Context) *ContextStack {
	return &ContextStack{
		Server: server,
		Client: client,
	}
}

func (cs *ContextStack) LogInit() (*zerolog.Logger, *zerolog.Logger) {
	config := cs.Server.Value(MainConfig).(*StartUpConfig)
	return config.Logger.Trace(), config.Logger.Err()
}
