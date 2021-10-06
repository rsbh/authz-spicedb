package schema

import (
	"context"
	"fmt"
	"log"

	pb "github.com/authzed/authzed-go/proto/authzed/api/v1alpha1"
)

func (s *Schema) Load() {

	request := &pb.WriteSchemaRequest{Schema: schema}
	resp, err := s.client.WriteSchema(context.Background(), request)
	if err != nil {
		log.Fatalf("failed to write schema: %s", err)
	}
	fmt.Println("Output", resp)
}
