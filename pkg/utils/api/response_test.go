package api

import "testing"

type User struct {
	Name string
}

func TestResList(t *testing.T) {
	var list []User
	ResList(nil, list)
}
