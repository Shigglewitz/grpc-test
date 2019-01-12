package client

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	pb "grpc-test/routeguide"

	"google.golang.org/grpc"
)

var conn *grpc.ClientConn
var client pb.RouteGuideClient

func handler(w http.ResponseWriter, r *http.Request) {
	feature, err := client.GetFeature(context.Background(), &pb.Point{
		Latitude:  409146138,
		Longitude: -746188906,
	})
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%v", feature)
}

func main() {
	http.HandleFunc("/", handler)
	portStr := os.Getenv("GO_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", portStr), nil))
}

func StartClientServer() {
	serverAddr := os.Getenv("SERVER_ADDR")
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client = pb.NewRouteGuideClient(conn)

	main()
}

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
