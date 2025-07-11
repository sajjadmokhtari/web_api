package config

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Password PasswordConfig
	Logger   LoggerConfig
	Otp      OtpConfig
	JWT      JwTConfig
}

type ServerConfig struct {
	Port    string
	RunMode string
}

type LoggerConfig struct {
	Filepath string
	Encoding string
	Level    string
	Logger   string
}

type PostgresConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DbName          string
	SSLMode         string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

type RedisConfig struct {
	Host               string
	Port               string
	User               string
	Password           string
	Db                 string
	SSLMode            bool
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	IdleCheckFrequency time.Duration
	PoolSize           int
	PoolTimeout        time.Duration
}

type PasswordConfig struct {
	IncludeChars     bool
	IncludeDigits    bool
	MinLength        int
	MaxLength        int
	IncludeUppercase bool
	IncludeLowercase bool
}

type OtpConfig struct {
	ExpireTime time.Duration
	Digits     int
	Limiter    time.Duration
}

type JwTConfig struct {
	Secret                     string
	RefreshSecret              string
	AccessTokenExpireDuration  time.Duration
	RefreshTokenExpireDuration time.Duration
}

func GetConfig() *Config {
	cfgPath := getConfigpath(os.Getenv("APP_ENV"))
	v, err := LoadConfig(cfgPath, "yml")

	if err != nil {
		log.Fatalf("error in load config: %v", err)

	}
	cfg, err := ParsConfig(v)

	if err != nil {
		log.Fatalf("error in parse config: %v", err)

	}

	return cfg

}

func ParsConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("unable to Parse Config %s", err)
		return nil, err
	}
	return &cfg, err

} // tabdil be struc config

func LoadConfig(filename string, fileType string) (*viper.Viper, error) {

	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(filename)

	// Add multiple config paths to handle different working directories
	v.AddConfigPath(".")
	v.AddConfigPath("..")
	v.AddConfigPath("../..")
	v.AddConfigPath("config")
	v.AddConfigPath("../config")
	v.AddConfigPath("../../config")
	v.AddConfigPath("src/config")
	v.AddConfigPath("../src/config")

	v.AutomaticEnv()

	err := v.ReadInConfig()

	if err != nil {
		log.Printf("unable to read Config %s", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")

		}
		return nil, err
	}
	return v, nil

} //gereft file ro va tabdil kard be struct viper

func getConfigpath(env string) string {
	if env == "docker" {
		return "config-docker"
	} else if env == "production" {
		return "config-production"
	} else {
		return "config-development"
	}

} //gereftan file
