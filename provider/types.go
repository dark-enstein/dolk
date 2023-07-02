package provider

import (
	"fmt"

	"github.com/dark-enstein/dolk/config"
	"github.com/dark-enstein/dolk/shape"
)

type Worker struct {
	Provider string
	Config   config.Config
}

func (w Worker) Deploy() (*shape.Shape, error) {
	return &shape.Shape{State: []byte("success")}, fmt.Errorf("nil\n")
}

func Init(p string) *Worker {
	return &Worker{Provider: p}
}
