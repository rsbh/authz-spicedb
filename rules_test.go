package main

import (
	"testing"

	"github.com/rsbh/authz-spicedb/internal/authz"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func setup() *authz.Authz {
	return authz.New(
		"localhost:50051",
		grpc.WithInsecure(),
	)
}

func TestRules(t *testing.T) {
	t.Run("Should return true if user is part of group", func(t *testing.T) {
		a := setup()
		a.Permission.Add("group:g1#member@user:u1")
		result := a.Permission.Check("group:g1#view@user:u1")
		assert.Equal(t, result, true)
	})

	t.Run("Should return false if user is part of group", func(t *testing.T) {
		a := setup()
		result := a.Permission.Check("group:g1#view@user:u2")
		assert.Equal(t, result, false)
	})
}
