package models

import (
	"encoding/json"

	"github.com/seungjulee/dummy-solana-indexer/pkg/types"
)

type Account struct {
	AccountId string `json:"id"`
	AccountType string `json:"accountType"`
	Tokens int `json:"tokens" gorm:"index:,sort:desc"`
	CallbackTimeMS int `json:"callbackTimeMS"`
	Version int `json:"version"`
	Data string `json:"data"`
}

func ConvertAccount(account types.Account) (Account, error) {
	jsonBytes, err := json.Marshal(account.Data)
	if err != nil {
		return Account{}, err
	}
	strData := string(jsonBytes)

	return Account{
		AccountId: account.Id,
		AccountType: string(account.AccountType),
		Tokens: account.Tokens,
		CallbackTimeMS: account.CallbackTimeMS,
		Version: account.Version,
		Data: strData,
	}, nil
}