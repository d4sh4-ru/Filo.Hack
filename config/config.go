package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config Конфигурация API сервера
type Config struct {
	Env        string `yaml:"env" env-required:"true"`
	Postgres   `yaml:"postgres" env-required:"true"`
	HTTPServer `yaml:"http_server" env-required:"true"`
}

// Postgres Конфигурация базы данных.
type Postgres struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     string `yaml:"port" env-required:"true"`
	User     string `yaml:"username" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	Database string `yaml:"database" env-required:"true"`
	Schema   string `yaml:"schema"`
}

// HTTPServer Конфигурация сервера HTTP.
type HTTPServer struct {
	Host         string        `yaml:"host" env-required:"true"`
	Port         string        `yaml:"port" env-required:"true"`
	ReadTimeout  time.Duration `yaml:"read_timeout" env-required:"true"`
	WriteTimeout time.Duration `yaml:"write_timeout" env-required:"true"`
	IdleTimeout  time.Duration `yaml:"idle_timeout" env-required:"true"`
	JWTSecret    string        `yaml:"jwt_secret" env-required:"true"`
}

// MustLoad - Загрузка и проверка конфигурации API сервера.
// В случае ошибки загрузки или проверки завершает работу процесса с ошибкой.
func MustLoad(configPath string) *Config {
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		log.Fatalf("config file %s not found", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("config file %s invalid: %v", configPath, err)
	}
	return &cfg
}
