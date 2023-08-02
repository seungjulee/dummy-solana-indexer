package indexer

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/seungjulee/dummy-solana-indexer/pkg/logger"
	"github.com/seungjulee/dummy-solana-indexer/pkg/types"
)

func loadFakeAccount(path string) ([]types.Account, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	logger.Info("Successfully loaded account inputs")

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var accounts []types.Account
	json.Unmarshal(byteValue, &accounts)

	return accounts, nil
}

func fetchAccountResponseGenerator(accounts []types.Account) <-chan types.Account {
	ch := make(chan types.Account)
	go func() {
		for _, ra := range accounts {
			// Each account comes into the system at a continuous uniform (random) distribution between 0 and 1000ms
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			ch <- ra
		}
		close(ch)
	}()
	return ch
}