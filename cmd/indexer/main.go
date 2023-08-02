package main

import (
	"github.com/seungjulee/dummy-solana-indexer/pkg/datastore"
	"github.com/seungjulee/dummy-solana-indexer/pkg/indexer"
	"github.com/seungjulee/dummy-solana-indexer/pkg/logger"
	"go.uber.org/zap"
)

type Config struct {
	IndexerConfig *indexer.Config
	SqliteConfig *datastore.SqliteConfig
}

func main() {
	cfg := &Config{
		IndexerConfig: &indexer.Config{
			FakeAccountPath: "./input.json",
		},
		SqliteConfig: &datastore.SqliteConfig{
			SqlitePath: "./test.db",
		},
	}

	logger.Info("initializing app with config", zap.Any("config", cfg))

	db, err := datastore.NewSqllite(cfg.SqliteConfig)
	if err != nil {
		logger.Fatal(err.Error())
		panic(err)
	}

	a := indexer.New(cfg.IndexerConfig, db)
	err = a.SchedulePeriodicIndex()
	if err != nil {
		logger.Fatal(err.Error())
		panic(err)
	}
}