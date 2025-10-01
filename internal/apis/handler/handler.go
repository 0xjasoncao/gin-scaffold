package handler

import (
	"github.com/0xjasoncao/gin-scaffold/internal/apis/handler/V1/auth"
	"github.com/0xjasoncao/gin-scaffold/internal/apis/handler/V1/user"
)

type V1 struct {
	User *user.Handler
	Auth *auth.Handler
}

type Handler struct {
	V1 *V1
}
