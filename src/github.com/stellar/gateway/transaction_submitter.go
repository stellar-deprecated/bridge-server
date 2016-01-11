package gateway

import (
	"errors"
	"log"

	"github.com/stellar/gateway/horizon"
	"github.com/stellar/go-stellar-base/build"
	"github.com/stellar/go-stellar-base/keypair"
)

type TransactionSubmitter struct {
	Horizon           *horizon.Horizon
	AvailableChannels chan ChannelAccount
}

type ChannelAccount struct {
	Keypair        keypair.KP
	Seed           string
	SequenceNumber uint64
}

func NewTransactionSubmitter(horizon *horizon.Horizon, channelsSeeds []string) (ts TransactionSubmitter, err error) {
	ts.Horizon = horizon

	newChannels := make(chan ChannelAccount, len(channelsSeeds))
	errorChannel := make(chan error)

	for _, channelSeed := range channelsSeeds {
		channelSeed := channelSeed // Create new instance for the goroutine.
		go func() {
			channelKeypair, err := keypair.Parse(channelSeed)
			if err != nil {
				errorChannel <- err
			}
			channelAccount, err := horizon.LoadAccount(channelKeypair.Address())
			if err != nil {
				errorChannel <- err
			}

			newChannels <- ChannelAccount{
				channelKeypair,
				channelSeed,
				channelAccount.SequenceNumber,
			}
		}()
	}

	ts.AvailableChannels = make(chan ChannelAccount, len(channelsSeeds))

	for i := 0; i < len(channelsSeeds); i++ {
		select {
		case ch := <-newChannels:
			ts.AvailableChannels <- ch
		case err = <-errorChannel:
			return
		}
	}

	return
}

func (ts *TransactionSubmitter) SubmitTransaction(opSourceSeed string, operation interface{}) (response horizon.SubmitTransactionResponse, err error) {
	channelAccount := <-ts.AvailableChannels

	_, err = keypair.Parse(opSourceSeed)
	if err != nil {
		log.Print("Invalid opSourceSeed")
		return
	}

	mutator, ok := operation.(build.TransactionMutator)
	if !ok {
		log.Print("Cannot cast to build.TransactionMutator")
		err = errors.New("Cannot cast to build.TransactionMutator")
		return
	}

	tx := build.Transaction(
		build.SourceAccount{channelAccount.Seed},
		build.Sequence{channelAccount.SequenceNumber + 1},
		mutator,
	)

	signersSeeds := []string{opSourceSeed}
	if channelAccount.Seed != opSourceSeed {
		// We're using channels
		signersSeeds = append(signersSeeds, channelAccount.Seed)
	}
	txe := tx.Sign(signersSeeds...)
	txeB64, err := txe.Base64()

	if err != nil {
		log.Print("Cannot encode transaction envelope ", err)
		return
	}

	response, err = ts.Horizon.SubmitTransaction(txeB64)
	if err != nil {
		log.Print("Error submitting transaction ", err)
		// Sync sequence number
		// TODO do it only if it's BAD_SEQ error
		accountResponse, _ := ts.Horizon.LoadAccount(channelAccount.Keypair.Address())
		channelAccount.SequenceNumber = accountResponse.SequenceNumber
		return
	} else {
		channelAccount.SequenceNumber++
	}

	ts.AvailableChannels <- channelAccount
	return
}
