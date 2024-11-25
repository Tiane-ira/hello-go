package configs

import (
	"fmt"
	"hello-go/pkg/env"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var config = new(Configs)
var data []byte

type Configs struct {
	Mysql struct {
		Host            string `yaml:"host"`
		Port            string `yaml:"port"`
		Db              string `yaml:"db"`
		Username        string `yaml:"username"`
		Password        string `yaml:"password"`
		ConnMaxLifeTine int    `yaml:"connMaxLifeTine"`
		MaxIdleConn     int    `yaml:"maxIdleConn"`
		MaxOpenConn     int    `yaml:"maxOpenConn"`
	} `yaml:"mysql"`

	Redis struct {
		Host        string `yaml:"host"`
		Port        string `yaml:"port"`
		Db          string `yaml:"db"`
		Password    string `yaml:"password"`
		MaxRetries  int    `yaml:"maxRetries"`
		MaxIdleConn int    `yaml:"maxIdleConn"`
		PoolSize    int    `yaml:"poolSize"`
	} `yaml:"redis"`
}

func init() {
	configName := fmt.Sprintf("app-%s", env.Active())
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(ConfigPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	// 动态加载配置
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
}

func Get() *Configs {
	return config
}
