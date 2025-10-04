package user

import (
	"context"
)

type Repo interface {
	FindUserByMobile(ctx context.Context, mobile string) (*User, error)
}
