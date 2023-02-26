package main

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"testing"
	"time"

	pb "dcard-intern/proto/proto_gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = flag.String("addr", "localhost:50051", "The address to connect to")

func populate_article(t *testing.T) string {
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Error("did not connect")
	}
	defer conn.Close()
	client := pb.NewListClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := client.SetList(ctx)
	if err != nil {
		t.Error("client.SetList failed")
	}
	for i := 1; i <= 5; i++ {
		article_content := "Article-" + strconv.Itoa(i)
		article := &pb.ListRequest{Article: &article_content}
		if err := stream.Send(article); err != nil {
			t.Error("stream.Send() failed")
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Error("stream.CloseAndRecv failed")
	}
	log.Printf("Head identifier: %s", reply)
	return reply.GetHead()
}

type Node struct {
	Article     string
	NextPageKey string
}

func send_request(t *testing.T, url string) Node {
	response, err := http.Get(url)
	if err != nil {
		t.Error("Unable to request server")
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	var resp_node Node
	json.Unmarshal(body, &resp_node)
	return resp_node
}

func check_article(t *testing.T, head string) {
	node := send_request(t, "http://localhost/GetHead/"+head)
	for node.NextPageKey != "" {
		node = send_request(t, "http://localhost/GetPage/"+node.NextPageKey)
		if node.Article == "" {
			t.Error("Article content is empty")
			log.Printf("Article content: %s", node.NextPageKey)
		}
	}
}

func Test(t *testing.T) {
	flag.Parse()
	head := populate_article(t)
	check_article(t, head)
}
