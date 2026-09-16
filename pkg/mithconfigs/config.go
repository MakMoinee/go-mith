package mithconfigs

import "github.com/spf13/viper"

var Registry *viper.Viper

// Set reading the configurations from settings.yaml
func Set() {
	var err error
	viper.AddConfigPath(".")
	viper.AddConfigPath("../..")
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	Registry = viper.GetViper()

}
