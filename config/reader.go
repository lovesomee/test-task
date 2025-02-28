package config

import "github.com/spf13/viper"

func Read() Settings {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var settings Settings
	if err := viper.Unmarshal(&settings); err != nil {
		panic(err)
	}

	return settings
}
