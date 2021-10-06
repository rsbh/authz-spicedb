package permission

import (
	"log"

	authzed "github.com/authzed/authzed-go/v0"
	"google.golang.org/grpc"
)

type Permission struct {
	client *authzed.Client
}

func New(endpoint string, opts ...grpc.DialOption) *Permission {
	client, err := authzed.NewClient(endpoint, opts...)

	if err != nil {
		log.Fatalf("unable to initialize client: %s", err)
	}

	return &Permission{
		client: client,
	}
}
