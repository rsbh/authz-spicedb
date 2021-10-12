package main

import (
	"fmt"
	"testing"
)

func BenchmarkRules(b *testing.B) {
	a := setup()

	for i := 0; i < 1000; i++ {
		a.Permission.Add(fmt.Sprintf("group:%d#member@user:u%d", i/10, i))
	}
	for i := 0; i < 1000; i++ {
		a.Permission.Add(fmt.Sprintf("resource/firehose:f%d#manager@group:g%d", i, i/10))
	}
	a.Permission.Add("group:f_admins#member@user:u2")
	a.Permission.Add("project:p1#firehose_admins@group:f_admins#member")
	a.Permission.Add("resource/firehose:f1#manager@project:p1#firehose_admins")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.Permission.Check("group:g1#view@user:u1")
	}

}
