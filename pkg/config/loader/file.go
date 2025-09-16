package loader

import (
	"context"
	"os"
)

// FileLoader 配置文件加载器
type FileLoader struct {
	location string
}

func NewFileLoader(path string) *FileLoader {
	fl := &FileLoader{location: path}
	return fl
}

func (f *FileLoader) Load() ([]byte, error) {
	return os.ReadFile(f.location)
}

func (f *FileLoader) Watch(ctx context.Context, onChange func([]byte)) error {
	return nil
}
