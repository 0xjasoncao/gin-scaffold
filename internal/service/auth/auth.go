package auth

import (
	"context"
	"github.com/0xjasoncao/gin-scaffold/internal/apis/dto"
)

type Service interface {
	Login(ctx context.Context, request dto.LoginRequest) error
}

func NewService() Service {
	return &service{}
}

type service struct{}

func (srv *service) Login(ctx context.Context, request dto.LoginRequest) error {
	return nil
}
