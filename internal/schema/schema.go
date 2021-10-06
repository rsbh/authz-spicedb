package schema

import (
	"log"

	authzed "github.com/authzed/authzed-go/v1alpha1"
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

type Schema struct {
	client *authzed.Client
}

func New(endpoint string, opts ...grpc.DialOption) *Schema {
	client, err := authzed.NewClient(endpoint, opts...)

	if err != nil {
		log.Fatalf("unable to initialize client: %s", err)
	}

	return &Schema{
		client: client,
	}
}
