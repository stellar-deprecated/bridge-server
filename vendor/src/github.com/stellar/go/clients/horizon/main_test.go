package horizon

import (
	"fmt"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stellar/go/support/errors"
	"github.com/stellar/go/support/http/httptest"
	"golang.org/x/net/context"
)

func ExampleClient_StreamLedgers() {
	client := DefaultPublicNetClient
	cursor := Cursor("now")

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		// Stop streaming after 60 seconds.
		time.Sleep(60 * time.Second)
		cancel()
	}()

	err := client.StreamLedgers(ctx, &cursor, func(l Ledger) {
		fmt.Println(l.Sequence)
	})

	if err != nil {
		fmt.Println(err)
	}
}

func ExampleClient_SubmitTransaction() {
	client := DefaultPublicNetClient
	transactionEnvelopeXdr := "AAAAABSxFjMo7qcQlJBlrZQypSqYsHA5hHaYxk5hFXwiehh6AAAAZAAIdakAAABZAAAAAAAAAAAAAAABAAAAAAAAAAEAAAAAFLEWMyjupxCUkGWtlDKlKpiwcDmEdpjGTmEVfCJ6GHoAAAAAAAAAAACYloAAAAAAAAAAASJ6GHoAAABAp0FnKOQ9lJPDXPTh/a91xoZ8BaznwLj59sdDGK94eGzCOk7oetw7Yw50yOSZg2mqXAST6Agc9Ao/f5T9gB+GCw=="

	response, err := client.SubmitTransaction(transactionEnvelopeXdr)
	if err != nil {
		fmt.Println(err)
		herr, isHorizonError := err.(*Error)
		if isHorizonError {
			resultCodes, err := herr.ResultCodes()
			if err != nil {
				fmt.Println("failed to extract result codes from horizon response")
				return
			}
			fmt.Println(resultCodes)
		}
		return
	}

	fmt.Println("Success")
	fmt.Println(response)
}

func TestHorizon(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Package: github.com/stellar/go/horizon")
}

