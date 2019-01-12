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
	myClient := client
	if "reuse" != os.Getenv("CONN_POLICY") {
		serverAddr := os.Getenv("SERVER_ADDR")
		myConn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		defer myConn.Close()

		myClient = pb.NewRouteGuideClient(myConn)
	}

	feature, err := myClient.GetFeature(context.Background(), &pb.Point{
		Latitude:  409146138,
		Longitude: -746188906,
	})
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%v\n", feature)
}

func main() {
	http.HandleFunc("/", handler)
	portStr := os.Getenv("GO_PORT")
	log.Printf("Listening on port %s\n", portStr)
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
