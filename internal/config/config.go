package config

import (
	env "github.com/fajarardiyanto/flt-go-env/lib"
	interfaceClient "github.com/fajarardiyanto/flt-go-listener/interfaces"
	"github.com/fajarardiyanto/flt-go-logger/interfaces"
	log "github.com/fajarardiyanto/flt-go-logger/lib"
	interfaceJaeger "github.com/fajarardiyanto/flt-go-tracer/interfaces"
	"github.com/pkg/errors"
	"os"
)

var (
	config Config
	logger interfaces.Logger
)

type Config struct {
	Client interfaceClient.Client     `json:"client"`
	Server interfaceClient.ClientRest `json:"server"`
}

func GetLogger() interfaces.Logger {
	return logger
}

func GetConfig() *Config {
	return &config
}

func init() {
	if err := env.LoadEnv(".env"); err != nil {
		causer := errors.Cause(err)
		if os.IsNotExist(causer) {
			GetLogger().Info("Using default env config")
		} else {
			GetLogger().Error(causer).Quit()
		}
	}

	GetConfig().Client = interfaceClient.Client{
		Name:    env.EnvString("APPLICATION_NAME", "AFAIK Client Rest"),
		Host:    env.EnvString("CLIENT_HOST", "127.0.0.1"),
		Port:    env.EnvInt("CLIENT_PORT", 8801),
		Timeout: env.EnvInt("CLIENT_TIMEOUT", 3000),
	}

	GetConfig().Server = interfaceClient.ClientRest{
		Name:    env.EnvString("APPLICATION_NAME", "AFAIK Client Rest"),
		Host:    env.EnvString("LISTENER_HOST", "127.0.0.1"),
		Port:    env.EnvInt("LISTENER_PORT", 8082),
		Timeout: env.EnvInt("LISTENER_TIMEOUT", 10), // Second
		Jaeger: interfaceJaeger.JaegerConfig{
			Host:   env.EnvString("JAEGER_HOST", "0.0.0.0"),
			Port:   env.EnvString("JAEGER_PORT", "6831"),
			Enable: env.EnvBool("JAEGER_ENABLE", false),
		},
	}
}

func Init() {
	logger = log.NewLib()
	logger.Init("News Client")
}
