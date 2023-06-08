package services

import (
	"os"
	"path"

	"github.com/uees/hidedomain/utils"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Listen    string `yaml:"listen"`
	SecretKey string `yaml:"secret_key"`
	Debug     bool   `yaml:"debug"`
	DbDriver  string `yaml:"driver"`
	DSN       string `yaml:"dsn"`
	//KeyMap    map[string]string `yaml:"keymap"`
}

func (conf *Config) LoadConf() error {
	data, err := os.ReadFile(path.Join(utils.BaseDir(), "config.yml"))
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, conf); err != nil {
		return err
	}

	return nil
}
