package server

import (
	"context"
	"fmt"
	pb "grpc-test/routeguide"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func StartServer() {
	portStr := os.Getenv("GO_PORT")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", portStr))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterRouteGuideServer(grpcServer, &routeGuideServer{})
	// determine whether to use TLS
	log.Printf("Listening on %s\n", portStr)
	grpcServer.Serve(lis)
}

type routeGuideServer struct {
}

func (s *routeGuideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	log.Printf("Received request: %v\n", point)
	return &pb.Feature{
		Name:     "unknown",
		Location: point,
	}, nil
}