var _ = Describe("Horizon", func() {
	var (
		client *Client
		hmock  *httptest.Client
	)

	BeforeEach(func() {
		hmock = httptest.NewClient()
		client = &Client{
			URL:  "https://localhost",
			HTTP: hmock,
		}
	})

	Describe("LoadAccount", func() {
		It("success response", func() {
			hmock.On(
				"GET",
				"https://localhost/accounts/GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H",
			).ReturnString(200, accountResponse)

			account, err := client.LoadAccount("GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H")
			Expect(err).To(BeNil())
			Expect(account.ID).To(Equal("GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H"))
			Expect(account.PT).To(Equal("1"))
			Expect(account.Signers[0].Key).To(Equal("XBT5HNPK6DAL6222MAWTLHNOZSDKPJ2AKNEQ5Q324CHHCNQFQ7EHBHZN"))
			Expect(account.Signers[0].Type).To(Equal("sha256_hash"))
			Expect(account.Data["test"]).To(Equal("R0NCVkwzU1FGRVZLUkxQNkFKNDdVS0tXWUVCWTQ1V0hBSkhDRVpLVldNVEdNQ1Q0SDROS1FZTEg="))
			Expect(account.GetNativeBalance()).To(Equal("948522307.6146000"))
		})

		It("failure response", func() {
			hmock.On(
				"GET",
				"https://localhost/accounts/GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H",
			).ReturnString(404, notFoundResponse)

			_, err := client.LoadAccount("GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H")
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(HavePrefix("Horizon error"))
			horizonError, ok := err.(*Error)
			Expect(ok).To(BeTrue())
			Expect(horizonError.Problem.Title).To(Equal("Resource Missing"))
		})

		It("connection error", func() {
			hmock.On(
				"GET",
				"https://localhost/accounts/GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H",
			).ReturnError("http.Client error")

			_, err := client.LoadAccount("GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H")
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("http.Client error"))
			_, ok := err.(*Error)
			Expect(ok).To(BeFalse())
		})
	})

	Describe("LoadAccountOffers", func() {
		It("success response", func() {
			hmock.On(
				"GET",
				"https://localhost/accounts/GC2BQYBXFOVPRDH35D5HT2AFVCDGXJM5YVTAF5THFSAISYOWAJQKRESK/offers?cursor=a&limit=50&order=desc",
			).ReturnString(200, accountOffersResponse)

			offers, err := client.LoadAccountOffers("GC2BQYBXFOVPRDH35D5HT2AFVCDGXJM5YVTAF5THFSAISYOWAJQKRESK", Cursor("a"), Limit(50), OrderDesc)
			Expect(err).To(BeNil())
			Expect(len(offers.Embedded.Records)).To(Equal(2))
			Expect(offers.Embedded.Records[0].ID).To(Equal(int64(161)))
			Expect(offers.Embedded.Records[0].Seller).To(Equal("GC2BQYBXFOVPRDH35D5HT2AFVCDGXJM5YVTAF5THFSAISYOWAJQKRESK"))
			Expect(offers.Embedded.Records[0].Price).To(Equal("450000.0000000"))
			Expect(offers.Embedded.Records[0].Buying.Type).To(Equal("native"))
			Expect(offers.Embedded.Records[0].Selling.Type).To(Equal("credit_alphanum4"))
			Expect(offers.Embedded.Records[0].Selling.Code).To(Equal("XBT"))
			Expect(offers.Embedded.Records[0].Selling.Issuer).To(Equal("GDI73WJ4SX7LOG3XZDJC3KCK6ED6E5NBYK2JUBQSPBCNNWEG3ZN7T75U"))
		})

		It("failure response", func() {
			hmock.On(
				"GET",
				"https://localhost/accounts/GC2BQYBXFOVPRDH35D5HT2AFVCDGXJM5YVTAF5THFSAISYOWAJQKRESK/offers",
			).ReturnString(404, notFoundResponse)

			_, err := client.LoadAccountOffers("GC2BQYBXFOVPRDH35D5HT2AFVCDGXJM5YVTAF5THFSAISYOWAJQKRESK")
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(HavePrefix("Horizon error"))
			horizonError, ok := err.(*Error)
			Expect(ok).To(BeTrue())
			Expect(horizonError.Problem.Title).To(Equal("Resource Missing"))
		})

		It("connection error", func() {
			hmock.On(
				"GET",
				"https://localhost/accounts/GC2BQYBXFOVPRDH35D5HT2AFVCDGXJM5YVTAF5THFSAISYOWAJQKRESK/offers",
			).ReturnError("http.Client error")

			_, err := client.LoadAccountOffers("GC2BQYBXFOVPRDH35D5HT2AFVCDGXJM5YVTAF5THFSAISYOWAJQKRESK")
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("http.Client error"))
			_, ok := err.(*Error)
			Expect(ok).To(BeFalse())
		})

		It("overridden location", func() {
			hmock.On(
				"GET",
				"https://localhost/beepboop",
			).ReturnString(200, accountOffersResponse)

			offers, err := client.LoadAccountOffers("GC2BQYBXFOVPRDH35D5HT2AFVCDGXJM5YVTAF5THFSAISYOWAJQKRESK", At("https://localhost/beepboop"))
			Expect(err).To(BeNil())
			Expect(len(offers.Embedded.Records)).To(Equal(2))
			Expect(offers.Embedded.Records[0].ID).To(Equal(int64(161)))
			Expect(offers.Embedded.Records[0].Seller).To(Equal("GC2BQYBXFOVPRDH35D5HT2AFVCDGXJM5YVTAF5THFSAISYOWAJQKRESK"))
			Expect(offers.Embedded.Records[0].Price).To(Equal("450000.0000000"))
			Expect(offers.Embedded.Records[0].Buying.Type).To(Equal("native"))
			Expect(offers.Embedded.Records[0].Selling.Type).To(Equal("credit_alphanum4"))
			Expect(offers.Embedded.Records[0].Selling.Code).To(Equal("XBT"))
			Expect(offers.Embedded.Records[0].Selling.Issuer).To(Equal("GDI73WJ4SX7LOG3XZDJC3KCK6ED6E5NBYK2JUBQSPBCNNWEG3ZN7T75U"))
		})
	})

	Describe("LoadOrderBook", func() {
		It("success response", func() {
			hmock.On(
				"GET",
				"https://localhost/order_book?buying_asset_code=DEMO&buying_asset_issuer=GBAMBOOZDWZPVV52RCLJQYMQNXOBLOXWNQAY2IF2FREV2WL46DBCH3BE&buying_asset_type=credit_alphanum4&selling_asset_code=&selling_asset_issuer=&selling_asset_type=native",
			).ReturnString(200, orderBookResponse)

			orderBook, err := client.LoadOrderBook(Asset{Type: "native"}, Asset{"credit_alphanum4", "DEMO", "GBAMBOOZDWZPVV52RCLJQYMQNXOBLOXWNQAY2IF2FREV2WL46DBCH3BE"})
			Expect(err).To(BeNil())
			Expect(orderBook.Selling.Type).To(Equal("native"))
			Expect(orderBook.Buying.Type).To(Equal("credit_alphanum4"))
			Expect(orderBook.Buying.Code).To(Equal("DEMO"))
			Expect(orderBook.Buying.Issuer).To(Equal("GBAMBOOZDWZPVV52RCLJQYMQNXOBLOXWNQAY2IF2FREV2WL46DBCH3BE"))

			Expect(len(orderBook.Bids)).To(Equal(20))
			Expect(orderBook.Bids[0].Price).To(Equal("0.0024937"))
			Expect(orderBook.Bids[0].Amount).To(Equal("0.4363975"))
			Expect(orderBook.Bids[0].PriceR.N).To(Equal(int32(24937)))
			Expect(orderBook.Bids[0].PriceR.D).To(Equal(int32(10000000)))

			Expect(len(orderBook.Asks)).To(Equal(20))
			Expect(orderBook.Asks[0].Price).To(Equal("0.0025093"))
			Expect(orderBook.Asks[0].Amount).To(Equal("1248.9663104"))
			Expect(orderBook.Asks[0].PriceR.N).To(Equal(int32(2017413)))
			Expect(orderBook.Asks[0].PriceR.D).To(Equal(int32(803984111)))
		})

		It("success response with limit", func() {
			hmock.On(
				"GET",
				"https://localhost/order_book?buying_asset_code=DEMO&buying_asset_issuer=GBAMBOOZDWZPVV52RCLJQYMQNXOBLOXWNQAY2IF2FREV2WL46DBCH3BE&buying_asset_type=credit_alphanum4&limit=20&selling_asset_code=&selling_asset_issuer=&selling_asset_type=native",
			).ReturnString(200, orderBookResponse)

			orderBook, err := client.LoadOrderBook(Asset{Type: "native"}, Asset{"credit_alphanum4", "DEMO", "GBAMBOOZDWZPVV52RCLJQYMQNXOBLOXWNQAY2IF2FREV2WL46DBCH3BE"}, Limit(20))
			Expect(err).To(BeNil())
			Expect(orderBook.Selling.Type).To(Equal("native"))
			Expect(orderBook.Buying.Type).To(Equal("credit_alphanum4"))
			Expect(orderBook.Buying.Code).To(Equal("DEMO"))
			Expect(orderBook.Buying.Issuer).To(Equal("GBAMBOOZDWZPVV52RCLJQYMQNXOBLOXWNQAY2IF2FREV2WL46DBCH3BE"))

			Expect(len(orderBook.Bids)).To(Equal(20))
			Expect(orderBook.Bids[0].Price).To(Equal("0.0024937"))
			Expect(orderBook.Bids[0].Amount).To(Equal("0.4363975"))
			Expect(orderBook.Bids[0].PriceR.N).To(Equal(int32(24937)))
			Expect(orderBook.Bids[0].PriceR.D).To(Equal(int32(10000000)))

			Expect(len(orderBook.Asks)).To(Equal(20))
			Expect(orderBook.Asks[0].Price).To(Equal("0.0025093"))
			Expect(orderBook.Asks[0].Amount).To(Equal("1248.9663104"))
			Expect(orderBook.Asks[0].PriceR.N).To(Equal(int32(2017413)))
			Expect(orderBook.Asks[0].PriceR.D).To(Equal(int32(803984111)))
		})

		It("failure response", func() {
			hmock.On(
				"GET",
				"https://localhost/order_book?buying_asset_code=DEMO&buying_asset_issuer=GBAMBOOZDWZPVV52RCLJQYMQNXOBLOXWNQAY2IF2FREV2WL46DBCH3BE&buying_asset_type=credit_alphanum4&limit=20&selling_asset_code=&selling_asset_issuer=&selling_asset_type=native",
			).ReturnString(404, notFoundResponse)

			_, err := client.LoadOrderBook(Asset{Type: "native"}, Asset{"credit_alphanum4", "DEMO", "GBAMBOOZDWZPVV52RCLJQYMQNXOBLOXWNQAY2IF2FREV2WL46DBCH3BE"}, Limit(20))
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(HavePrefix("Horizon error"))
			horizonError, ok := err.(*Error)
			Expect(ok).To(BeTrue())
			Expect(horizonError.Problem.Title).To(Equal("Resource Missing"))
		})

		It("connection error", func() {
			hmock.On(
				"GET",
				"https://localhost/order_book?buying_asset_code=DEMO&buying_asset_issuer=GBAMBOOZDWZPVV52RCLJQYMQNXOBLOXWNQAY2IF2FREV2WL46DBCH3BE&buying_asset_type=credit_alphanum4&limit=20&selling_asset_code=&selling_asset_issuer=&selling_asset_type=native",
			).ReturnError("http.Client error")

			_, err := client.LoadOrderBook(Asset{Type: "native"}, Asset{"credit_alphanum4", "DEMO", "GBAMBOOZDWZPVV52RCLJQYMQNXOBLOXWNQAY2IF2FREV2WL46DBCH3BE"}, Limit(20))
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("http.Client error"))
			_, ok := err.(*Error)
			Expect(ok).To(BeFalse())
		})
	})

	Describe("SubmitTransaction", func() {
		var tx = "AAAAADSMMRmQGDH6EJzkgi/7PoKhphMHyNGQgDp2tlS/dhGXAAAAZAAT3TUAAAAwAAAAAAAAAAAAAAABAAAAAAAAAAMAAAABSU5SAAAAAAA0jDEZkBgx+hCc5IIv+z6CoaYTB8jRkIA6drZUv3YRlwAAAAFVU0QAAAAAADSMMRmQGDH6EJzkgi/7PoKhphMHyNGQgDp2tlS/dhGXAAAAAAX14QAAAAAKAAAAAQAAAAAAAAAAAAAAAAAAAAG/dhGXAAAAQLuStfImg0OeeGAQmvLkJSZ1MPSkCzCYNbGqX5oYNuuOqZ5SmWhEsC7uOD9ha4V7KengiwNlc0oMNqBVo22S7gk="

		It("success response", func() {
			hmock.
				On("POST", "https://localhost/transactions").
				ReturnString(200, submitResponse)

			account, err := client.SubmitTransaction(tx)
			Expect(err).To(BeNil())
			Expect(account.Ledger).To(Equal(int32(3128812)))
		})

		It("failure response", func() {
			hmock.
				On("POST", "https://localhost/transactions").
				ReturnString(400, transactionFailure)

			_, err := client.SubmitTransaction(tx)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Horizon error"))
			horizonError, ok := errors.Cause(err).(*Error)
			Expect(ok).To(BeTrue())
			Expect(horizonError.Problem.Title).To(Equal("Transaction Failed"))
		})

		It("connection error", func() {
			hmock.
				On("POST", "https://localhost/transactions").
				ReturnError("http.Client error")

			_, err := client.SubmitTransaction(tx)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("http.Client error"))
			_, ok := err.(*Error)
			Expect(ok).To(BeFalse())
		})
	})
})

