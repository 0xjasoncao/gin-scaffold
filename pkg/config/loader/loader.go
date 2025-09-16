package loader

import "context"

// Loader 配置加载器接口
type Loader interface {
	// Load 加载配置数据
	Load() ([]byte, error)
	// Watch 监听配置变化
	Watch(ctx context.Context, onChange func([]byte)) error
}
