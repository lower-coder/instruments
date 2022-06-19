// package config viper.go.
package config

import "github.com/spf13/viper"

// ViperWrapper alias viper.Viper.
type ViperWrapper = viper.Viper

// NewViper for config wrapper.
//  @return *ViperWrapper
func NewViper() *ViperWrapper {
	return viper.New()
}
