package util

import "github.com/spf13/viper"

// bakal nyimpan semua konfigurasi yang dipakek
// isinha bakal dibaca sama viper dari config file atau env
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	// viper bakal baca file config dari path
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// baca file config
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// unmarshal config ke struct
	err = viper.Unmarshal(&config)
	return
}
