package cfg

import (
	"fmt"
	"gin.go.tpl/kernel/constant"
	"github.com/spf13/viper"
	"sync"
)

var (
	api  *Cfg
	once sync.Once
)

func NewCfgApi(path string) *Cfg {
	once.Do(func() {
		api, _ = Cfg{}.LoadConfig(path)
	})
	return api
}

type GlobalConf struct {
	ENV string `mapstructure:"env"`
}

type GinConf struct {
	Mode string `mapstructure:"mode"`
}

type LogConf struct {
	Level            uint32 `mapstructure:"level"`
	LogFile          string `mapstructure:"log_file"`
	LogRotationTime  int    `mapstructure:"log_rotation_time"`
	LogMaxAge        int    `mapstructure:"log_max_age"`
	LogRotationCount uint   `mapstructure:"log_rotation_count"`
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

type DatabaseConf struct {
	IsMigration            bool      `mapstructure:"is_migration"`
	DBDriver               string    `mapstructure:"db_driver"`
	PrefixTable            string    `mapstructure:"prefix_table"`
	SingularTable          bool      `mapstructure:"singular_table"`
	NoLowerCase            bool      `mapstructure:"no_lower_case"`
	IsPool                 int       `mapstructure:"is_pool"`
	MaxIdleConn            int       `mapstructure:"max_idle_conn"`
	MaxOpenConn            int       `mapstructure:"max_open_conn"`
	MaxLifetime            uint      `mapstructure:"max_lifetime"`
	SkipDefaultTransaction bool      `mapstructure:"skip_default_transaction"`
	PrepareStmt            bool      `mapstructure:"prepare_stmt"`
	PgSql                  PgSqlConf `mapstructure:"pgsql"`
	MySql                  MySqlConf `mapstructure:"mysql"`
}

type PgSqlConf struct {
	DSN string `mapstructure:"dsn"`
}

type MySqlConf struct {
	DSN string `mapstructure:"dsn"`
}

type RedisConf struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	Auth string `mapstructure:"auth"`
	DB   int    `mapstructure:"db"`
}

type LimiterConf struct {
	GeneratorInterval uint `mapstructure:"generator_interval"`
	GeneratorNum      int  `mapstructure:"generator_num"`
}

type Cfg struct {
	Global   GlobalConf   `mapstructure:"global"`
	Gin      GinConf      `mapstructure:"gin"`
	Log      LogConf      `mapstructure:"log"`
	Server   ServerConf   `mapstructure:"server"`
	Database DatabaseConf `mapstructure:"database"`
	PgSql    PgSqlConf    `mapstructure:"pgsql"`
	MySql    MySqlConf    `mapstructure:"mysql"`
	Redis    RedisConf    `mapstructure:"redis"`
	Limiter  LimiterConf  `mapstructure:"limiter"`
}

func (Cfg) LoadConfig(path string) (config *Cfg, err error) {
	config = &Cfg{}

	viper.AddConfigPath(path)
	viper.SetConfigType("ini")
	viper.SetConfigName("app.global")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("当前配置加载失败, 错误: %v", err))
	}

	err = viper.Unmarshal(config)
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
	err = viper.Unmarshal(config)
	if err != nil {
		panic(fmt.Errorf("%s配置加载失败, 错误: %v", config.Global.ENV, err))
	}

	return config, err
}
