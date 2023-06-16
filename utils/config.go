package utils

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Listen    string `yaml:"listen"`
	SecretKey string `yaml:"secret_key"`
	Debug     bool   `yaml:"debug"`
	DbDriver  string `yaml:"driver"`
	DSN       string `yaml:"dsn"`
	LogFile   string `yaml:"logfile"`
	//KeyMap    map[string]string `yaml:"keymap"`
}

func (conf *Config) LoadConf() error {
	data, err := os.ReadFile("config.yml")
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, conf); err != nil {
		return err
	}

	return nil
}
