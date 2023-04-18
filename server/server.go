package server

import (
	"context"
	"dapr-pluggable-state-store-demo/components"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type DemoStateStoreServer struct {
	SockPath string
	components.UnimplementedStateStoreServer
}

// Initializes the state store component with the given metadata.
func (s *DemoStateStoreServer) Init(_ context.Context, _ *components.InitRequest) (*components.InitResponse, error) {
	// empty resonponse
	return &components.InitResponse{}, nil
}

// Returns a list of implemented state store features.
func (s *DemoStateStoreServer) Features(_ context.Context, _ *components.FeaturesRequest) (*components.FeaturesResponse, error) {
	rsp := &components.FeaturesResponse{
		Features: make([]string, 0),
	}
	return rsp, nil
}

// Deletes the specified key from the state store.
func (s *DemoStateStoreServer) Delete(_ context.Context, _ *components.DeleteRequest) (*components.DeleteResponse, error) {
	return &components.DeleteResponse{}, nil
}

// Get data from the given key.
func (s *DemoStateStoreServer) Get(_ context.Context, req *components.GetRequest) (*components.GetResponse, error) {
	key := req.Key
	return &components.GetResponse{
		Data:        []byte("key=" + key),
		ContentType: "string",
	}, nil
}

// Sets the value of the specified key.
func (s *DemoStateStoreServer) Set(_ context.Context, _ *components.SetRequest) (*components.SetResponse, error) {
	return &components.SetResponse{}, nil
}

// Ping the state store. Used for liveness porpuses.
func (s *DemoStateStoreServer) Ping(_ context.Context, _ *components.PingRequest) (*components.PingResponse, error) {
	return &components.PingResponse{}, nil
}

// Deletes many keys at once.
func (s *DemoStateStoreServer) BulkDelete(_ context.Context, _ *components.BulkDeleteRequest) (*components.BulkDeleteResponse, error) {
	return &components.BulkDeleteResponse{}, nil
}

// Retrieves many keys at once.
func (s *DemoStateStoreServer) BulkGet(_ context.Context, _ *components.BulkGetRequest) (*components.BulkGetResponse, error) {
	return &components.BulkGetResponse{}, nil
}

// Set the value of many keys at once.
func (s *DemoStateStoreServer) BulkSet(_ context.Context, _ *components.BulkSetRequest) (*components.BulkSetResponse, error) {
	return &components.BulkSetResponse{}, nil
}

// launch the server
func (s *DemoStateStoreServer) Serve() error {
	lis, err := net.Listen("unix", s.SockPath)
	if err != nil {
		return err
	}
	// implement Call using biz handler
	grpcServer := grpc.NewServer()
	components.RegisterStateStoreServer(grpcServer, s)
	reflection.Register(grpcServer)
	// bind and restful
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	log.Printf("serve started succeed.")
	return nil
}
