package configs

import (
	"github.com/spf13/viper"
	"time"
)

type Server struct {
	Host            string
	Port            int
	RunMode         string
	ClientAuth      bool
	ReadTimeOut     int
	WriteTimeOut    int
	MaxHeaderBytes  int
	RateLimit       int
	RateLimitPeriod int
}

type Database struct {
	Host            string
	User            string
	Password        string
	Port            int
	Type            string
	Name            string
	SslMode         string
	TimeZone        string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
	ShowDbLog       bool
}

type Redis struct {
	Host     string
	Port     int
	Password string
	Db       int
	PoolSize int
}

type RabbitMq struct {
	Host     string
	Port     int
	User     string
	Password string
}

type JWT struct {
	Secret     string
	EffectTime time.Duration
}

type Config struct {
	Server    Server
	Database  Database
	Redis     Redis
	Rabbitmq  RabbitMq
	JWTConfig JWT
}

var config *Config

func InitConfig() {
	// 默认先加载dev 配置
	viper := viper.New()

	// 获取环境变量
	viper.AutomaticEnv()
	config = &Config{
		Server: Server{
			Host:            viper.GetString("SERVER_HOST"),
			Port:            viper.GetInt("SERVER_PORT"),
			RunMode:         viper.GetString("SERVER_RUN_MODE"),
			ClientAuth:      false,
			ReadTimeOut:     60,
			WriteTimeOut:    60,
			MaxHeaderBytes:  1048576,
			RateLimit:       1000,
			RateLimitPeriod: 3600,
		},
		Database: Database{
			Host:            viper.GetString("DB_HOST"),
			User:            viper.GetString("DB_USER"),
			Password:        viper.GetString("DB_PASSWORD"),
			Port:            viper.GetInt("DB_PORT"),
			Type:            viper.GetString("DB_TYPE"),
			Name:            viper.GetString("DB_NAME"),
			SslMode:         "disable",
			TimeZone:        "Asia/Shanghai",
			MaxIdleConns:    0,
			MaxOpenConns:    0,
			ConnMaxLifetime: 3600,
			ShowDbLog:       false,
		},
		Redis: Redis{
			Host:     viper.GetString("REDIS_HOST"),
			Port:     viper.GetInt("REDIS_PORT"),
			Password: viper.GetString("REDIS_PASSWORD"),
			Db:       0,
			PoolSize: 10,
		},
		Rabbitmq: RabbitMq{
			Host:     viper.GetString("RABBITMQ_HOST"),
			Port:     viper.GetInt("RABBITMQ_PORT"),
			User:     viper.GetString("RABBITMQ_USER"),
			Password: viper.GetString("RABBITMQ_PASSWORD"),
		},
		JWTConfig: JWT{
			Secret:     viper.GetString("JWT_SECRET"),
			EffectTime: time.Hour * 24,
		},
	}
}

func Get() Config {
	return *config
}
