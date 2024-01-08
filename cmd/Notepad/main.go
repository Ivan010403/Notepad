package main

import (
	"fmt"
	"net/http"
	"notepad/internal/config"
	"notepad/internal/handlers/delete"
	"notepad/internal/handlers/new"
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
		logger.Fatal("can not read config --->", zap.Error(err))
	}
	logger.Info("successfull config setup")

	db, err := postgres.New(cfg.DataBase.Host, cfg.DataBase.User, cfg.DataBase.Password, cfg.DataBase.Dbname, cfg.DataBase.Port)
	if err != nil {
		logger.Fatal("can not connect with database --->", zap.Error(err))
	}
	logger.Info("successfull database setup")
	_ = db

	mux := http.NewServeMux()
	mux.HandleFunc("/new", new.New(logger, db))
	mux.HandleFunc("/delete", delete.Delete(logger, db))

	srv := http.Server{
		Addr:         cfg.Address,
		Handler:      mux,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal("fatal in server --->", zap.Error(err))
	}
}

func setupZapLogger() (*zap.Logger, error) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	if err != nil {
		return nil, fmt.Errorf("can not setup logger ---> %w", err)
	}
	logger.Info("successfull logger setup")
	return logger, nil
}
