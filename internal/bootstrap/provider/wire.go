package provider

import (
	"github.com/google/wire"
)

var BasicProviderSet = wire.NewSet(
	InitGorm,
	InitRedisCli,
	InitMemoryCache,
	InitCache,
	InitTokenService,
)
