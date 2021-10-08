package permission

import (
	"context"
	"fmt"
	"log"
	"strings"

	pb "github.com/authzed/authzed-go/proto/authzed/api/v0"
)

func (p Permission) Check(object string, subject string, relation string) {
	o := strings.Split(object, ":")
	s := strings.Split(subject, ":")
	ctx := context.Background()

	resp, err := p.client.Check(ctx, &pb.CheckRequest{
		User: &pb.User{
			UserOneof: &pb.User_Userset{
				Userset: &pb.ObjectAndRelation{
					Namespace: s[0],
					ObjectId:  s[1],
					Relation:  "...",
				}}},
		TestUserset: &pb.ObjectAndRelation{
			Namespace: o[0],
			ObjectId:  o[1],
			Relation:  relation,
		}})

	if err != nil {
		log.Fatalf("failed to check permission: %s", err)
	}

	fmt.Println(resp)

}
