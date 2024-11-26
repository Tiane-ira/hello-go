package configs

import (
	"fmt"
	"hello-go/utils/env"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var config = new(Configs)

type Configs struct {
	App struct {
		Name      string `yaml:"name" json:"name" `
		Version   string `yaml:"version" json:"version"`
		Port      int    `yaml:"port" json:"port"`
		ApiPrefix string `yaml:"api-prefix" json:"apiPrefix" mapstructure:"api-prefix"`
		LogFile   string `yaml:"log-file" json:"logFile" mapstructure:"log-file"`
	} `yaml:"app" json:"app"`
	Mysql struct {
		Host          string `yaml:"host" json:"host"`
		Port          string `yaml:"port" json:"port"`
		Username      string `yaml:"username" json:"username"`
		Password      string `yaml:"password" json:"password"`
		Db            string `yaml:"db" json:"db"`
		Params        string `yaml:"params" json:"params"`
		TablePrefix   string `yaml:"table-prefix" json:"tablePrefix" mapstructure:"table-prefix"`
		SingularTable bool   `yaml:"singular-table" json:"singularTable" mapstructure:"singular-table"`
		MaxIdleConns  int    `yaml:"max-idle-conns" json:"maxIdleConns" mapstructure:"max-idle-conns"`
		MaxOpenConns  int    `yaml:"max-open-conns" json:"maxOpenConns" mapstructure:"max-open-conns"`
		LogMode       string `yaml:"log-mode" json:"logMode" mapstructure:"log-mode"`
	} `yaml:"mysql" json:"mysql"`

	Redis struct {
		Host         string `yaml:"host" json:"host"`
		Port         string `yaml:"port" json:"port"`
		Db           int    `yaml:"db" json:"db"`
		Password     string `yaml:"password" json:"password"`
		MaxRetries   int    `yaml:"max-retries" json:"maxRetries" mapstructure:"max-retries"`
		MaxIdleConns int    `yaml:"max-idle-conns" json:"maxIdleConns" mapstructure:"max-idle-conns"`
		PoolSize     int    `yaml:"pool-size" json:"poolSize" mapstructure:"pool-size"`
	} `yaml:"redis" json:"redis"`
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
