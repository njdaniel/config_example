package main

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

func main() {
	configName := flag.StringP("config", "c", "", "Configuration path filename")
	flag.Parse()
	fmt.Println(*configName)
}
