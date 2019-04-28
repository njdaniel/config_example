package main

import (
	"fmt"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	configName := flag.StringP("config", "c", "", "Configuration path filename")
	flag.Parse()
	fmt.Println(*configName)
	viper.SetConfigFile(*configName)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		// Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	fmt.Printf("name: %s\n", viper.GetString("name"))
}
