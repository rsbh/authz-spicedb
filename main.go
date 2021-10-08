package main

import (
	"github.com/rsbh/authz-spicedb/internal/authz"
	"google.golang.org/grpc"
)

func main() {
	a := authz.New(
		"localhost:50051",
		grpc.WithInsecure(),
	)
	a.Schema.Load()
	a.Permission.Add("group:g1", "user:user1", "member")
	a.Permission.Check("group:g1", "user:user1", "member")
	a.Permission.Check("group:g1", "user:user2", "member")

}
