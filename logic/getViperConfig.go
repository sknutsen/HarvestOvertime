package logic

import "github.com/spf13/viper"

func GetViperConfig() *viper.Viper {
	var vp *viper.Viper = viper.New()

	vp.SetConfigName("appsettings")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")

	return vp
}
