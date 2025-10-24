package model

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/censys-sample/proto/gen"
)

// KVModel defines the interface for key-value operations.
type KVModel interface {
	Put(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, bool, error)
	Delete(ctx context.Context, key string) (bool, error)
}

// KVStoreClient implements the KVModel interface.
type KVStoreClient struct {
	client pb.KVStoreClient
}

// NewKVStoreClient establishes the gRPC connection and returns a new KVModel instance.
func NewKVStoreClient(addr string) (KVModel, *grpc.ClientConn, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	client := pb.NewKVStoreClient(conn)
	return &KVStoreClient{client: client}, conn, nil
}

func (c *KVStoreClient) Put(ctx context.Context, key, value string) error {
	_, err := c.client.Put(ctx, &pb.PutRequest{
		Key:   key,
		Value: value,
	})
	return err
}

func (c *KVStoreClient) Get(ctx context.Context, key string) (string, bool, error) {
	res, err := c.client.Get(ctx, &pb.GetRequest{Key: key})
	if err != nil {
		return "", false, err
	}
	return res.Value, res.Found, nil
}

func (c *KVStoreClient) Delete(ctx context.Context, key string) (bool, error) {
	res, err := c.client.Delete(ctx, &pb.DeleteRequest{Key: key})
	if err != nil {
		return false, err
	}
	return res.Success, err
}
