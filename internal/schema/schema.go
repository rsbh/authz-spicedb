package schema

import (
	"context"
	"fmt"
	"log"

	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	authzed "github.com/authzed/authzed-go/v1"
	"google.golang.org/grpc"
)

type Schema struct {
	client *authzed.Client
}

func New(endpoint string, opts ...grpc.DialOption) *Schema {
	client, err := authzed.NewClient(endpoint, opts...)

	if err != nil {
		log.Fatalf("unable to initialize client: %s", err)
	}

	return &Schema{
		client: client,
	}
}

func (s *Schema) Load() {

	request := &pb.WriteSchemaRequest{Schema: schema}
	resp, err := s.client.WriteSchema(context.Background(), request)
	if err != nil {
		log.Fatalf("failed to write schema: %s", err)
	}
	fmt.Println("Output", resp)
}

func (s *Schema) Read() {
	request := &pb.ReadSchemaRequest{}
	resp, err := s.client.ReadSchema(context.Background(), request)
	if err != nil {
		log.Fatalf("failed to write schema: %s", err)
	}
	fmt.Println(resp)
}
