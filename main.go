package main

import (
	"context"
	"fmt"
	"log"

	spicedb "github.com/authzed/authzed-go/proto/authzed/api/v0"
	"google.golang.org/grpc"
)

func main() {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()
	add_config(cc)
	check_access(cc)
}

func add_config(cc *grpc.ClientConn) {
	client := spicedb.NewNamespaceServiceClient(cc)
	request := &spicedb.WriteConfigRequest{
		Configs: []*spicedb.NamespaceDefinition{
			{
				Name: "test",
				Relation: []*spicedb.Relation{
					{
						Name: "user",
					},
				},
			},
		},
	}
	resp, err := client.WriteConfig(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

func check_access(cc *grpc.ClientConn) {
	client := spicedb.NewACLServiceClient(cc)
	request := &spicedb.CheckRequest{
		TestUserset: &spicedb.ObjectAndRelation{
			ObjectId:  "abcd",
			Relation:  "owner",
			Namespace: "default",
		},
		User: &spicedb.User{
			UserOneof: &spicedb.User_Userset{
				Userset: &spicedb.ObjectAndRelation{
					ObjectId:  "abcd",
					Relation:  "owner",
					Namespace: "default",
				},
			},
		},
	}
	resp, err := client.Check(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