var accountResponse = `{
  "_links": {
    "self": {
      "href": "https://horizon-testnet.stellar.org/accounts/GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H"
    },
    "transactions": {
      "href": "https://horizon-testnet.stellar.org/accounts/GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H/transactions{?cursor,limit,order}",
      "templated": true
    },
    "operations": {
      "href": "https://horizon-testnet.stellar.org/accounts/GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H/operations{?cursor,limit,order}",
      "templated": true
    },
    "payments": {
      "href": "https://horizon-testnet.stellar.org/accounts/GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H/payments{?cursor,limit,order}",
      "templated": true
    },
    "effects": {
      "href": "https://horizon-testnet.stellar.org/accounts/GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H/effects{?cursor,limit,order}",
      "templated": true
    },
    "offers": {
      "href": "https://horizon-testnet.stellar.org/accounts/GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H/Offers{?cursor,limit,order}",
      "templated": true
    }
  },
  "id": "GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H",
  "paging_token": "1",
  "account_id": "GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H",
  "sequence": "7384",
  "subentry_count": 0,
  "thresholds": {
    "low_threshold": 0,
    "med_threshold": 0,
    "high_threshold": 0
  },
  "flags": {
    "auth_required": false,
    "auth_revocable": false
  },
  "balances": [
    {
      "balance": "948522307.6146000",
      "asset_type": "native"
    }
  ],
  "signers": [
    {
      "public_key": "XBT5HNPK6DAL6222MAWTLHNOZSDKPJ2AKNEQ5Q324CHHCNQFQ7EHBHZN",
      "weight": 1,
      "key": "XBT5HNPK6DAL6222MAWTLHNOZSDKPJ2AKNEQ5Q324CHHCNQFQ7EHBHZN",
      "type": "sha256_hash"
    },
    {
      "public_key": "GDQHKHMFW5ICTQYM3QWCXMSZ56BNHMQG6NH6SGV3ZNZ72KRHYV5XINCE",
      "weight": 1,
      "key": "GDQHKHMFW5ICTQYM3QWCXMSZ56BNHMQG6NH6SGV3ZNZ72KRHYV5XINCE",
      "type": "ed25519_public_key"
    }
  ],
  "data": {
    "test": "R0NCVkwzU1FGRVZLUkxQNkFKNDdVS0tXWUVCWTQ1V0hBSkhDRVpLVldNVEdNQ1Q0SDROS1FZTEg="
  }
}`

