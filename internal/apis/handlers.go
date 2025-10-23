package apis

import (
	"gin-scaffold/internal/apis/docs"
	"gin-scaffold/internal/apis/system"
	"github.com/google/wire"
)

type RouterHandlers struct {
	Swagger *docs.SwaggerHandler
	System  *system.Handlers
	//v2 *v2.Handlers
}

var ProviderSet = wire.NewSet(
	docs.NewSwaggerHandler,
	system.ProviderSet,
	wire.Struct(new(RouterHandlers), "*"),
)
