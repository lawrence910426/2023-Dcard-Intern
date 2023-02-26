package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	pb "dcard-intern/proto/dcard-intern"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedListServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SetList(stream pb.List_SetListServer) error {
	for {
		article, err := stream.Recv()
		if err != io.EOF {
			head := "qwer"
			return stream.SendAndClose(&pb.ListResponse{Head: &head})
		}
		if err != nil {
			return err
		}
		log.Printf("Received: %v", article.GetArticle())
	}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterListServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
