package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"time"
)

type Config struct {
	Group string `mapstructure:"group"`
}

func GetEnvInfo() bool {
	viper.AutomaticEnv()
	fmt.Println("env:", os.Getenv("ENV"))
	return viper.GetBool("ENV")
}
func main() {
	v := viper.New()
	s, _ := os.Getwd()
	fmt.Println(s)
	v.SetConfigFile("./config-debug.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	cfg := new(Config)
	v.Unmarshal(cfg)
	viper.AutomaticEnv()
	v.WatchConfig()
	// 监听变化
	v.OnConfigChange(func(in fsnotify.Event) {
		v.ReadInConfig()
		v.Unmarshal(cfg)
		fmt.Printf("changed:%+v\n", cfg)
	})
	fmt.Println(GetEnvInfo())
	fmt.Println(v.Get("group"), cfg.Group)
	time.Sleep(time.Minute)
}
