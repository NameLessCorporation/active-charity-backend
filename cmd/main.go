package main

import (
	"flag"
	"io/ioutil"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"

	"github.com/NameLessCorporation/active-charity-backend/app"
	"github.com/NameLessCorporation/active-charity-backend/config"
)

const defaultConfigPath = "./config/default-config.yaml"

var (
	configPath string
	certPath   string
)

func init() {
	flag.StringVar(&configPath, "config-path", "./config/default-config.yaml", "path to config")
	flag.StringVar(&certPath, "cert-path", "", "path to certificates for bd connection")
}

func main() {
	flag.Parse()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		logger.Fatal("failed config file reading`",
			zap.String("error", err.Error()),
		)
	}

	var config config.Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		logger.Fatal("failed config unmarshalling",
			zap.String("error", err.Error()),
		)
	}

	app := app.NewApp(logger, &config)
	err = app.StartApp(certPath)
	if err != nil {
		logger.Fatal("failed start app",
			zap.String("error", err.Error()),
		)
	}
}
