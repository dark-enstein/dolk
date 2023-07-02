package main

import (
	"fmt"
	"net"
	"os"

	dolk "github.com/dark-enstein/dolk/api/v1"
	"github.com/dark-enstein/dolk/auto"
	"github.com/dark-enstein/dolk/dlog"
	"google.golang.org/grpc"
)

var (
	ErrNoPortSpecified  = fmt.Errorf("no port specified")
	ErrListenerNotStart = fmt.Errorf("listener couldn't start")
	ErrServerNotStart   = fmt.Errorf("server couldn't start")
)

func main() {
	log := dlog.NewLogger().Err
	trace := dlog.NewLogger().Trace

	parsed := parse(os.Args[1:])
	port, ok := parsed["port"]
	if !ok {
		log.Fatal().Str("scope", "entrypoint").Err(ErrNoPortSpecified).Send()
		return
	}
	trace.Info().Msgf("grpc server started on port%v\n", port)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal().Str("scope", "entrypoint").Err(err).Msg(ErrListenerNotStart.Error())
		return
	}

	s := grpc.NewServer()
	dolk.RegisterDolkServer(s, &auto.Server{Logger: dlog.NewLogger()})
	if err := s.Serve(listener); err != nil {
		log.Fatal().Str("scope", "entrypoint").Err(err).Msg(ErrServerNotStart.Error())
	}
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
