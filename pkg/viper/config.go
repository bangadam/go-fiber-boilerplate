package viper

import (
	"time"

	"github.com/spf13/viper"
)

type EnvConfig struct {
	FileName string
	FileType string
	Path string
	IdleTimeout time.Duration
}

func (e *EnvConfig) ReadConfig() error {
	// set viper config
	viper.SetConfigName(e.FileName)
	// set config file type
	viper.SetConfigType(e.FileType)
	// set config file path
	viper.AddConfigPath(e.Path)
	// read config file
	viper.AutomaticEnv()
	viper.WatchConfig()
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}