package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

var (
	global *Config
)

// LoadGlobalConfig 加载全局配置
func LoadGlobalConfig(fpath string) error {
	c, err := ParseConfig(fpath)
	if err != nil {
		return err
	}
	global = c
	return nil
}

// GetGlobalConfig 获取全局配置
func GetGlobalConfig() *Config {
	if global == nil {
		return &Config{}
	}
	return global
}

// ParseConfig 解析配置文件
func ParseConfig(fpath string) (*Config, error) {
	var c Config
	_, err := toml.DecodeFile(fpath, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

//Config .
type Config struct {
	MySQL            MySQL    `toml:"mysql"`
	Etcd             Etcd     `toml:"etcd"`
	Micro            Micro    `toml:"micro"`
	RabbitMq         RabbitMq `toml:"rabbitmq"`
	LogLevel         int8     `toml:"loglevel"`
	Jaeger           Jaeger   `toml:"jaeger"`
	Memory           Memory   `toml:"memory_cache"`
	Workspace        string   `toml:"workspace"`
	RootPackageSlash int      `toml:"root_package_slash"`
	Secrete          string   `toml:"secrete"`
}

//Jaeger .
type Jaeger struct {
	ServiceName string `toml:"service_name"`
	URL         string `toml:"url"`
}

//Memory .
type Memory struct {
	DefaultExpiration int `toml:"default_expiration"`
	IntervalClear     int `toml:"interval_clear"`
}

//MySQL .
type MySQL struct {
	Debug      bool   `toml:"debug"`
	Host       string `toml:"host"`
	Port       int    `toml:"port"`
	User       string `toml:"user"`
	Password   string `toml:"password"`
	DBName     string `toml:"db_name"`
	Parameters string `toml:"parameters"`
}

//DSN .
func (a MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		a.User, a.Password, a.Host, a.Port, a.DBName, a.Parameters)
}

//RabbitMq .
type RabbitMq struct {
	User             string `toml:"user"`
	Password         string `toml:"password"`
	Host             string `toml:"host"`
	Port             int    `toml:"port"`
	ChatExchangeName string `toml:"chat_exchange_name"`
}

//DSN .
func (a RabbitMq) DSN() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/",
		a.User, a.Password, a.Host, a.Port)
}

//Etcd .
type Etcd struct {
	Addrs            []string `toml:"addrs"`
	RegisterTTL      int      `toml:"register_ttl"`
	RegisterInterval int      `toml:"register_interval"`
}

//Micro .
type Micro struct {
	Name    string `toml:"name"`
	Version string `toml:"version"`
}
