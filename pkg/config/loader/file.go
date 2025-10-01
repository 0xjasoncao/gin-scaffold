package loader

import (
	"context"
	"os"
)

// LocalConfigLoader 本地配置文件加载器
type LocalConfigLoader struct {
	location string
}

func NewLocalConfigLoader(path string) *LocalConfigLoader {
	fl := &LocalConfigLoader{location: path}
	return fl
}

func (f *LocalConfigLoader) Load() ([]byte, error) {
	return os.ReadFile(f.location)
}

func (f *LocalConfigLoader) Watch(ctx context.Context, onChange func([]byte)) error {
	return nil
}
