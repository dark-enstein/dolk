package main

import (
	"context"
	"fmt"
	"net"
	"os"

	dolk "github.com/dark-enstein/dolk/api/v1"
	"github.com/dark-enstein/dolk/auto"
	"github.com/dark-enstein/dolk/internal"
	"google.golang.org/grpc"
)

var (
	ErrNoPortSpecified  = fmt.Errorf("no port specified")
	ErrListenerNotStart = fmt.Errorf("listener couldn't start")
	ErrServerNotStart   = fmt.Errorf("server couldn't start")
)

type StartUpInit struct {
	log internal.Logger
	ctx context.Context
}

func main() {
	i := StartUpInit{}
	i.ctx = context.Background()
	genke := internal.NewLogger(false)
	log := genke.Err()
	trace := genke.Trace()

	// load server config

	parsed := parse(os.Args[1:])
	port, ok := parsed["port"]
	if !ok {
		log.Fatal().Str("scope", "entrypoint").Err(ErrNoPortSpecified).Send()
		return
	}
	trace.Info().Msgf("initiating connection to port %v", port)

	sConfig := &internal.StartUpConfig{internal.NewLogger(false), port}
	trace.Info().Msg("populating startup config")
	i.ctx = context.WithValue(i.ctx, internal.MainConfig, sConfig)
	trace.Info().Msg("loading context")

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal().Str("scope", "entrypoint").Err(err).Msg(
			ErrListenerNotStart.Error())
		return
	}
	trace.Info().Msg("listener started")

	s := grpc.NewServer()
	trace.Info().Msg("server initialized")
	dolk.RegisterDolkServer(s, &auto.Server{Ctx: i.ctx})
	trace.Info().Msg("registered dolk server")

	trace.Info().Msgf("server started on port %v", port)

	if err := s.Serve(listener); err != nil {
		log.Fatal().Str("scope", "entrypoint").Err(err).Msg(
			ErrServerNotStart.Error())
		return
	}
}

// RetrieveFromCtx retrieves the value to the key stored in context
func (i *StartUpInit) RetrieveFromCtx(key string) *internal.StartUpConfig {
	if key == "" {
		return &internal.StartUpConfig{}

	}
	return i.ctx.Value(key).(*internal.StartUpConfig)
}

func parse(s []string) map[string]string {
	parsed := make(map[string]string)
	for k, v := range s {
		switch v {
		case "--port", "-p":
			parsed["port"] = fmt.Sprintf(":%v", s[k+1])
		}
	}
	return parsed
}
