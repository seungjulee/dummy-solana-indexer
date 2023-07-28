package datastore

import (
	"github.com/seungjulee/fake-solana-indexer/pkg/types"
)

type Datastore interface {
	SaveAccount(types.Account) error
}