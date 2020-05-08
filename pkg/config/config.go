package config

import (
	"github.com/pufferApi/pkg/proxy"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Puffers map[string]proxy.Config
}

func YAML() (config *Config, err error) {
	return YAMLfromFile("./Pufferfile")
}

func YAMLfromFile(file string) (config *Config, err error) {
	f, err := os.Open(file)
	if err != nil {
		return &Config{}, err
	}
	err = yaml.NewDecoder(f).Decode(&config)
	return config, err
}