package permission

import (
	"log"

	authzed "github.com/authzed/authzed-go/v0"
	"google.golang.org/grpc"
)

type PermissionHandler struct {
	client *authzed.Client
}

func NewHandler() *PermissionHandler {
	client, err := authzed.NewClient(
		"localhost:50051",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("unable to initialize client: %s", err)
	}

	handler := &PermissionHandler{
		client: client,
	}

	return handler
}
