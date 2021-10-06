package permission

import (
	"context"
	"fmt"
	"log"

	pb "github.com/authzed/authzed-go/proto/authzed/api/v0"
)

func (h *PermissionHandler) Add() {

	request := &pb.WriteRequest{Updates: []*pb.RelationTupleUpdate{
		{ // Emilia is a Writer on Post 1
			Operation: pb.RelationTupleUpdate_CREATE,
			Tuple: &pb.RelationTuple{
				User: &pb.User{UserOneof: &pb.User_Userset{Userset: &pb.ObjectAndRelation{
					Namespace: "blog/user",
					ObjectId:  "emilia",
					Relation:  "...",
				}}},
				ObjectAndRelation: &pb.ObjectAndRelation{
					Namespace: "blog/post",
					ObjectId:  "1",
					Relation:  "writer",
				},
			},
		},
	}}

	resp, err := h.client.Write(context.Background(), request)

	if err != nil {
		log.Fatalf("failed to check permission: %s", err)
	}

	fmt.Println(resp)
}
