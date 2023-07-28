package indexer

import (
	"context"
	"math/rand"
	"time"

	"github.com/seungjulee/fake-solana-indexer/pkg/datastore"
	"github.com/seungjulee/fake-solana-indexer/pkg/logger"
	"github.com/seungjulee/fake-solana-indexer/pkg/types"
	"go.uber.org/zap"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Indexer interface {
	SchedulePeriodicIndex() error
}

type indexer struct {
	cfg *Config
	db datastore.Datastore
	latestVersionByAccount map[string]int
	cancelCallbackFnByAccount map[string]context.CancelFunc
}

type Config struct {
	FakeAccountPath string
}

func New(cfg *Config, db datastore.Datastore) Indexer {
	return &indexer{
		cfg: cfg,
		db: db,
		latestVersionByAccount: make(map[string]int),
		cancelCallbackFnByAccount: make(map[string]context.CancelFunc),
	}
}

func (a *indexer) SchedulePeriodicIndex() error {
	logger.Info("Scheduling periodic index")

	accounts, err := loadFakeAccount("./input.json")
	if err != nil {
		return err
	}
	for account := range fetchAccountResponseGenerator(accounts) {
		if err != nil {
			return err
		}
		a.db.SaveAccount(account)
		logger.Info("indexed account into db", zap.String("id", account.Id), zap.Int("version", account.Version))

		// Check the last version exists
		if lastVer, ok := a.latestVersionByAccount[account.Id]; ok && lastVer < account.Version {
			if cancelCB, ok := a.cancelCallbackFnByAccount[account.Id]; ok {
				// Attempt cancel. If the callback time has already been passed, this is no-op as `logCallback()`
				// already finished running
				cancelCB()
			}
		}

		// Create a new context
		ctx := context.Background()
		// Create a new context, with its cancellation function
		ctx, cancel := context.WithCancel(ctx)
		go logCallback(ctx, account)

		// Save the latest version by account id
		a.latestVersionByAccount[account.Id] = account.Version
		// Save cancel callback by account id
		a.cancelCallbackFnByAccount[account.Id] = cancel
	}

	// wait enough time to make sure that remaining callback gets fired
	<-time.After(10 * time.Second)



	return nil
}

// logCallback prints the callback log if it's within the callback time. Otherwise, it
// prints the callback cancell log
func logCallback(ctx context.Context, account types.Account) {
	// only one of the following function will execute, whichever comes first.
	select {
	// wait for the account's callback time
	case <-time.After(time.Duration(account.CallbackTimeMS) * time.Millisecond):
		logger.Info("callback fires", zap.String("id", account.Id), zap.Int("version", account.Version))
	// wait until cancel is called
	case <-ctx.Done():
		logger.Info("callback cancelled", zap.String("id", account.Id), zap.Int("version", account.Version))
	}
}


func (a *indexer) IndexAccount(ctx context.Context) error {
	return nil
}
