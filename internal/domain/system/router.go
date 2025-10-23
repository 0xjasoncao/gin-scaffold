package system

import "gin-scaffold/internal/domain/shared"

type Router struct {
	Group   string
	Path    string
	Comment string
	Method  string
	shared.BasicInfo
}
