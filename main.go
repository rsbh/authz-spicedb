package main

import (
	"github.com/rsbh/authz-spicedb/internal/permission"
	"github.com/rsbh/authz-spicedb/internal/schema"
)

func main() {
	schema.Read()

	ph := permission.NewHandler()
	ph.Add()
	ph.Check()
}
