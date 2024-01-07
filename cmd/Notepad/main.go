package main

import (
	"fmt"
	"net/http"
	"notepad/internal/config"
	"notepad/internal/storage/postgres"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	logger, err := setupZapLogger()
	defer logger.Sync()
	if err != nil {
		panic(err)
	}

	cfg, err := config.ReadConfig()
	if err != nil {
		logger.Fatal("can not read config:", zap.Error(err))
	}
	_ = cfg
	logger.Info("successfull config setup")

	db, err := postgres.New(cfg.DataBase.Host, cfg.DataBase.User, cfg.DataBase.Password, cfg.DataBase.Dbname, cfg.DataBase.Port)
	if err != nil {
		logger.Fatal("can not connect with database", zap.Error(err))
	}
	logger.Info("successfull database setup")
	_ = db

	mux := http.NewServeMux()
	_ = mux
	//Подключаем все обработчики ручек и эндпоинты
}

func setupZapLogger() (*zap.Logger, error) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	if err != nil {
		return nil, fmt.Errorf("can not setup logger: %w", err)
	}
	logger.Info("successfull logger setup")
	return logger, nil
}
