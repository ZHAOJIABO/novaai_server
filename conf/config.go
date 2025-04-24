package conf

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Server struct {
		HttpAddress string `yaml:"http_address"`
		GrpcAddress string `yaml:"grpc_address"`
	} `yaml:"server"`
}

func LoadConfig() (*Config, error) {
	config := &Config{}

	file, err := os.ReadFile("conf/config.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
