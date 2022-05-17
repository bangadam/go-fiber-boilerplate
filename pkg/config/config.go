package config

import "github.com/spf13/viper"

type EnvConfig struct {
	FileName string
	Path     string
}

/* ReadConfig reads the config file stream into viper and return errors if any
* Example Instance struct
* FileName  	: "config.yaml"
* Path    		: "/temp/config/"
 */
func (e *EnvConfig) ReadConfig() error {
	viper.SetConfigName(e.FileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(e.Path)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
