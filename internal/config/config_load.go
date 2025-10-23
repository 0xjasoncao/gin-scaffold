package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"reflect"
)

// MustLoad 加载文件夹下的所有配置到config.C中
func MustLoad(configDir string) {
	viper.AutomaticEnv()
	entries, err := os.ReadDir(configDir)
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		fullPath := filepath.Join(configDir, entry.Name())
		if entry.IsDir() {
			continue
		}
		loadFile(fullPath)
		fmt.Println("[CONFIG] - ", fullPath, "load successfully.")
	}
	if err := viper.Unmarshal(C); err != nil {
		panic(err)
	}
	C.Dir = configDir
	InitStructRefs(C)

}

// 加载配置到viper中
func loadFile(fullPath string) {
	v := viper.New()
	v.SetConfigFile(fullPath)

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.MergeConfigMap(v.AllSettings()); err != nil {
		panic(err)
	}

}

// InitStructRefs 遍历结构体字段，如果是指针、slice、map为 nil，初始化为默认值
func InitStructRefs(ptr any) {
	v := reflect.ValueOf(ptr)
	if v.Kind() != reflect.Pointer || v.IsNil() {
		panic("InitStructRefs requires a non-nil pointer to struct")
	}

	v = v.Elem()
	if v.Kind() != reflect.Struct {
		panic("InitStructRefs requires pointer to struct")
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		// 如果字段是 unexported，跳过
		if !field.CanSet() {
			continue
		}
		switch field.Kind() {
		case reflect.Slice:
			if field.IsNil() {
				field.Set(reflect.MakeSlice(field.Type(), 0, 0))
			}
		case reflect.Map:
			if field.IsNil() {
				field.Set(reflect.MakeMap(field.Type()))
			}
		case reflect.Pointer:
			if field.IsNil() {
				// 创建一个新的值
				newVal := reflect.New(field.Type().Elem())
				field.Set(newVal)
			} else {
				// 递归初始化指针指向的结构体
				InitStructRefs(field.Interface())
			}
		case reflect.Struct:
			// 递归初始化嵌套结构体
			InitStructRefs(field.Addr().Interface())

		}
	}
}
