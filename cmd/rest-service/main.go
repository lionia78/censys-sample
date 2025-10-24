package main

import (
	"log"
	"os"

	"github.com/censys-sample/internal/app/rest-service/controller"
	"github.com/censys-sample/internal/app/rest-service/model"
	"github.com/censys-sample/internal/app/rest-service/transport"
)

func main() {
	// Getting gRPC address from env variable
	grpcAddr := os.Getenv("KV_SERVICE_ADDR")
	if grpcAddr == "" {
		// Fallback - using default if not set by environment variable
		grpcAddr = "localhost:50051"
	}

	httpAddr := ":8080"

	// 1. Initialize the Model (gRPC Client)
	kvModel, conn, err := model.NewKVStoreClient(grpcAddr)
	if err != nil {
		log.Fatalf("[main] failed to initialize KV model (gRPC client): %v", err)
	}
	defer conn.Close()

	// 2. Initialize the Controller with the Model
	kvController := &controller.KVController{
		KVService: kvModel,
	}

	// 3. Start the HTTP Server
	transport.StartServer(kvController, httpAddr)
}