var accountOffersResponse = `{
  "_links": {
    "self": {
      "href": "https://horizon.stellar.org/accounts/GC2BQYBXFOVPRDH35D5HT2AFVCDGXJM5YVTAF5THFSAISYOWAJQKRESK/offers?order=asc\u0026limit=10\u0026cursor="
    },
    "next": {
      "href": "https://horizon.stellar.org/accounts/GC2BQYBXFOVPRDH35D5HT2AFVCDGXJM5YVTAF5THFSAISYOWAJQKRESK/offers?order=asc\u0026limit=10\u0026cursor=2539"
    },
    "prev": {
      "href": "https://horizon.stellar.org/accounts/GC2BQYBXFOVPRDH35D5HT2AFVCDGXJM5YVTAF5THFSAISYOWAJQKRESK/offers?order=desc\u0026limit=10\u0026cursor=161"
    }
  },
  "_embedded": {
    "records": [
      {
        "_links": {
          "self": {
            "href": "https://horizon.stellar.org/offers/161"
          },
          "offer_maker": {
            "href": "https://horizon.stellar.org/accounts/GC2BQYBXFOVPRDH35D5HT2AFVCDGXJM5YVTAF5THFSAISYOWAJQKRESK"
          }
        },
        "id": 161,
        "paging_token": "161",
        "seller": "GC2BQYBXFOVPRDH35D5HT2AFVCDGXJM5YVTAF5THFSAISYOWAJQKRESK",
        "selling": {
          "asset_type": "credit_alphanum4",
          "asset_code": "XBT",
          "asset_issuer": "GDI73WJ4SX7LOG3XZDJC3KCK6ED6E5NBYK2JUBQSPBCNNWEG3ZN7T75U"
        },
        "buying": {
          "asset_type": "native"
        },
        "amount": "0.0000100",
        "price_r": {
          "n": 450000,
          "d": 1
        },
        "price": "450000.0000000"
      },
      {
        "_links": {
          "self": {
            "href": "https://horizon.stellar.org/offers/2539"
          },
          "offer_maker": {
            "href": "https://horizon.stellar.org/accounts/GC2BQYBXFOVPRDH35D5HT2AFVCDGXJM5YVTAF5THFSAISYOWAJQKRESK"
          }
        },
        "id": 2539,
        "paging_token": "2539",
        "seller": "GC2BQYBXFOVPRDH35D5HT2AFVCDGXJM5YVTAF5THFSAISYOWAJQKRESK",
        "selling": {
          "asset_type": "credit_alphanum4",
          "asset_code": "EUR",
          "asset_issuer": "GDI73WJ4SX7LOG3XZDJC3KCK6ED6E5NBYK2JUBQSPBCNNWEG3ZN7T75U"
        },
        "buying": {
          "asset_type": "native"
        },
        "amount": "4.9129252",
        "price_r": {
          "n": 588,
          "d": 1
        },
        "price": "588.0000000"
      }
    ]
  }
}`

