package config

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

func TestViper(t *testing.T) {

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.ReadInConfig()

	fmt.Println(viper.GetString("peer.privateKey"))
	fmt.Println(viper.GetString("cli.address"))

}
