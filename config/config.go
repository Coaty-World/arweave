package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	ArweaveConfig *ArweaveConfig `mapstructure:"arweave"`
	CDNConfig     *CDN           `mapstructure:"cdn"`
}

type ArweaveConfig struct {
	Store string `mapstructure:"store"`
}

type CDN struct {
	CharacterItems string `mapstructure:"character_items"`
}

func NewConfig() (Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	config := Config{}

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}
	return config, nil
}
