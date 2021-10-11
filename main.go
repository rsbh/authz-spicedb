package main

import (
	"time"

	"github.com/rsbh/authz-spicedb/internal/authz"
	"google.golang.org/grpc"
)

func main() {
	a := authz.New(
		"localhost:50051",
		grpc.WithInsecure(),
	)
	a.Schema.Load()
	a.Permission.Add("group:g1#member@user:u1")
	a.Permission.Add("group:f_admins#member@user:u2")
	a.Permission.Add("resource/firehose:f1#manager@group:g1")
	a.Permission.Add("project:p1#firehose_admins@group:f_admins#member")
	a.Permission.Add("resource/firehose:f1#manager@project:p1#firehose_admins")
	//
	a.Permission.Check("group:g1#view@user:u1")
	a.Permission.Check("resource/firehose:f1#manage@user:u1")
	a.Permission.Check("resource/firehose:f1#manage@user:u2")

	time.Sleep(3 * time.Second)
	a.Permission.Remove("group")


	time.Sleep(3 * time.Second)
	a.Permission.Check("group:g1#view@user:u1")
	a.Permission.Check("resource/firehose:f1#manage@user:u1")
	a.Permission.Check("resource/firehose:f1#manage@user:u2")
}
