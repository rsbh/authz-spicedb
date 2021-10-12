package main

import (
	"testing"

	"github.com/rsbh/authz-spicedb/internal/authz"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func setup() *authz.Authz {
	a := authz.New(
		"localhost:50051",
		grpc.WithInsecure(),
	)
	a.Schema.Load()
	return a
}

func TestRules(t *testing.T) {
	t.Run("Should return true if user is part of group", func(t *testing.T) {
		a := setup()
		a.Permission.Add("group:g1#member@user:u1")
		result := a.Permission.Check("group:g1#view@user:u1")
		assert.Equal(t, result, true)
	})

	t.Run("Should return false if user is not part of group", func(t *testing.T) {
		a := setup()
		result := a.Permission.Check("group:g1#view@user:u2")
		assert.Equal(t, result, false)
	})

	t.Run("Should return false if user doesnt have access to resource", func(t *testing.T) {
		a := setup()
		result := a.Permission.Check("resource/firehose:f1#manage@user:u2")
		assert.Equal(t, result, false)
	})

	t.Run("Should return true if user has access to resource", func(t *testing.T) {
		a := setup()
		a.Permission.Add("group:g1#member@user:u1")
		a.Permission.Add("resource/firehose:f1#manager@group:g1")
		result := a.Permission.Check("resource/firehose:f1#manage@user:u1")
		assert.Equal(t, result, true)
	})

}
