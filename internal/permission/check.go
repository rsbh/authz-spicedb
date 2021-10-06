package permission

import (
	"context"
	"fmt"
	"log"

	pb "github.com/authzed/authzed-go/proto/authzed/api/v0"
)

func (h *PermissionHandler) Check() {

	ctx := context.Background()

	emilia := &pb.User{UserOneof: &pb.User_Userset{Userset: &pb.ObjectAndRelation{
		Namespace: "blog/post",
		ObjectId:  "emilia",
		Relation:  "...",
	}}}

	post1Reader := &pb.ObjectAndRelation{Namespace: "blog/post", ObjectId: "1", Relation: "reader"}

	resp, err := h.client.Check(ctx, &pb.CheckRequest{User: emilia, TestUserset: post1Reader})
	if err != nil {
		log.Fatalf("failed to check permission: %s", err)
	}

	fmt.Println(resp)

}
