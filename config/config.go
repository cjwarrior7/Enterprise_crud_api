package configs

import (
	"strings"
	"github.com/spf13/viper"
)

type cacheConfig struct {
	Host     string // CACHE_HOST
	PoolSize int    // CACHE_POOLSIZE
}

type logConfig struct {
	LogFile  string
	LogLevel string
}

type httpConfig struct {
	HostPort string
}

// Config - configuration object
type Config struct {
	Cache      cacheConfig
	Log        logConfig
	HttpConfig httpConfig
}

var conf *Config

// GetConfig - Function to get Config
func GetConfig() *Config {
	if conf != nil {
		return conf
	}
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	cacheConf := cacheConfig{
		Host:     v.GetString("cache.host"),
		PoolSize: v.GetInt("cache.poolsize"),
	}

	logConf := logConfig{
		LogFile:  v.GetString("log.file"),
		LogLevel: v.GetString("log.level"),
	}

	httpConf := httpConfig{
		HostPort: v.GetString("http.host"),
	}
	conf = &Config{
		Cache:      cacheConf,
		Log:        logConf,
		HttpConfig: httpConf,
	}
	return conf
}
