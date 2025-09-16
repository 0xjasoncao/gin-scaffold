package parser

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

const (
	YAML = "yaml"
	YML  = "yml"
	JSON = "json"
	TOML = "toml"
)

// Parser 配置解析器接口
type Parser interface {
	Parse(data []byte, v interface{}) error
}

type parser func(data []byte, v interface{}) error

func (p parser) Parse(data []byte, v interface{}) error { return p(data, v) }

// 预定义解析器实例
var (
	// YAMLParser YAML格式解析器
	YAMLParser Parser = parser(func(data []byte, v interface{}) error {
		return yaml.Unmarshal(data, v)
	})

	// JSONParser JSON格式解析器
	JSONParser Parser = parser(func(data []byte, v interface{}) error {
		return json.Unmarshal(data, v)
	})

	// TOMLParser TOML格式解析器
	TOMLParser Parser = parser(func(data []byte, v interface{}) error {
		_, err := toml.Decode(string(data), v)
		return err
	})
)

func GetParser(fileType string) Parser {
	switch fileType {
	case YAML, YML:
		return YAMLParser
	case JSON:
		return JSONParser
	case TOML:
		return TOMLParser
	default:
		return nil
	}
}
