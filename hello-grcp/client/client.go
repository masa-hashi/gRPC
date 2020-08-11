package main

import (
	"context"
	"log"
	"os"

	pb "github.com/masa-hashi/hello-grpc"
	"google.golang.org/grpc"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/credentials"
)

func main () {
	addr := "localhost:50051"
	cred, err := credentials.NewClientTLSFromFile("server.crt", "")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(cred))
	if err != nil {
		log.Fatalf("did not connct: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := os.Args[1]

	ctx := context.Background()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})

	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			log.Printf("gRPC Error (message: %s)", s.Message())
			for _, d := range s.Details() {
				switch info := d.(type) {
					case *errdetails.RetryInfo:
						log.Printf(" RetryInfo: %v", info)
				}
			}
			os.Exit(1)
		} else {
			log.Fatalf("could not greet: %v", err)
		}
	}
	log.Printf("Greeting: %s" , r.Message)
}