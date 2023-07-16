package provider

import (
	"github.com/dark-enstein/dolk/internal"
	awspile "github.com/dark-enstein/dolk/provider/aws"
	"github.com/dark-enstein/dolk/shape"
)

type AWSBuilder struct {
	Worker *Worker
}

func (a *AWSBuilder) Deploy() ([]byte, error) {
	w := a.Worker
	pile := &awspile.SageMakerBuilder{
		UUID:     w.UUID,
		Provider: w.Provider,
		Version:  w.Version,
		Name:     w.Name,
		Tags:     w.Tags,
		Options:  w.Options,
		Stack:    w.Stack,
	}
	return pile.Deploy()
}

type Worker struct {
	UUID     string
	Provider string
	Version  string
	Name     string
	Tags     []string
	Options  string
	Stack    *internal.ContextStack
}

func (w *Worker) Deploy() (*shape.Shape, error) {
	trace, log := w.Stack.LogInit()
	var s = &shape.Shape{}
	switch w.Provider {
	case internal.AWS:
		trace.Info().Msgf("provider found: %v", internal.AWS)
		build := &AWSBuilder{
			Worker: w,
		}
		trace.Info().Msgf("builder built: %v. deploying", build)
		resp, err := build.Deploy()
		if err != nil {
			log.Info().Msgf("error while creating the infra: %v", err)
			return nil, err
		}
		s.State = resp
	}
	return s, nil
}
