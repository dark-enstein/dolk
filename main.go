package main

import (
	"fmt"
	"log"
	"net"
	"os"

	dolk "github.com/dark-enstein/dolk/api/v1"
	"github.com/dark-enstein/dolk/auto"
	"google.golang.org/grpc"
)

func main() {
	parsed := parse(os.Args[1:])
	port, ok := parsed["port"]
	if !ok {
		log.Println("no port specified")
		return
	}
	fmt.Println("grpc server started on port:", port)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Println(err)
		return
	}

	s := grpc.NewServer()
	dolk.RegisterDolkServer(s, &auto.Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalln("server couldn't start", err)
	}
}

func parse(s []string) map[string]string {
	parsed := make(map[string]string)
	fmt.Println(parsed)
	for k, v := range s {
		//fmt.Println(s[k+1])
		switch v {
		case "--port", "-p":
			parsed["port"] = fmt.Sprintf(":%v", s[k+1])
		}
	}
	return parsed
}
