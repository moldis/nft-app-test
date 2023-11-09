package database

import (
	"artemb/nft/pkg/config"
	"artemb/nft/pkg/db/model"
	"fmt"
	"time"

	"github.com/glebarez/sqlite"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

func NewInMemoryDB(cfg *config.DB) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(cfg.DSN), &gorm.Config{
		Logger: newLogger(),
	})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database %s", err.Error()))
	}

	err = db.AutoMigrate(&model.Collections{}, &model.MintedCollections{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func newLogger() zapgorm2.Logger {
	logLevel := logger.Info
	switch zap.L().Level() {
	case zap.DebugLevel, zap.InfoLevel:
		logLevel = logger.Info
	case zap.WarnLevel:
		logLevel = logger.Warn
	case zap.ErrorLevel, zap.FatalLevel, zap.PanicLevel:
		fallthrough
	default:
		logLevel = logger.Error
	}

	gormLogger := zapgorm2.Logger{
		ZapLogger:                 zap.L(),
		LogLevel:                  logLevel,
		SlowThreshold:             5000 * time.Millisecond,
		SkipCallerLookup:          false,
		IgnoreRecordNotFoundError: false,
		Context:                   nil,
	}
	gormLogger.SetAsDefault()
	return gormLogger
}
