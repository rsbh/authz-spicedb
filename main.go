// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	spicedb "github.com/authzed/authzed-go/proto/authzed/api/v0"
// 	"google.golang.org/grpc"
// )

// func main() {
// 	opts := grpc.WithInsecure()
// 	cc, err := grpc.Dial("localhost:50051", opts)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer cc.Close()
// 	add_config(cc)
// 	check_access(cc)
// 	read_config(cc)
// }

// func add_config(cc *grpc.ClientConn) {
// 	client := spicedb.NewNamespaceServiceClient(cc)
// 	request := &spicedb.WriteConfigRequest{
// 		Configs: []*spicedb.NamespaceDefinition{
// 			{
// 				Name: "default",
// 				Relation: []*spicedb.Relation{
// 					{
// 						Name: "owner",
// 					},
// 				},
// 			},
// 		},
// 	}
// 	resp, err := client.WriteConfig(context.Background(), request)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(resp)
// }

// func read_config(cc *grpc.ClientConn) {
// 	client := spicedb.NewNamespaceServiceClient(cc)
// 	request := &spicedb.ReadConfigRequest{
// 		Namespace: "default",
// 	}
// 	resp, err := client.ReadConfig(context.Background(), request)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(resp)
// }

// func check_access(cc *grpc.ClientConn) {
// 	client := spicedb.NewACLServiceClient(cc)
// 	request := &spicedb.CheckRequest{
// 		TestUserset: &spicedb.ObjectAndRelation{
// 			ObjectId:  "abcd",
// 			Relation:  "owner",
// 			Namespace: "default",
// 		},
// 		User: &spicedb.User{
// 			UserOneof: &spicedb.User_Userset{
// 				Userset: &spicedb.ObjectAndRelation{
// 					ObjectId:  "abcd",
// 					Relation:  "owner",
// 					Namespace: "default",
// 				},
// 			},
// 		},
// 	}
// 	resp, err := client.Check(context.Background(), request)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(resp)
// }

package main

import (
	"context"
	"fmt"
	"log"

	pbv0 "github.com/authzed/authzed-go/proto/authzed/api/v0"
	pbv1 "github.com/authzed/authzed-go/proto/authzed/api/v1alpha1"
	authzedv0 "github.com/authzed/authzed-go/v0"
	authzedv1 "github.com/authzed/authzed-go/v1alpha1"
	"google.golang.org/grpc"
)

const schema = `
definition blog/user {}

definition blog/post {
    relation reader: blog/user
    relation writer: blog/user

    permission read = reader + writer
    permission write = writer
}


`

func main() {
	setup()
	add_permission()
	check()
}

func setup() {
	client, err := authzedv1.NewClient(
		"localhost:50051",
		// grpcutil.WithBearerToken("t_your_token_here_1234567deadbeef"),
		// grpcutil.WithSystemCerts(false),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("unable to initialize client: %s", err)
	}

	request := &pbv1.WriteSchemaRequest{Schema: schema}
	resp, err := client.WriteSchema(context.Background(), request)
	if err != nil {
		log.Fatalf("failed to write schema: %s", err)
	}
	fmt.Println("Output", resp)
}

func add_permission() {
	client, err := authzedv0.NewClient(
		"localhost:50051",
		// grpcutil.WithBearerToken("t_your_token_here_1234567deadbeef"),
		// grpcutil.WithSystemCerts(false),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("unable to initialize client: %s", err)
	}

	request := &pbv0.WriteRequest{Updates: []*pbv0.RelationTupleUpdate{
		{ // Emilia is a Writer on Post 1
			Operation: pbv0.RelationTupleUpdate_CREATE,
			Tuple: &pbv0.RelationTuple{
				User: &pbv0.User{UserOneof: &pbv0.User_Userset{Userset: &pbv0.ObjectAndRelation{
					Namespace: "blog/user",
					ObjectId:  "emilia",
					Relation:  "...",
				}}},
				ObjectAndRelation: &pbv0.ObjectAndRelation{
					Namespace: "blog/post",
					ObjectId:  "1",
					Relation:  "writer",
				},
			},
		},
	}}

	resp, err := client.Write(context.Background(), request)

	if err != nil {
		log.Fatalf("failed to check permission: %s", err)
	}

	fmt.Println(resp)
}

func check() {
	client, err := authzedv0.NewClient(
		"localhost:50051",
		// grpcutil.WithBearerToken("t_your_token_here_1234567deadbeef"),
		// grpcutil.WithSystemCerts(false),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("unable to initialize client: %s", err)
	}

	ctx := context.Background()

	emilia := &pbv0.User{UserOneof: &pbv0.User_Userset{Userset: &pbv0.ObjectAndRelation{
		Namespace: "blog/post",
		ObjectId:  "emilia",
		Relation:  "...",
	}}}

	// beatrice := &pbv0.User{UserOneof: &pbv0.User_Userset{Userset: &pbv0.ObjectAndRelation{
	// 		Namespace: "user",
	// 		ObjectId:  "beatrice",
	// 		Relation:  "...",
	// }}}

	post1Reader := &pbv0.ObjectAndRelation{Namespace: "blog/post", ObjectId: "1", Relation: "reader"}
	// post1Writer := &pbv0.ObjectAndRelation{Namespace: "post", ObjectId: "1", Relation: "writer"}

	resp, err := client.Check(ctx, &pbv0.CheckRequest{User: emilia, TestUserset: post1Reader})
	if err != nil {
		log.Fatalf("failed to check permission: %s", err)
	}

	fmt.Println(resp)

}
