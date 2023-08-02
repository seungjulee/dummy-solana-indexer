package datastore

import (
	"github.com/seungjulee/dummy-solana-indexer/pkg/datastore/models"
	"github.com/seungjulee/dummy-solana-indexer/pkg/types"
)

type Datastore interface {
	SaveAccount(types.Account) error
	GetHigestTokenValueAccount() (models.Account, error)
}