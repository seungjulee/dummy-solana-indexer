package models

import (
	"encoding/json"
	"time"

	"github.com/seungjulee/fake-solana-indexer/pkg/types"
	"gorm.io/gorm"
)

type Account struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Id string `json:"id" gorm:"index:idx_id_version,priority:2"`
	AccountType string `json:"accountType"`
	Tokens int `json:"tokens"`
	CallbackTimeMS int `json:"callbackTimeMS"`
	Version int `json:"version" gorm:"index:idx_id_version,priority:1"`
	Data string `json:"data"`
}

func ConvertAccount(account types.Account) (Account, error) {
	jsonBytes, err := json.Marshal(account.Data)
	if err != nil {
		return Account{}, err
	}
	strData := string(jsonBytes)

	return Account{
		Id: account.Id,
		AccountType: string(account.AccountType),
		Tokens: account.Tokens,
		CallbackTimeMS: account.CallbackTimeMS,
		Version: account.Version,
		Data: strData,
	}, nil
}