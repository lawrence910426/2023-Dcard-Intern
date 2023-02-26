package main

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	pb "dcard-intern/proto/dcard-intern"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = flag.String("addr", "localhost:50051", "The address to connect to")

func populate_article(t *testing.T) string {
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Error("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewListClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := client.SetList(ctx)
	if err != nil {
		t.Error("client.SetList failed: %v", err)
	}
	for i := 1; i <= 5; i++ {
		article_content := "Article-" + strconv.Itoa(i)
		article := &pb.ListRequest{Article: &article_content}
		if err := stream.Send(article); err != nil {
			t.Error("stream.Send() failed: %v", err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Error("stream.CloseAndRecv failed: %v", err)
	}
	log.Printf("Head identifier: %v", reply)
	return reply
}

type Node struct {
	Article     string
	NextPageKey string
}

func send_request(t *testing.T, url string) Node {
	request, _ := http.NewRequest("GET", url)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	var resp_node Node
	json.Unmarshal(body, &resp_node)
	return resp_node
}

func check_article(t *testing.T, head string) {
	node := send_request("http://localhost/GetHead/" + head)
	for node.NextPageKey != nil {
		node = send_request("http://localhost/GetPage/" + node.NextPageKey)
		if node.Article == nil:
			t.Error("Article content is nil")
			log.Printf("Article content: %v", node.Article)
	}
}

func Test(t *testing.T) {
	flag.Parse()
	head := populate_article(t)
	check_article(t, head)
}
