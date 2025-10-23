package server

import (
	context "context"
	"github.com/censys-sample/internal/app/kv-service/kvstore"
	"log"

	pb "github.com/censys-sample/proto/gen"
)

// KVServer implements the gRPC KVStore service.
type KVServer struct {
	pb.UnimplementedKVStoreServer
	store kvstore.Store
}

func NewKVServer(s kvstore.Store) *KVServer {
	return &KVServer{store: s}
}

func (s *KVServer) Put(ctx context.Context, req *pb.PutRequest) (*pb.PutResponse, error) {
	log.Printf("Put key=%s", req.Key)
	s.store.Put(req.Key, req.Value)
	return &pb.PutResponse{
		Success: true,
	}, nil
}

func (s *KVServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	log.Printf("Get key=%s", req.Key)
	v, ok := s.store.Get(req.Key)
	return &pb.GetResponse{
		Found: ok,
		Value: v,
	}, nil
}

func (s *KVServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	log.Printf("Delete key=%s", req.Key)
	ok := s.store.Delete(req.Key)
	return &pb.DeleteResponse{
		Success: ok,
	}, nil
}
