package client

import (
	"context"
	"log"
	"os"

	pb "grpc-test/routeguide"

	"google.golang.org/grpc"
)

func GetFeature() {
	// serverAddr := "127.0.0.1:8080"
	serverAddr := os.Getenv("SERVER_ADDR")
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewRouteGuideClient(conn)

	feature, err := client.GetFeature(context.Background(), &pb.Point{
		Latitude:  409146138,
		Longitude: -746188906,
	})
	if err != nil {
		panic(err)
	}
	log.Println(feature)
}
