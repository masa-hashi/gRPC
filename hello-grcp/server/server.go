package main

import (
	"context" 
	"log"
	"net" 

	pb "github.com/masa-hashi/hello-grpc" 
	"google.golang.org/grpc"
	"github.com/golang/protobuf/ptypes/duration"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) { 
	log.Printf("Received: %v", in.Name)	
	st, _ := status.New(codes.Aborted, "aborted").WithDetails(&errdetails.RetryInfo{
			RetryDelay: &duration.Duration{ 
				Seconds: 3,
				Nanos: 0, 
		},
	})
	return nil, st.Err()
}

func main() {	addr := ":50051" 
	lis, err := net.Listen("tcp", addr) 
	if err != nil { 
		log.Fatalf("failed to listen: %v", err) 
	}	
	s := grpc.NewServer() 
	pb.RegisterGreeterServer(s, &server{}) 
	
	log.Printf("gRPC server listening on " + addr)
	if err := s.Serve(lis); err != nil { 
		log.Fatalf("failed to serve: %v", err) 
	}
}