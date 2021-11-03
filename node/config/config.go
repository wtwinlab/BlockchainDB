package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Options struct {
	SelfID         string `mapstructure:"self-id"`
	ServerNodeAddr string `mapstructure:"server-node-addr"`
	EthNode        string `mapstructure:"eth-node"`
	EthHexAddr     string `mapstructure:"eth-hexaddr"`
	EthHexKey      string
}

func ReadConfig(opts interface{}, configfile string) (err error) {
	err = viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		return err
	}
	if configfile != "" {
		//viper.SetConfigFile(configfile)
		viper.SetConfigName(configfile)
	} else {
		viper.SetConfigName("blockchaindb")
	}

	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(opts)
	if err != nil {
		return err
	}

	return nil
}
