package main

import (
	"github.com/censys-sample/internal/app/kv-service/kvstore"
	"github.com/censys-sample/internal/app/kv-service/server"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/censys-sample/proto/gen"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	store := kvstore.NewInMemoryStore()
	pb.RegisterKVStoreServer(srv, server.NewKVServer(store))
	log.Printf("gRPC server listening on :50051")
	if err = srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
