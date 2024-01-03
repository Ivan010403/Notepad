package main

import (
	"fmt"
	"net/http"
	"notepad/internal/config"
	"notepad/internal/storage/postgres"

	_ "github.com/lib/pq"
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

	db, err := postgres.New()
	if err != nil {
		logger.Fatal("can not connect with database", zap.Error(err))
	}
	_ = db
	//Подключаем все обработчики ручек и эндпоинты

	srv := &http.Server{
		Addr: cfg.Address,
		// Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	logger.Info("http server started")

	//TODO: graceful shutdown
	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal("server stopped")
	}
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
