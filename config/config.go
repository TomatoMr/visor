package config

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Pid string `yaml:"pid"`
	SnapPath string `yaml:"snap_path"`
	AlterLimit float64 `yaml:"alter_limit"`
	Interval int `yaml:"interval"`
}

var conf Config

func InitConfig(configPath string) error {
	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return errors.Wrap(err, "Read config file failed")
	}

	if err = yaml.Unmarshal(configFile, &conf); err != nil {
		return errors.Wrap(err, "Unmarshal config file failed.")
	}
	return nil
}

func GetConfig() Config {
	return conf
}
