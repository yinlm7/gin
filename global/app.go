package global

import (
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      Configuration
	PoolRedis   *redis.Pool
	PoolMongo   *mongo.Client
}

type Configuration struct {
	App     Apps    `mapstructure:"app" json:"app" yaml:"app"`
	Redis   Redis   `mapstructure:"redis" json:"reids" yaml:"redis"`
	Mongodb Mongodb `mapstructure:"mongodb" json:"mongodb" yaml:"mongodb"`
}

type Apps struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env"`
	Port    string `mapstructure:"port" json:"port" yaml:"port"`
	AppName string `mapstructure:"app_name" json:"app_name" yaml:"app_name"`
	AppUrl  string `mapstructure:"app_url" json:"app_url" yaml:"app_url"`
}

type Mongodb struct {
	Addr      string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Port      string `mapstructure:"port" json:"port" yaml:"port"`
	User      string `mapstructure:"user" json:"user" yaml:"user"`
	Password  string `mapstructure:"password" json:"password" yaml:"password"`
	DefaultDB string `mapstructure:"defaultDB" json:"defaultDB" yaml:"defaultDB"`
	MgoUrl    string `mapstructure:"mgoUrl" json:"mgoUrl" yaml:"mgoUrl"`
	ClientNum string `mapstructure:"clientNum" json:"clientNum" yaml:"clientNum"`
}

type Redis struct {
	Addr        string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password    string `mapstructure:"password" json:"password" yaml:"password"`
	DefaultDB   int    `mapstructure:"defaultDB" json:"defaultDB" yaml:"defaultDB"`
	MaxIdle     int    `mapstructure:"maxIdle" json:"maxIdle" yaml:"maxIdle"`
	MaxActive   int    `mapstructure:"maxActive" json:"maxActive" yaml:"maxActive"`
	IdleTimeout int    `mapstructure:"idleTimeout" json:"idleTimeout" yaml:"idleTimeout"`
	Wait        bool   `mapstructure:"wait" json:"wait" yaml:"wait"`
}

var App = new(Application)
