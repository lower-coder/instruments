package config

import "github.com/spf13/viper"

// New 创建一个 viper 实例
//  @return *viper.Viper 
func New() *viper.Viper {
	return viper.New()
}
