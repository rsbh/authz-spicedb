package permission

import (
	"context"
	"fmt"
	"log"
	"strings"

	pb "github.com/authzed/authzed-go/proto/authzed/api/v0"
)

func (p *Permission) Add(object string, subject string, relation string) {
	o := strings.Split(object, ":")
	s := strings.Split(subject, ":")
	request := &pb.WriteRequest{Updates: []*pb.RelationTupleUpdate{
		{
			Operation: pb.RelationTupleUpdate_CREATE,
			Tuple: &pb.RelationTuple{
				User: &pb.User{UserOneof: &pb.User_Userset{Userset: &pb.ObjectAndRelation{
					Namespace: s[0],
					ObjectId:  s[1],
					Relation:  "...",
				}}},
				ObjectAndRelation: &pb.ObjectAndRelation{
					Namespace: o[0],
					ObjectId:  o[1],
					Relation:  relation,
				},
			},
		},
	}}

	resp, err := p.client.Write(context.Background(), request)

	if err != nil {
		log.Fatalf("failed to check permission: %s", err)
	}

	fmt.Println(resp)
}
