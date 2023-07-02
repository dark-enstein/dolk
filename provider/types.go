package provider

import (
	"github.com/dark-enstein/dolk/engine"
	"github.com/dark-enstein/dolk/shape"
)

type Worker struct {
	Provider string
	Config   engine.Config
}

func (w Worker) Deploy() (*shape.Shape, error) {
	return nil, nil
}

func Init(p string) *Worker {
	return &Worker{Provider: p}
}
