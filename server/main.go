package main

import (
	"log"
	"net"

	"github.com/1x-eng/struct_laas/server/logservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("oops, failed to listen: %v", err)
	}
	s := grpc.NewServer()
	logservice.Register(s)

	// TODO: Maybe enable this only in dev.
	reflection.Register(s)

	log.Printf("struct_laas gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("oops, failed to serve: %v", err)
	}
}
