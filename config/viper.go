// package config viper.go
package config

import "github.com/spf13/viper"

// NewViper for config wrapper.
//  @return *viper.Viper 
func NewViper() *viper.Viper {
	return viper.New()
}
