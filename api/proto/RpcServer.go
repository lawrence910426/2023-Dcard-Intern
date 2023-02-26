package RpcServer

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	pb "dcard-intern/proto/proto_gen"
	"dcard-intern/services"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedListServer
}

func (s *server) SetList(stream pb.List_SetListServer) error {
	head := uuid.New().String()
	previous := head
	for {
		content, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.ListResponse{Head: &head})
		}
		if err != nil {
			return err
		}
		article := content.GetArticle()
		log.Printf("Received: %v", article)

		new_id := uuid.New().String()
		services.Set(previous, new_id)
		services.Set(new_id, article)
		previous = new_id
	}
}

func StartServer() {
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
