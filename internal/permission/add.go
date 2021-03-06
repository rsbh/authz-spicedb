package permission

import (
	"context"
	// "fmt"
	"log"

	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/spicedb/pkg/tuple"
)

func (p *Permission) Add(str string) {
	request := &pb.WriteRelationshipsRequest{Updates: []*pb.RelationshipUpdate{
		{
			Operation:    pb.RelationshipUpdate_OPERATION_CREATE,
			Relationship: tuple.ParseRel(str),
		},
	}}

	_, err := p.client.WriteRelationships(context.Background(), request)

	if err != nil {
		log.Fatalf("failed to check permission: %s", err)
	}

	// fmt.Println(resp)
}
