package main

import (
	"github.com/censys-sample/internal/app/kv-service/kvstore"
	"github.com/censys-sample/internal/app/kv-service/server"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	pb "github.com/censys-sample/proto/gen"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("[main] failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	store := kvstore.NewInMemoryStore()
	pb.RegisterKVStoreServer(srv, server.NewKVServer(store))
	log.Printf("[main] gRPC server listening on :50051")
	if err = srv.Serve(lis); err != nil {
		log.Fatalf("[main] failed to serve: %v", err)
	}

	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, os.Interrupt, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	log.Printf("[main] login service is up and running...")

	sig := <-terminate
	log.Printf("[main] registered signal='%s', shutting down...", sig.String())
}
