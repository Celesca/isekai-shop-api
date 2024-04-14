package config

import (
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// สิ่งที่เปลี่ยนแปลงได้ตลอดเวลา deploy ขึ้น cloud

type (
	Config struct {
		Server   *Server   `mapstructure:"server" validate:"required"`
		OAuth2   *OAuth2   `mapstructure:"oauth2" validate:"required"`
		State    *State    `mapstructure:"state" validate:"required"`
		Database *Database `mapstructure:"database" validate:"required"`
	}

	Server struct {
		Port         int           `mapstructure:"port" validate:"required"`
		AllowOrigins []string      `mapstructure:"allowOrigins" validate:"required"`
		Timeout      time.Duration `mapstructure:"timeout" validate:"required"`
		BodyLimit    string        `mapstructure:"bodyLimit" validate:"required"`
	}

	OAuth2 struct {
		PlayerRedirectUrl string   `mapstructure:"playerRedirectUrl" validate:"required"`
		AdminRedirectUrl  string   `mapstructure:"adminRedirectUrl" validate:"required"`
		ClientID          string   `mapstructure:"clientId" validate:"required"`
		ClientSecret      string   `mapstructure:"clientSecret" validate:"required"`
		Endpoints         endpoint `mapstructure:"endpoints" validate:"required"`
		Scopes            []string `mapstructure:"scopes" validate:"required"`
		UserInfoUrl       string   `mapstructure:"userInfoUrl" validate:"required"`
		RevokeUrl         string   `mapstructure:"revokeUrl" validate:"required"`
	}

	endpoint struct {
		AuthUrl       string `mapstructure:"authUrl" validate:"required"`
		TokenUrl      string `mapstructure:"tokenUrl" validate:"required"`
		DeviceAuthUrl string `mapstructure:"deviceAuthUrl" validate:"required"`
	}

	State struct {
		Secret    string        `mapstructure:"secret" validate:"required"`
		ExpiresAt time.Duration `mapstructure:"expiresAt" validate:"required"`
		Issuer    string        `mapstructure:"issuer" validate:"required"`
	}

	Database struct {
		Host     string `mapstructure:"host" validate:"required"`
		Port     int    `mapstructure:"port" validate:"required"`
		User     string `mapstructure:"user" validate:"required"`
		Password string `mapstructure:"password" validate:"required"`
		DBName   string `mapstructure:"dbname" validate:"required"`
		SSLMode  string `mapstructure:"sslmode" validate:"required"`
		Schema   string `mapstructure:"schema" validate:"required"`
	}
)

// Singleton Pattern จะต้องมี Instance เดียว เวลา Start App จะโหลดแค่ครั้งเดียว
// จะไม่มีการโหลดอีกแล้ว แม้จะมีการเรียกใช้อีกรอบกก็ตาม

var (
	once           sync.Once
	configInstance *Config
)

func ConfigGetting() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
		viper.AutomaticEnv()

		// ตอนจะ Deploy ขึ้น Could เราจะเป็น SERVER_PORT แทน server.port
		// Build once deploy multiple in different environment
		// server.port -> SERVER_PORT
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		// from config.yaml to json and map to struct

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}

		// Validate

		validating := validator.New()

		if err := validating.Struct(configInstance); err != nil {
			panic(err)
		}
	})
	return configInstance
}
