package permission

import (
	"context"
	"fmt"
	"github.com/authzed/spicedb/pkg/tuple"
	"log"

	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
)

func (p Permission) Check(str string) {
	ctx := context.Background()
	rel := tuple.ParseRel(str)
	resp, err := p.client.CheckPermission(ctx, &pb.CheckPermissionRequest{
		Resource:   rel.Resource,
		Subject:    rel.Subject,
		Permission: rel.Relation,
	})

	if err != nil {
		log.Fatalf("failed to check permission: %s", err)
	}

	fmt.Println(resp.GetPermissionship())

}
