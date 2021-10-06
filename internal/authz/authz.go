package authz

import (
	"github.com/rsbh/authz-spicedb/internal/permission"
	"github.com/rsbh/authz-spicedb/internal/schema"
	"google.golang.org/grpc"
)

type Authz struct {
	Permission *permission.Permission
	Schema     *schema.Schema
}

func New(endpoint string, opts ...grpc.DialOption) *Authz {
	authz := Authz{}
	authz.Permission = permission.New(endpoint, opts...)
	authz.Schema = schema.New(endpoint, opts...)
	return &authz
}
