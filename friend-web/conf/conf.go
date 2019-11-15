package conf

import (
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
	Web              Web     `toml:"web"`
	Etcd             Etcd    `toml:"etcd"`
	Micro            Micro   `toml:"micro"`
	Remote           Remote  `toml:"remote"`
	Hystrix          Hystrix `toml:"hystrix"`
	LogLevel         int8    `toml:"loglevel"`
	Workspace        string  `toml:"workspace"`
	RootPackageSlash int     `toml:"root_package_slash"`
}

//Web .
type Web struct {
	Port string `toml:"port"`
}

//Hystrix .
type Hystrix struct {
	DefaultTimeout               int `toml:"default_timeout"`
	DefaultMaxConcurrent         int `toml:"default_max_concurrent"`
	DefaultVolumeThreshold       int `toml:"default_volume_threshold"`
	DefaultSleepWindow           int `toml:"default_sleep_window"`
	DefaultErrorPercentThreshold int `toml:"default_error_percent_threshold"`
}

//Remote .
type Remote struct {
	Auth string `toml:"auth"`
	User string `toml:"user"`
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
