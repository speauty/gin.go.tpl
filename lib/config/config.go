package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type GlobalConf struct {
	ENV      string `mapstructure:"env"`
	DBDriver string `mapstructure:"db_driver"`
}

type PgSqlConf struct {
	Source string `mapstructure:"source"`
}

type RedisConf struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	Auth string `mapstructure:"auth"`
	DB   string `mapstructure:"db"`
}

type Config struct {
	Global GlobalConf `mapstructure:"global"`
	PgSql  PgSqlConf  `mapstructure:"pgsql"`
	Redis  RedisConf  `mapstructure:"redis"`
}

func (Config) LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("ini")
	viper.SetConfigName("app.global")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("当前配置加载失败, 错误: %v", err))
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("默认配置加载失败, 错误: %v", err))
	}
	if config.Global.ENV == "" {
		panic("未检测到当前环境(env)")
	}
	viper.SetConfigName("app." + config.Global.ENV)
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("%s配置加载失败, 错误: %v", config.Global.ENV, err))
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("%s配置加载失败, 错误: %v", config.Global.ENV, err))
	}
	return config, err
}
