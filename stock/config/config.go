package config

import (
	"gopkg.in/ini.v1"
)

type Config struct {
	StockConfig  StockConfig
	XueQiuConfig XueQiuConfig
}

type StockConfig struct {
	Codes string `ini:"codes"`
}

type XueQiuConfig struct {
	CookieXQAToken string `ini:"cookie_xq_a_token"`
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

	err = cfg.Section("xueqiu").MapTo(&configs.XueQiuConfig)
	if err != nil {
		return configs, err
	}
	return configs, nil
}