var orderBookResponse = `{
  "bids": [
    {
      "price_r": {
        "n": 24937,
        "d": 10000000
      },
      "price": "0.0024937",
      "amount": "0.4363975"
    },
    {
      "price_r": {
        "n": 24817,
        "d": 10000000
      },
      "price": "0.0024817",
      "amount": "0.6800000"
    },
    {
      "price_r": {
        "n": 1787762,
        "d": 737341097
      },
      "price": "0.0024246",
      "amount": "0.7078008"
    },
    {
      "price_r": {
        "n": 4638606,
        "d": 1914900241
      },
      "price": "0.0024224",
      "amount": "32.0067435"
    },
    {
      "price_r": {
        "n": 2048926,
        "d": 845926319
      },
      "price": "0.0024221",
      "amount": "5.9303650"
    },
    {
      "price_r": {
        "n": 4315154,
        "d": 1782416791
      },
      "price": "0.0024210",
      "amount": "2.6341583"
    },
    {
      "price_r": {
        "n": 3360181,
        "d": 1397479136
      },
      "price": "0.0024045",
      "amount": "5.9948532"
    },
    {
      "price_r": {
        "n": 2367836,
        "d": 985908229
      },
      "price": "0.0024017",
      "amount": "3.8896537"
    },
    {
      "price_r": {
        "n": 4687363,
        "d": 1952976585
      },
      "price": "0.0024001",
      "amount": "1.5747618"
    },
    {
      "price_r": {
        "n": 903753,
        "d": 380636870
      },
      "price": "0.0023743",
      "amount": "1.6182054"
    },
    {
      "price_r": {
        "n": 2562439,
        "d": 1081514977
      },
      "price": "0.0023693",
      "amount": "15.1310429"
    },
    {
      "price_r": {
        "n": 2588843,
        "d": 1129671233
      },
      "price": "0.0022917",
      "amount": "2.7172038"
    },
    {
      "price_r": {
        "n": 3249035,
        "d": 1425861493
      },
      "price": "0.0022786",
      "amount": "6.7610234"
    },
    {
      "price_r": {
        "n": 629489,
        "d": 284529942
      },
      "price": "0.0022124",
      "amount": "8.6216043"
    },
    {
      "price_r": {
        "n": 1428194,
        "d": 664535371
      },
      "price": "0.0021492",
      "amount": "11.0263350"
    },
    {
      "price_r": {
        "n": 1653667,
        "d": 771446377
      },
      "price": "0.0021436",
      "amount": "26.0527506"
    },
    {
      "price_r": {
        "n": 3613348,
        "d": 1709911165
      },
      "price": "0.0021132",
      "amount": "1.6923954"
    },
    {
      "price_r": {
        "n": 2674223,
        "d": 1280335392
      },
      "price": "0.0020887",
      "amount": "0.9882259"
    },
    {
      "price_r": {
        "n": 3594842,
        "d": 1769335169
      },
      "price": "0.0020317",
      "amount": "6.6846233"
    },
    {
      "price_r": {
        "n": 1526497,
        "d": 751849545
      },
      "price": "0.0020303",
      "amount": "3.5964310"
    }
  ],
  "asks": [
    {
      "price_r": {
        "n": 2017413,
        "d": 803984111
      },
      "price": "0.0025093",
      "amount": "1248.9663104"
    },
    {
      "price_r": {
        "n": 2687972,
        "d": 1067615183
      },
      "price": "0.0025177",
      "amount": "6286.5014925"
    },
    {
      "price_r": {
        "n": 845303,
        "d": 332925720
      },
      "price": "0.0025390",
      "amount": "1203.8364195"
    },
    {
      "price_r": {
        "n": 5147713,
        "d": 2017340695
      },
      "price": "0.0025517",
      "amount": "668.0464888"
    },
    {
      "price_r": {
        "n": 2372938,
        "d": 877879233
      },
      "price": "0.0027030",
      "amount": "4953.5042925"
    },
    {
      "price_r": {
        "n": 5177131,
        "d": 1895808254
      },
      "price": "0.0027308",
      "amount": "3691.8772552"
    },
    {
      "price_r": {
        "n": 2219932,
        "d": 812231813
      },
      "price": "0.0027331",
      "amount": "1948.1788496"
    },
    {
      "price_r": {
        "n": 4285123,
        "d": 1556796383
      },
      "price": "0.0027525",
      "amount": "5274.3733332"
    },
    {
      "price_r": {
        "n": 3945179,
        "d": 1402780548
      },
      "price": "0.0028124",
      "amount": "1361.9590574"
    },
    {
      "price_r": {
        "n": 4683683,
        "d": 1664729678
      },
      "price": "0.0028135",
      "amount": "5076.0909147"
    },
    {
      "price_r": {
        "n": 1489326,
        "d": 524639179
      },
      "price": "0.0028388",
      "amount": "2303.2370107"
    },
    {
      "price_r": {
        "n": 3365104,
        "d": 1176168157
      },
      "price": "0.0028611",
      "amount": "8080.5751770"
    },
    {
      "price_r": {
        "n": 2580607,
        "d": 899476885
      },
      "price": "0.0028690",
      "amount": "3733.5054174"
    },
    {
      "price_r": {
        "n": 5213871,
        "d": 1788590825
      },
      "price": "0.0029151",
      "amount": "485.7370041"
    },
    {
      "price_r": {
        "n": 4234565,
        "d": 1447134374
      },
      "price": "0.0029262",
      "amount": "7936.6430110"
    },
    {
      "price_r": {
        "n": 674413,
        "d": 230022877
      },
      "price": "0.0029319",
      "amount": "101.5325328"
    },
    {
      "price_r": {
        "n": 1554515,
        "d": 514487004
      },
      "price": "0.0030215",
      "amount": "5407.8562112"
    },
    {
      "price_r": {
        "n": 5638983,
        "d": 1850050675
      },
      "price": "0.0030480",
      "amount": "3024.9341116"
    },
    {
      "price_r": {
        "n": 31027,
        "d": 10000000
      },
      "price": "0.0031027",
      "amount": "18911.2836169"
    },
    {
      "price_r": {
        "n": 15899,
        "d": 5000000
      },
      "price": "0.0031798",
      "amount": "3767.4827430"
    }
  ],
  "base": {
    "asset_type": "native"
  },
  "counter": {
    "asset_type": "credit_alphanum4",
    "asset_code": "DEMO",
    "asset_issuer": "GBAMBOOZDWZPVV52RCLJQYMQNXOBLOXWNQAY2IF2FREV2WL46DBCH3BE"
  }
}`

