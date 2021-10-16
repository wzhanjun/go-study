package config

import (
	"gopkg.in/ini.v1"
)

type Config struct {
	StockConfig StockConfig
}

type StockConfig struct {
	Codes string `ini:"codes"`
}

func NewConfig() (configs Config, err error) {
	cfg, err := ini.Load("./config/config.ini")
	if err != nil {
		return configs, err
	}
	err = cfg.Section("stock").MapTo(&configs.StockConfig)
	if err != nil {
		return configs, err
	}
	return configs, nil
}
