package config

type Redis struct {
	Addr           string   `yaml:"addr"`
	DB             int      `yaml:"db"`
	Password       string   `yaml:"password"`
	UseCluster     bool     `yaml:"use-cluster"`
	ClusterAddress []string `yaml:"cluster-address"`
	Open           bool     `yaml:"open"`
}
