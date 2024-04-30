package config

import (
	"log"

	"strconv"

	env "github.com/Netflix/go-env"
	"github.com/behrouz-rfa/kentech/pkg/logger"
	"github.com/joho/godotenv"
)

const DbTimeout = 10

var conf = &Config{}

type Config struct {
	// Server config
	Server struct {
		Port int    `env:"SERVER_PORT"`
		Host string `env:"SERVER_PORT"`
	}

	// Database config
	Database struct {
		Host      string `env:"DB_HOST"`
		Port      int    `env:"DB_PORT"`
		User      string `env:"DB_USER"`
		Password  string `env:"DB_PASS"`
		Name      string `env:"DB_NAME"`
		SSL       string `env:"DB_SSL"`
		Clustered bool   `env:"DB_CLUSTERED"`
	}

	Jwt struct {
		Secret string `env:"JWT_KEY"`
	}

	// Env config
	Env struct {
		Env      string `env:"ENV"`
		LogLevel string `env:"LOG_LEVEL"`
	}

	// logging config
	logger *logger.Entry
}

func LoadTest() {
}

func Load() {
	lg := logger.General.Component("config")
	err := godotenv.Load(".env.local")
	if err != nil {
		err = godotenv.Load(".env")
		if err != nil {
			lg.Info("No .env file found, loading variables from environment")
		}
	}

	_, err = env.UnmarshalFromEnviron(conf)
	if err != nil {
		lg.WithError(err).Fatal("failed to unmarshal config")
	}

	conf.logger = lg
}

func LightLoad() {
	err := godotenv.Load(".env.local")
	if err != nil {
		err = godotenv.Load(".env")
		if err != nil {
			log.Println("No .env file found, loading variables from environment")
		}
	}

	_, err = env.UnmarshalFromEnviron(conf)
	if err != nil {
		log.Println("failed to unmarshal config")
	}
}

func Get() *Config {
	return conf
}

func (c *Config) DbConnectionString() string {
	authSegment := ""

	if c.Database.User != "" && c.Database.Password != "" {
		authSegment = c.Database.User + ":" + c.Database.Password + "@"
	}

	prefix := "mongodb://"

	if c.Database.Clustered {
		prefix = "mongodb+srv://"

		return prefix + authSegment + c.Database.Host
	}

	return prefix +
		authSegment +
		c.Database.Host +
		":" +
		strconv.Itoa(c.Database.Port)
}
