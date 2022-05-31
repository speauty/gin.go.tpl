package config

import (
	"fmt"
	"gin.go.tpl/lib/constant"
	"github.com/spf13/viper"
)

type GlobalConf struct {
	ENV      string `mapstructure:"env"`
	DBDriver string `mapstructure:"db_driver"`
}

type GinConf struct {
	Mode string `mapstructure:"mode"`
}

type LogConf struct {
	Level uint32 `mapstructure:"level"`
}

type ServerConf struct {
	Protocol string  `mapstructure:"protocol"`
	Domain   string  `mapstructure:"domain"`
	Port     string  `mapstructure:"port"`
	SSL      SSLConf `mapstructure:"ssl"`
}

func (sc ServerConf) GetAddr() string {
	addr := sc.Domain
	if sc.Port != "" {
		addr += ":" + sc.Port
	}
	return addr
}

func (sc ServerConf) IsHttp() bool {
	return sc.Protocol == constant.ProtocolHttp
}

func (sc ServerConf) IsHttps() bool {
	return sc.Protocol == constant.ProtocolHttps
}

type SSLConf struct {
	Certificate    string `mapstructure:"certificate"`
	CertificateKey string `mapstructure:"certificate_key"`
}

type PgSqlConf struct {
	Source string `mapstructure:"source"`
}

type RedisConf struct {
	Host      string `mapstructure:"host"`
	Port      string `mapstructure:"port"`
	Auth      string `mapstructure:"auth"`
	DB        string `mapstructure:"db"`
	MaxIdle   int    `mapstructure:"max_idle"`
	MaxActive int    `mapstructure:"max_active"`
}

type Config struct {
	Global GlobalConf `mapstructure:"global"`
	Gin    GinConf    `mapstructure:"gin"`
	Log    LogConf    `mapstructure:"log"`
	Server ServerConf `mapstructure:"server"`
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
