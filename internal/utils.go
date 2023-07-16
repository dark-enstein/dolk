package internal

import (
	"context"

	"github.com/dark-enstein/dolk/dlog"
)

const (
	MainConfig = "_mainConfig"
)

type StartUpConfig struct {
	Logger *dlog.Logger
	Port   string
}

type ContextStack struct {
	Server context.Context
	Client context.Context
}
