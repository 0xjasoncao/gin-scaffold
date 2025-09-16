package config

import (
	"context"
	"fmt"
	"github.com/0xjasoncao/gin-scaffold/configs/config"
	"github.com/0xjasoncao/gin-scaffold/pkg/logging"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"log"
	"testing"
)

func init() {
	err := config.Load("/Users/jasoncao/Develop/myDevelop/gin-scaffold/config/dev")
	if err != nil {
		log.Panic(err)
	}
	config.C.PrintWithJSON()
	logging.InitLogger(config.C)

}

func TestLoggingConfig(t *testing.T) {
	ctx := logging.NewStackContext(context.Background(), errors.New("test error "))

	ctx = logging.NewTagContext(ctx, "log-test")
	logging.WithContext(ctx).Info("hello")
	zap.L().Error("错误")
	zap.L().Info("错误3")

	production, _ := zap.NewDevelopment()
	fmt.Println()
	production.With(zap.String("tag", "test")).Error("sdfsdf")

}

func TestLogging(t *testing.T) {

}