var notFoundResponse = `{
  "type": "https://stellar.org/horizon-errors/not_found",
  "title": "Resource Missing",
  "status": 404,
  "detail": "The resource at the url requested was not found.  This is usually occurs for one of two reasons:  The url requested is not valid, or no data in our database could be found with the parameters provided.",
  "instance": "horizon-live-001/61KdRW8tKi-18408110"
}`

var submitResponse = `{
  "_links": {
    "transaction": {
      "href": "https://horizon-testnet.stellar.org/transactions/ee14b93fcd31d4cfe835b941a0a8744e23a6677097db1fafe0552d8657bed940"
    }
  },
  "hash": "ee14b93fcd31d4cfe835b941a0a8744e23a6677097db1fafe0552d8657bed940",
  "ledger": 3128812,
  "envelope_xdr": "AAAAADSMMRmQGDH6EJzkgi/7PoKhphMHyNGQgDp2tlS/dhGXAAAAZAAT3TUAAAAwAAAAAAAAAAAAAAABAAAAAAAAAAMAAAABSU5SAAAAAAA0jDEZkBgx+hCc5IIv+z6CoaYTB8jRkIA6drZUv3YRlwAAAAFVU0QAAAAAADSMMRmQGDH6EJzkgi/7PoKhphMHyNGQgDp2tlS/dhGXAAAAAAX14QAAAAAKAAAAAQAAAAAAAAAAAAAAAAAAAAG/dhGXAAAAQLuStfImg0OeeGAQmvLkJSZ1MPSkCzCYNbGqX5oYNuuOqZ5SmWhEsC7uOD9ha4V7KengiwNlc0oMNqBVo22S7gk=",
  "result_xdr": "AAAAAAAAAGQAAAAAAAAAAQAAAAAAAAADAAAAAAAAAAAAAAAAAAAAADSMMRmQGDH6EJzkgi/7PoKhphMHyNGQgDp2tlS/dhGXAAAAAAAAAPEAAAABSU5SAAAAAAA0jDEZkBgx+hCc5IIv+z6CoaYTB8jRkIA6drZUv3YRlwAAAAFVU0QAAAAAADSMMRmQGDH6EJzkgi/7PoKhphMHyNGQgDp2tlS/dhGXAAAAAAX14QAAAAAKAAAAAQAAAAAAAAAAAAAAAA==",
  "result_meta_xdr": "AAAAAAAAAAEAAAACAAAAAAAvoHwAAAACAAAAADSMMRmQGDH6EJzkgi/7PoKhphMHyNGQgDp2tlS/dhGXAAAAAAAAAPEAAAABSU5SAAAAAAA0jDEZkBgx+hCc5IIv+z6CoaYTB8jRkIA6drZUv3YRlwAAAAFVU0QAAAAAADSMMRmQGDH6EJzkgi/7PoKhphMHyNGQgDp2tlS/dhGXAAAAAAX14QAAAAAKAAAAAQAAAAAAAAAAAAAAAAAAAAEAL6B8AAAAAAAAAAA0jDEZkBgx+hCc5IIv+z6CoaYTB8jRkIA6drZUv3YRlwAAABZ9zvNAABPdNQAAADAAAAAEAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAA=="
}`

var transactionFailure = `{
  "type": "https://stellar.org/horizon-errors/transaction_failed",
  "title": "Transaction Failed",
  "status": 400,
  "detail": "The transaction failed when submitted to the stellar network. The extras.result_codes field on this response contains further details.  Descriptions of each code can be found at: https://www.stellar.org/developers/learn/concepts/list-of-operations.html",
  "instance": "horizon-testnet-001.prd.stellar001.internal.stellar-ops.com/4elYz2fHhC-528285",
  "extras": {
    "envelope_xdr": "AAAAAKpmDL6Z4hvZmkTBkYpHftan4ogzTaO4XTB7joLgQnYYAAAAZAAAAAAABeoyAAAAAAAAAAEAAAAAAAAAAQAAAAAAAAABAAAAAD3sEVVGZGi/NoC3ta/8f/YZKMzyi9ZJpOi0H47x7IqYAAAAAAAAAAAF9eEAAAAAAAAAAAA=",
    "result_codes": {
      "transaction": "tx_no_source_account"
    },
    "result_xdr": "AAAAAAAAAAD////4AAAAAA=="
  }
}`
