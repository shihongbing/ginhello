package config

import (
	"ginhello/logger"
	"io/ioutil"
	"time"
)
import "gopkg.in/yaml.v2"

type app struct {
	AppName string `yaml:"appName"`
	Port    string `yaml:"port"`
	Debug	bool   `yaml:"debug"`
}

type db struct {
	DataBaseUrl    string        `yaml:"databaseUrl"`
	MaxIdleConns    int           `yaml:"maxIdleConns"`
	MaxOpenConns    int           `yaml:"maxOpenConns"`
	ConnMaxLifeTime int           `yaml:"connMaxLifetime"`
	ShowSql         bool          `yaml:"showSql"`
	ENGINE          string        `yaml:"engine"`
	CHARSET         string        `yaml:"charset"`
	PREFIX          string        `yaml:"prefix"`
	SlowThreshold   time.Duration `yaml:"slowThreshold"`
}

type log struct {
	Level string `yaml:"level"`
}

//Config   系统配置
type config struct{
	AppConfig app `yaml:"app"`
	DBConfig db `yaml:"datasource"`
	LogConfig log `yaml:"log"`
}

var (
	AppConfig   app
	DBConfig    db
	LogConfig   log
)

/**
 初始化系统配置
 */
func init(){
	application, err := ioutil.ReadFile("application.yaml")
	if err != nil {
		logger.Logger.Error("LoadingConfig error: %s ")
	}
	conf := new(config)
	yaml.Unmarshal(application,&conf)
	AppConfig = conf.AppConfig
	DBConfig = conf.DBConfig
	LogConfig = conf.LogConfig
}
