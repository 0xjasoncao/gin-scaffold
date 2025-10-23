package config

// JWT jwt token config
type JWT struct {
	Key              string `mapstructure:"key" json:"secret"`                            // 签名密钥（HMAC算法使用）
	ExpiresAtSeconds int    `mapstructure:"expires-at-seconds" json:"expires-at-seconds"` // 访问令牌过期时间（单位：秒）
	Issuer           string `mapstructure:"issuer" json:"issuer"`                         // 令牌签发者（可选，用于标识服务）
}
