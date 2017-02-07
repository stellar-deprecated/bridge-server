package submitter

import (
	"errors"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stellar/gateway/db/entities"
	"github.com/stellar/gateway/horizon"
	"github.com/stellar/gateway/mocks"
	b "github.com/stellar/go/build"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTransactionSubmitter(t *testing.T) {
	mockHorizon := new(mocks.MockHorizon)
	mockEntityManager := new(mocks.MockEntityManager)
	mocks.PredefinedTime = time.Now()

	Convey("TransactionSubmitter", t, func() {
		seed := "SDZT3EJZ7FZRYNTLOZ7VH6G5UYBFO2IO3Q5PGONMILPCZU3AL7QNZHTE"
		accountID := "GCLOMB72ODBFUGK4E2BK7VMR3RNZ5WSTMEOGNA2YUVHFR3WMH2XBAB6H"

		Convey("LoadAccount", func() {
			transactionSubmitter := NewTransactionSubmitter(
				mockHorizon,
				mockEntityManager,
				"Test SDF Network ; September 2015",
				mocks.Now,
			)

			Convey("When seed is invalid", func() {
				_, err := transactionSubmitter.LoadAccount("invalidSeed")
				assert.NotNil(t, err)
			})

			Convey("When there is a problem loading an account", func() {
				mockHorizon.On(
					"LoadAccount",
					accountID,
				).Return(
					horizon.AccountResponse{},
					errors.New("Account not found"),
				).Once()

				_, err := transactionSubmitter.LoadAccount(seed)
				assert.NotNil(t, err)
				mockHorizon.AssertExpectations(t)
			})

			Convey("Successfully loads an account", func() {
				mockHorizon.On(
					"LoadAccount",
					accountID,
				).Return(
					horizon.AccountResponse{
						AccountID:      accountID,
						SequenceNumber: "10372672437354496",
					},
					nil,
				).Once()

				account, err := transactionSubmitter.LoadAccount(seed)
				assert.Nil(t, err)
				assert.Equal(t, account.Keypair.Address(), accountID)
				assert.Equal(t, account.Seed, seed)
				assert.Equal(t, account.SequenceNumber, uint64(10372672437354496))
				mockHorizon.AssertExpectations(t)
			})
		})

		Convey("SubmitTransaction", func() {
			Convey("Submits transaction without a memo", func() {
				operation := b.Payment(
					b.Destination{"GB3W7VQ2A2IOQIS4LUFUMRC2DWXONUDH24ROLE6RS4NGUNHVSXKCABOM"},
					b.NativeAmount{"100"},
				)

				Convey("Error response from horizon", func() {
					transactionSubmitter := NewTransactionSubmitter(
						mockHorizon,
						mockEntityManager,
						"Test SDF Network ; September 2015",
						mocks.Now,
					)

					mockHorizon.On(
						"LoadAccount",
						accountID,
					).Return(
						horizon.AccountResponse{
							AccountID:      accountID,
							SequenceNumber: "10372672437354496",
						},
						nil,
					).Once()

					err := transactionSubmitter.InitAccount(seed)
					assert.Nil(t, err)

					txB64 := "AAAAAJbmB/pwwloZXCaCr9WR3Fue2lNhHGaDWKVOWO7MPq4QAAAAZAAk2eQAAAABAAAAAAAAAAAAAAABAAAAAAAAAAEAAAAAd2/WGgaQ6CJcXQtGRFodrubQZ9ci5ZPRlxpqNPWV1CAAAAAAAAAAADuaygAAAAAAAAAAAcw+rhAAAABAyFjIMIZOtstCWtZlVBDj1AhTmsk5v1i2GGY4by2b5mgZoXXGgFTB8sfbQav0LzFKCcxY8h+9xPMT2e9xznAfDw=="

					// Persist sending transaction
					mockEntityManager.On(
						"Persist",
						mock.AnythingOfType("*entities.SentTransaction"),
					).Return(nil).Once().Run(func(args mock.Arguments) {
						transaction := args.Get(0).(*entities.SentTransaction)
						assert.Equal(t, "4f885999be6ea7891052a53e496bcfb5c5a1a5bfb31923f649b028fdc74dd050", transaction.TransactionID)
						assert.Equal(t, "sending", string(transaction.Status))
						assert.Equal(t, "GCLOMB72ODBFUGK4E2BK7VMR3RNZ5WSTMEOGNA2YUVHFR3WMH2XBAB6H", transaction.Source)
						assert.Equal(t, mocks.PredefinedTime, transaction.SubmittedAt)
						assert.Equal(t, txB64, transaction.EnvelopeXdr)
					})

					// Persist failure
					mockEntityManager.On(
						"Persist",
						mock.AnythingOfType("*entities.SentTransaction"),
					).Return(nil).Once().Run(func(args mock.Arguments) {
						transaction := args.Get(0).(*entities.SentTransaction)
						assert.Equal(t, "4f885999be6ea7891052a53e496bcfb5c5a1a5bfb31923f649b028fdc74dd050", transaction.TransactionID)
						assert.Equal(t, "failure", string(transaction.Status))
						assert.Equal(t, "GCLOMB72ODBFUGK4E2BK7VMR3RNZ5WSTMEOGNA2YUVHFR3WMH2XBAB6H", transaction.Source)
						assert.Equal(t, mocks.PredefinedTime, transaction.SubmittedAt)
						assert.Equal(t, txB64, transaction.EnvelopeXdr)
					})

					mockHorizon.On("SubmitTransaction", txB64).Return(
						horizon.SubmitTransactionResponse{
							Ledger: nil,
							Extras: &horizon.SubmitTransactionResponseExtras{
								ResultXdr: "AAAAAAAAAGT/////AAAAAQAAAAAAAAAB////+wAAAAA=", // no_destination
							},
						},
						nil,
					).Once()

					_, err = transactionSubmitter.SubmitTransaction(seed, operation, nil)
					assert.Nil(t, err)
					mockHorizon.AssertExpectations(t)
				})

				Convey("Bad Sequence response from horizon", func() {
					transactionSubmitter := NewTransactionSubmitter(
						mockHorizon,
						mockEntityManager,
						"Test SDF Network ; September 2015",
						mocks.Now,
					)

					mockHorizon.On(
						"LoadAccount",
						accountID,
					).Return(
						horizon.AccountResponse{
							AccountID:      accountID,
							SequenceNumber: "10372672437354496",
						},
						nil,
					).Once()

					err := transactionSubmitter.InitAccount(seed)
					assert.Nil(t, err)

					txB64 := "AAAAAJbmB/pwwloZXCaCr9WR3Fue2lNhHGaDWKVOWO7MPq4QAAAAZAAk2eQAAAABAAAAAAAAAAAAAAABAAAAAAAAAAEAAAAAd2/WGgaQ6CJcXQtGRFodrubQZ9ci5ZPRlxpqNPWV1CAAAAAAAAAAADuaygAAAAAAAAAAAcw+rhAAAABAyFjIMIZOtstCWtZlVBDj1AhTmsk5v1i2GGY4by2b5mgZoXXGgFTB8sfbQav0LzFKCcxY8h+9xPMT2e9xznAfDw=="

					// Persist sending transaction
					mockEntityManager.On(
						"Persist",
						mock.AnythingOfType("*entities.SentTransaction"),
					).Return(nil).Once().Run(func(args mock.Arguments) {
						transaction := args.Get(0).(*entities.SentTransaction)
						assert.Equal(t, "4f885999be6ea7891052a53e496bcfb5c5a1a5bfb31923f649b028fdc74dd050", transaction.TransactionID)
						assert.Equal(t, "sending", string(transaction.Status))
						assert.Equal(t, "GCLOMB72ODBFUGK4E2BK7VMR3RNZ5WSTMEOGNA2YUVHFR3WMH2XBAB6H", transaction.Source)
						assert.Equal(t, mocks.PredefinedTime, transaction.SubmittedAt)
						assert.Equal(t, txB64, transaction.EnvelopeXdr)
					})

					// Persist failure
					mockEntityManager.On(
						"Persist",
						mock.AnythingOfType("*entities.SentTransaction"),
					).Return(nil).Once().Run(func(args mock.Arguments) {
						transaction := args.Get(0).(*entities.SentTransaction)
						assert.Equal(t, "4f885999be6ea7891052a53e496bcfb5c5a1a5bfb31923f649b028fdc74dd050", transaction.TransactionID)
						assert.Equal(t, "failure", string(transaction.Status))
						assert.Equal(t, "GCLOMB72ODBFUGK4E2BK7VMR3RNZ5WSTMEOGNA2YUVHFR3WMH2XBAB6H", transaction.Source)
						assert.Equal(t, mocks.PredefinedTime, transaction.SubmittedAt)
						assert.Equal(t, txB64, transaction.EnvelopeXdr)
					})

					mockHorizon.On("SubmitTransaction", txB64).Return(
						horizon.SubmitTransactionResponse{
							Ledger: nil,
							Extras: &horizon.SubmitTransactionResponseExtras{
								ResultXdr: "AAAAAAAAAAD////7AAAAAA==", // tx_bad_seq
							},
						},
						nil,
					).Once()

					// Updating sequence number
					mockHorizon.On(
						"LoadAccount",
						accountID,
					).Return(
						horizon.AccountResponse{
							AccountID:      accountID,
							SequenceNumber: "100",
						},
						nil,
					).Once()

					_, err = transactionSubmitter.SubmitTransaction(seed, operation, nil)
					assert.Nil(t, err)
					assert.Equal(t, uint64(100), transactionSubmitter.Accounts[seed].SequenceNumber)
					mockHorizon.AssertExpectations(t)
				})

				Convey("Successfully submits a transaction", func() {
					transactionSubmitter := NewTransactionSubmitter(
						mockHorizon,
						mockEntityManager,
						"Test SDF Network ; September 2015",
						mocks.Now,
					)

					mockHorizon.On(
						"LoadAccount",
						accountID,
					).Return(
						horizon.AccountResponse{
							AccountID:      accountID,
							SequenceNumber: "10372672437354496",
						},
						nil,
					).Once()

					err := transactionSubmitter.InitAccount(seed)
					assert.Nil(t, err)

					txB64 := "AAAAAJbmB/pwwloZXCaCr9WR3Fue2lNhHGaDWKVOWO7MPq4QAAAAZAAk2eQAAAABAAAAAAAAAAAAAAABAAAAAAAAAAEAAAAAd2/WGgaQ6CJcXQtGRFodrubQZ9ci5ZPRlxpqNPWV1CAAAAAAAAAAADuaygAAAAAAAAAAAcw+rhAAAABAyFjIMIZOtstCWtZlVBDj1AhTmsk5v1i2GGY4by2b5mgZoXXGgFTB8sfbQav0LzFKCcxY8h+9xPMT2e9xznAfDw=="

					// Persist sending transaction
					mockEntityManager.On(
						"Persist",
						mock.AnythingOfType("*entities.SentTransaction"),
					).Return(nil).Once().Run(func(args mock.Arguments) {
						transaction := args.Get(0).(*entities.SentTransaction)
						assert.Equal(t, "4f885999be6ea7891052a53e496bcfb5c5a1a5bfb31923f649b028fdc74dd050", transaction.TransactionID)
						assert.Equal(t, "sending", string(transaction.Status))
						assert.Equal(t, "GCLOMB72ODBFUGK4E2BK7VMR3RNZ5WSTMEOGNA2YUVHFR3WMH2XBAB6H", transaction.Source)
						assert.Equal(t, mocks.PredefinedTime, transaction.SubmittedAt)
						assert.Equal(t, txB64, transaction.EnvelopeXdr)
					})

					// Persist failure
					mockEntityManager.On(
						"Persist",
						mock.AnythingOfType("*entities.SentTransaction"),
					).Return(nil).Once().Run(func(args mock.Arguments) {
						transaction := args.Get(0).(*entities.SentTransaction)
						assert.Equal(t, "4f885999be6ea7891052a53e496bcfb5c5a1a5bfb31923f649b028fdc74dd050", transaction.TransactionID)
						assert.Equal(t, "success", string(transaction.Status))
						assert.Equal(t, "GCLOMB72ODBFUGK4E2BK7VMR3RNZ5WSTMEOGNA2YUVHFR3WMH2XBAB6H", transaction.Source)
						assert.Equal(t, mocks.PredefinedTime, transaction.SubmittedAt)
						assert.Equal(t, txB64, transaction.EnvelopeXdr)
					})

					ledger := uint64(1486276)
					mockHorizon.On("SubmitTransaction", txB64).Return(
						horizon.SubmitTransactionResponse{Ledger: &ledger},
						nil,
					).Once()

					response, err := transactionSubmitter.SubmitTransaction(seed, operation, nil)
					assert.Nil(t, err)
					assert.Equal(t, *response.Ledger, ledger)
					assert.Equal(t, uint64(10372672437354497), transactionSubmitter.Accounts[seed].SequenceNumber)
					mockHorizon.AssertExpectations(t)
				})
			})

			Convey("Submits transaction with a memo", func() {
				operation := b.Payment(
					b.Destination{"GB3W7VQ2A2IOQIS4LUFUMRC2DWXONUDH24ROLE6RS4NGUNHVSXKCABOM"},
					b.NativeAmount{"100"},
				)

				memo := b.MemoText{"Testing!"}

				Convey("Successfully submits a transaction", func() {
					transactionSubmitter := NewTransactionSubmitter(
						mockHorizon,
						mockEntityManager,
						"Test SDF Network ; September 2015",
						mocks.Now,
					)

					mockHorizon.On(
						"LoadAccount",
						accountID,
					).Return(
						horizon.AccountResponse{
							AccountID:      accountID,
							SequenceNumber: "10372672437354496",
						},
						nil,
					).Once()

					err := transactionSubmitter.InitAccount(seed)
					assert.Nil(t, err)

					txB64 := "AAAAAJbmB/pwwloZXCaCr9WR3Fue2lNhHGaDWKVOWO7MPq4QAAAAZAAk2eQAAAABAAAAAAAAAAEAAAAIVGVzdGluZyEAAAABAAAAAAAAAAEAAAAAd2/WGgaQ6CJcXQtGRFodrubQZ9ci5ZPRlxpqNPWV1CAAAAAAAAAAADuaygAAAAAAAAAAAcw+rhAAAABAU5ahFsd28sVKSUFcmAiEf+zSLXhf9HG/pJuQirR0s43zs7Y43vM8T3sIvJWHgwMADaZiy/D+evYWd/vS/uO8Ag=="

					// Persist sending transaction
					mockEntityManager.On(
						"Persist",
						mock.AnythingOfType("*entities.SentTransaction"),
					).Return(nil).Once().Run(func(args mock.Arguments) {
						transaction := args.Get(0).(*entities.SentTransaction)
						assert.Equal(t, "60cb3c020b0c97352cbabdf68a822b04baea61927b0f1ac31260a9f8d0150316", transaction.TransactionID)
						assert.Equal(t, "sending", string(transaction.Status))
						assert.Equal(t, "GCLOMB72ODBFUGK4E2BK7VMR3RNZ5WSTMEOGNA2YUVHFR3WMH2XBAB6H", transaction.Source)
						assert.Equal(t, mocks.PredefinedTime, transaction.SubmittedAt)
						assert.Equal(t, txB64, transaction.EnvelopeXdr)
					})

					// Persist failure
					mockEntityManager.On(
						"Persist",
						mock.AnythingOfType("*entities.SentTransaction"),
					).Return(nil).Once().Run(func(args mock.Arguments) {
						transaction := args.Get(0).(*entities.SentTransaction)
						assert.Equal(t, "60cb3c020b0c97352cbabdf68a822b04baea61927b0f1ac31260a9f8d0150316", transaction.TransactionID)
						assert.Equal(t, "success", string(transaction.Status))
						assert.Equal(t, "GCLOMB72ODBFUGK4E2BK7VMR3RNZ5WSTMEOGNA2YUVHFR3WMH2XBAB6H", transaction.Source)
						assert.Equal(t, mocks.PredefinedTime, transaction.SubmittedAt)
						assert.Equal(t, txB64, transaction.EnvelopeXdr)
					})

					ledger := uint64(1486276)
					mockHorizon.On("SubmitTransaction", txB64).Return(
						horizon.SubmitTransactionResponse{Ledger: &ledger},
						nil,
					).Once()

					response, err := transactionSubmitter.SubmitTransaction(seed, operation, memo)
					assert.Nil(t, err)
					assert.Equal(t, *response.Ledger, ledger)
					assert.Equal(t, uint64(10372672437354497), transactionSubmitter.Accounts[seed].SequenceNumber)
					mockHorizon.AssertExpectations(t)
				})
			})
		})
	})
}
