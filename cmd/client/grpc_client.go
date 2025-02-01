package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/samekigor/quill-cli/proto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	SocketPath string
	GrpcClient *GRPCClient
)

type GRPCClient struct {
	conn *grpc.ClientConn
	Auth auth.AuthClient // Client form auth.proto
}

func NewGRPCClient(socketPath string) (*GRPCClient, error) {
	conn, err := grpc.NewClient(
		fmt.Sprintf("unix://%s", socketPath),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to daemon: %w", err)
	}

	return &GRPCClient{
		conn: conn,
		Auth: auth.NewAuthClient(conn),
	}, nil
}

func (c *GRPCClient) Close() {
	if err := c.conn.Close(); err != nil {
		log.Printf("Error closing gRPC connection: %v", err)
	}
}

func WithTimeout(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, timeout)
}
