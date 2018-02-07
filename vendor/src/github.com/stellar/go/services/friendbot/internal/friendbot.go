package internal

import (
	"strconv"
	"sync"

	b "github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/support/errors"
)

// Bot represents the friendbot subsystem.
type Bot struct {
	Horizon         *horizon.Client
	Secret          string
	Network         string
	StartingBalance string

	sequence uint64
	lock     sync.Mutex
}

// Pay funds the account at `destAddress`
func (bot *Bot) Pay(destAddress string) (*horizon.TransactionSuccess, error) {
	err := bot.checkSequenceRefresh()
	if err != nil {
		return nil, err
	}

	signed, err := bot.makeTx(destAddress)
	if err != nil {
		return nil, err
	}

	result, err := bot.Horizon.SubmitTransaction(signed)
	if err != nil {
		switch e := err.(type) {
		case *horizon.Error:
			bot.checkHandleBadSequence(e)
		}
	}
	return &result, err
}

func (bot *Bot) checkHandleBadSequence(err *horizon.Error) {
	if err.Problem.Type != "tx_bad_seq" {
		return
	}

	// force refresh sequence for bad sequence errors
	bot.lock.Lock()
	defer bot.lock.Unlock()
	bot.refreshSequence()
}

// establish initial sequence if needed
func (bot *Bot) checkSequenceRefresh() error {
	if bot.sequence != 0 {
		return nil
	}

	bot.lock.Lock()
	defer bot.lock.Unlock()

	// short-circuit here if the thread that previously had the lock was successful in refreshing the sequence
	if bot.sequence != 0 {
		return nil
	}

	return bot.refreshSequence()
}

func (bot *Bot) makeTx(destAddress string) (string, error) {
	bot.lock.Lock()
	defer bot.lock.Unlock()

	txn, err := b.Transaction(
		b.SourceAccount{AddressOrSeed: bot.Secret},
		b.Sequence{Sequence: bot.sequence + 1},
		b.Network{Passphrase: bot.Network},
		b.CreateAccount(
			b.Destination{AddressOrSeed: destAddress},
			b.NativeAmount{Amount: bot.StartingBalance},
		),
	)

	if err != nil {
		return "", errors.Wrap(err, "Error building a transaction")
	}

	txs, err := txn.Sign(bot.Secret)
	if err != nil {
		return "", errors.Wrap(err, "Error signing a transaction")
	}

	base64, err := txs.Base64()

	// only increment the in-memory sequence number if we are going to submit the transaction, while we hold the lock
	if err == nil {
		bot.sequence++
	}
	return base64, err
}

// refreshes the sequence from the bot account
func (bot *Bot) refreshSequence() error {
	botAccount, err := bot.Horizon.LoadAccount(bot.address())
	if err != nil {
		bot.sequence = 0
		return err
	}

	seq, err := strconv.ParseInt(botAccount.Sequence, 10, 0)
	if err != nil {
		bot.sequence = 0
		return err
	}

	bot.sequence = uint64(seq)
	return nil
}

func (bot *Bot) address() string {
	kp := keypair.MustParse(bot.Secret)
	return kp.Address()
}
