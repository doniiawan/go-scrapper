package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Env Env
}

type Env struct {
	Urls struct {
		TokopediaUrl string `yaml:"tokopedia_url"`
	} `yaml:"urls"`
	Databases struct {
		Mysql struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			Database string `yaml:"database"`
		} `yaml:"mysql"`
	} `yaml:"databases"`
}

func InitConfig() (c Config, err error) {
	envByte, err := os.ReadFile(".env.yaml")
	if err != nil {
		return c, err
	}

	err = yaml.Unmarshal(envByte, &c.Env)
	if err != nil {
		return c, err
	}
	return
}
