package schema

import (
	"context"
	"fmt"
	"log"

	pb "github.com/authzed/authzed-go/proto/authzed/api/v1alpha1"
	authzed "github.com/authzed/authzed-go/v1alpha1"
	"google.golang.org/grpc"
)

func Read() {
	client, err := authzed.NewClient(
		"localhost:50051",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("unable to initialize client: %s", err)
	}

	request := &pb.WriteSchemaRequest{Schema: schema}
	resp, err := client.WriteSchema(context.Background(), request)
	if err != nil {
		log.Fatalf("failed to write schema: %s", err)
	}
	fmt.Println("Output", resp)
}
