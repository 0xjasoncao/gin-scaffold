package token

import (
	"context"
	"fmt"
	"github.com/0xjasoncao/gin-scaffold/pkg/cache"
	"testing"
)

func init() {

}

func TestJwtToken_IssuingToken(t *testing.T) {
	st := &Settings{
		ExpiresAtSeconds: 120,
		Key:              DefaultKey,
		Issuer:           "0xJasoncao",
	}
	ctx := context.Background()

	memory, _ := cache.NewMemory()
	cache := memory
	s := NewStoreWithCache(cache)
	service, _ := NewTokenService(st, s)
	token, err := service.IssuingToken(ctx, "test userID")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("签发的token:")
	fmt.Println(token)

	parse, err := service.Parse(ctx, token)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("解析后的token:\n %+v", parse)

	fmt.Println("正在销毁token")
	err = service.DestroyToken(ctx, token)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(cache)
	fmt.Println(memory)
	v, _ := cache.Get(ctx, token)
	fmt.Println(v)
}
