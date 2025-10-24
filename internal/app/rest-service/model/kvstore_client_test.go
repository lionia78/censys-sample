package model

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/censys-sample/proto/gen"
)

func dialClient(t *testing.T) (pb.KVStoreClient, *grpc.ClientConn) {
	t.Helper()
	// assuming grpc server is running locally
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Skipf("gRPC server not available: %v", err)
	}
	return pb.NewKVStoreClient(conn), conn
}

func TestPut(t *testing.T) {
	t.Parallel()
	client, conn := dialClient(t)
	if conn != nil {
		defer conn.Close()
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := client.Put(ctx, &pb.PutRequest{Key: "test-put", Value: "value"})
	assert.NoError(t, err)
}

func TestGet_Missing(t *testing.T) {
	t.Parallel()
	client, conn := dialClient(t)
	if conn != nil {
		defer conn.Close()
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.Get(ctx, &pb.GetRequest{Key: "non-existent-key"})
	// If server returns an error it's acceptable to assert on that; otherwise check Found==false
	if assert.NoError(t, err) {
		assert.False(t, resp.Found)
		assert.Equal(t, "", resp.Value)
	}
}

func TestGet_Found(t *testing.T) {
	t.Parallel()
	client, conn := dialClient(t)
	if conn != nil {
		defer conn.Close()
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// ensure key exists
	_, err := client.Put(ctx, &pb.PutRequest{Key: "test-get", Value: "value"})
	assert.NoError(t, err)

	resp, err := client.Get(ctx, &pb.GetRequest{Key: "test-get"})
	if assert.NoError(t, err) {
		assert.True(t, resp.Found)
		assert.Equal(t, "value", resp.Value)
	}
}

func TestDelete_NonExistent(t *testing.T) {
	t.Parallel()
	client, conn := dialClient(t)
	if conn != nil {
		defer conn.Close()
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.Delete(ctx, &pb.DeleteRequest{Key: "nope"})
	if assert.NoError(t, err) {
		// depending on implementation Delete may return Success=false for non-existent key
		assert.False(t, resp.Success)
	}
}

func TestDelete_Existing(t *testing.T) {
	t.Parallel()
	client, conn := dialClient(t)
	if conn != nil {
		defer conn.Close()
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// create key
	_, err := client.Put(ctx, &pb.PutRequest{Key: "test-delete", Value: "value"})
	assert.NoError(t, err)

	// delete it
	resp, err := client.Delete(ctx, &pb.DeleteRequest{Key: "test-delete"})
	if assert.NoError(t, err) {
		assert.True(t, resp.Success)
	}

	// verify gone
	getResp, err := client.Get(ctx, &pb.GetRequest{Key: "test-delete"})
	if assert.NoError(t, err) {
		assert.False(t, getResp.Found)
	}
}
