package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func main() {
	viper.SetEnvPrefix("spf") // will be uppercased automatically
	viper.BindEnv("id")

	os.Setenv("SPF_ID", "13") // typically done outside of the app
	id := viper.Get("id")     // 13
	fmt.Println(id)
}
