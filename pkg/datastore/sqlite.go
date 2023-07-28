package datastore

import (
	"github.com/seungjulee/fake-solana-indexer/pkg/datastore/models"
	"github.com/seungjulee/fake-solana-indexer/pkg/logger"
	"github.com/seungjulee/fake-solana-indexer/pkg/types"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	ormLogger "gorm.io/gorm/logger"
)

type SqliteConfig struct {
	SqlitePath string
}

func NewSqllite(cfg *SqliteConfig) (Datastore, error) {
	db, err := gorm.Open(sqlite.Open(cfg.SqlitePath), &gorm.Config{
		Logger: ormLogger.Default.LogMode(ormLogger.Silent),
	  })
	if err != nil {
	  return nil, err
	}
	logger.Info("successfully connected to the db")

	// Migrate the schema
	logger.Info("migrate the schema tables")
	db.AutoMigrate(&models.Account{})

	return &sqliteDB{
		db: db,
	}, nil
}

type sqliteDB struct {
	db *gorm.DB
}

func (s *sqliteDB) SaveAccount(tAccount types.Account) error {
	account, err := models.ConvertAccount(tAccount)
	if err != nil {
		return err
	}
	return s.db.Create(&account).Error
}

func (s *sqliteDB) GetHigestTokenValueAccount() (models.Account, error) {

// 	SELECT
// 	account_id, account_type, tokens, MAX(version)
// FROM
// 	accounts
// GROUP BY
// 	account_id
// ORDER BY
// 	tokens DESC;

	var account models.Account
	if err := s.db.Model(&models.Account{}).Select("account_id, account_type, tokens, max(version) as latest_version").Group("account_id").Order("tokens desc").First(&account).Error; err != nil {
		return models.Account{}, err
	}
	return account, nil
}