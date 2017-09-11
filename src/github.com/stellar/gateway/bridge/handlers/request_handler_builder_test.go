package handlers

import (
	// "encoding/hex"
	// "errors"
	"net/http"
	"net/http/httptest"
	// "net/url"
	"strings"
	"testing"

	"github.com/facebookgo/inject"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stellar/gateway/bridge/config"
	"github.com/stellar/gateway/horizon"
	"github.com/stellar/gateway/mocks"
	"github.com/stellar/gateway/net"
	"github.com/stellar/gateway/test"
	"github.com/stretchr/testify/assert"
)

func TestRequestHandlerBuilder(t *testing.T) {
	c := &config.Config{NetworkPassphrase: "Test SDF Network ; September 2015"}
	mockHorizon := new(mocks.MockHorizon)
	mockHTTPClient := new(mocks.MockHTTPClient)
	mockTransactionSubmitter := new(mocks.MockTransactionSubmitter)
	mockFederationResolver := new(mocks.MockFederationResolver)
	mockStellartomlResolver := new(mocks.MockStellartomlResolver)
	requestHandler := RequestHandler{}

	// Inject mocks
	var g inject.Graph

	err := g.Provide(
		&inject.Object{Value: &requestHandler},
		&inject.Object{Value: c},
		&inject.Object{Value: mockHorizon},
		&inject.Object{Value: mockHTTPClient},
		&inject.Object{Value: mockTransactionSubmitter},
		&inject.Object{Value: mockFederationResolver},
		&inject.Object{Value: mockStellartomlResolver},
	)
	if err != nil {
		panic(err)
	}

	if err := g.Populate(); err != nil {
		panic(err)
	}

	testServer := httptest.NewServer(http.HandlerFunc(requestHandler.Builder))
	defer testServer.Close()

	Convey("Builder", t, func() {

		Convey("Empty Sequence Number", func() {
			data := test.StringToJSONMap(`{
		"source": "GBWJES3WOKK7PRLJKZVGIPVFGQSSGCRMY7H3GCZ7BEG6ZTDB4FZXTPJ5",
		"sequence_number": "",
		"operations": [
		{
				"type": "create_account",
				"body": {
					"destination": "GCOEGO43PFSLE4K7WRZQNRO3PIOTRLKRASP32W7DSPBF65XFT4V6PSV3",
					"starting_balance": "50"
				}
		}
		],
		"signers": ["SABY7FRMMJWPBTKQQ2ZN43AUJQ3Z2ZAK36VYSG2SPE2ABNQXA66H5E5G"]
		}`)
			// Loading sequence number
			mockHorizon.On(
				"LoadAccount",
				"GBWJES3WOKK7PRLJKZVGIPVFGQSSGCRMY7H3GCZ7BEG6ZTDB4FZXTPJ5",
			).Return(
				horizon.AccountResponse{
					SequenceNumber: "123",
				},
				nil,
			).Once()

			Convey("it should return correct XDR", func() {
				statusCode, response := net.JSONGetResponse(testServer, data)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				expected := test.StringToJSONMap(`{
		"transaction_envelope": "AAAAAGySS3ZylffFaVZqZD6lNCUjCizHz7MLPwkN7Mxh4XN5AAAAZAAAAAAAAAB7AAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAnEM7m3lksnFftHMGxdt6HTitUQSfvVvjk8JfduWfK+cAAAAAHc1lAAAAAAAAAAABn420/AAAAECXY+neSolhAeHUXf+UrOV6PjeJnvLM/HqjOlOEWD3hmu/z9aBksDu9zqa26jS14eMpZzq8sofnnvt248FUO+cP"
		}`)
				assert.Equal(t, expected, test.StringToJSONMap(responseString))
			})
		})

		Convey("CreateAccount", func() {
			data := test.StringToJSONMap(`{
  "source": "GBWJES3WOKK7PRLJKZVGIPVFGQSSGCRMY7H3GCZ7BEG6ZTDB4FZXTPJ5",
  "sequence_number": "123",
  "operations": [
    {
        "type": "create_account",
        "body": {
        	"destination": "GCOEGO43PFSLE4K7WRZQNRO3PIOTRLKRASP32W7DSPBF65XFT4V6PSV3",
        	"starting_balance": "50"
        }
    }
  ],
  "signers": ["SABY7FRMMJWPBTKQQ2ZN43AUJQ3Z2ZAK36VYSG2SPE2ABNQXA66H5E5G"]
}`)

			Convey("it should return correct XDR", func() {
				statusCode, response := net.JSONGetResponse(testServer, data)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				expected := test.StringToJSONMap(`{
  "transaction_envelope": "AAAAAGySS3ZylffFaVZqZD6lNCUjCizHz7MLPwkN7Mxh4XN5AAAAZAAAAAAAAAB7AAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAnEM7m3lksnFftHMGxdt6HTitUQSfvVvjk8JfduWfK+cAAAAAHc1lAAAAAAAAAAABn420/AAAAECXY+neSolhAeHUXf+UrOV6PjeJnvLM/HqjOlOEWD3hmu/z9aBksDu9zqa26jS14eMpZzq8sofnnvt248FUO+cP"
}`)
				assert.Equal(t, expected, test.StringToJSONMap(responseString))
			})
		})

		Convey("Payment", func() {
			data := test.StringToJSONMap(`{
  "source": "GBWJES3WOKK7PRLJKZVGIPVFGQSSGCRMY7H3GCZ7BEG6ZTDB4FZXTPJ5",
  "sequence_number": "123",
  "operations": [
    {
        "type": "payment",
        "body": {
        	"destination": "GCOEGO43PFSLE4K7WRZQNRO3PIOTRLKRASP32W7DSPBF65XFT4V6PSV3",
        	"amount": "100",
        	"asset": {
        		"code": "USD",
        		"issuer": "GACETOPHMOLSZLG5IQ3D6KQDKCAAYUYTTQHIEY6IGZE4VOBDD2YY6YAO"
        	}
        }
    }
  ],
  "signers": ["SABY7FRMMJWPBTKQQ2ZN43AUJQ3Z2ZAK36VYSG2SPE2ABNQXA66H5E5G"]
}`)

			Convey("it should return correct XDR", func() {
				statusCode, response := net.JSONGetResponse(testServer, data)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				expected := test.StringToJSONMap(`{
  "transaction_envelope": "AAAAAGySS3ZylffFaVZqZD6lNCUjCizHz7MLPwkN7Mxh4XN5AAAAZAAAAAAAAAB7AAAAAAAAAAAAAAABAAAAAAAAAAEAAAAAnEM7m3lksnFftHMGxdt6HTitUQSfvVvjk8JfduWfK+cAAAABVVNEAAAAAAAESbnnY5csrN1ENj8qA1CADFMTnA6CY8g2Scq4Ix6xjwAAAAA7msoAAAAAAAAAAAGfjbT8AAAAQGlQbmCv74lzQpjUOn8dsQ9/BFCKHSev6DLo4lS2wcS20GpfIjGZSXIAry/3porFM+3xrvBWlIH9Tr/QFKjqRAU="
}`)
				assert.Equal(t, expected, test.StringToJSONMap(responseString))
			})
		})

		Convey("PathPayment", func() {
			data := test.StringToJSONMap(`{
  "source": "GBWJES3WOKK7PRLJKZVGIPVFGQSSGCRMY7H3GCZ7BEG6ZTDB4FZXTPJ5",
  "sequence_number": "123",
  "operations": [
    {
        "type": "path_payment",
        "body": {
        	"source": "GBYLUAJBHGZMVAYCALM4ZRTAOY74NTSSULG42VKVY2EOWS5X2HFBB2VL",
        	"destination": "GCOEGO43PFSLE4K7WRZQNRO3PIOTRLKRASP32W7DSPBF65XFT4V6PSV3",
        	"destination_amount": "500",
        	"destination_asset": {
        		"code": "EUR",
        		"issuer": "GDOJMKTDLGGLROSSM5BV5MXIAQ3JZHASQFUV55WBJ45AFOUXSVVFGPTJ"
        	},
        	"send_max": "100",
        	"send_asset": {
        		"code": "USD",
        		"issuer": "GACETOPHMOLSZLG5IQ3D6KQDKCAAYUYTTQHIEY6IGZE4VOBDD2YY6YAO"
        	},
        	"path": [
        		{
	        		"code": "ABCDEFG",
	        		"issuer": "GD4RIHH2HWB4MPJN72G2VGLRPUXDODFNQG6DVU47HMSSSF3RIQ4UXALD"
	        	},
	        	{}
        	]
        }
    }
  ],
  "signers": ["SABY7FRMMJWPBTKQQ2ZN43AUJQ3Z2ZAK36VYSG2SPE2ABNQXA66H5E5G"]
}`)

			Convey("it should return correct XDR", func() {
				statusCode, response := net.JSONGetResponse(testServer, data)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				expected := test.StringToJSONMap(`{
  "transaction_envelope": "AAAAAGySS3ZylffFaVZqZD6lNCUjCizHz7MLPwkN7Mxh4XN5AAAAZAAAAAAAAAB7AAAAAAAAAAAAAAABAAAAAQAAAABwugEhObLKgwIC2czGYHY/xs5Sos3NVVXGiOtLt9HKEAAAAAIAAAABVVNEAAAAAAAESbnnY5csrN1ENj8qA1CADFMTnA6CY8g2Scq4Ix6xjwAAAAA7msoAAAAAAJxDO5t5ZLJxX7RzBsXbeh04rVEEn71b45PCX3blnyvnAAAAAUVVUgAAAAAA3JYqY1mMuLpSZ0NesugENpycEoFpXvbBTzoCupeValMAAAABKgXyAAAAAAIAAAACQUJDREVGRwAAAAAAAAAAAPkUHPo9g8Y9Lf6NqplxfS43DK2BvDrTnzslKRdxRDlLAAAAAAAAAAAAAAABn420/AAAAEA9DEvKZhLwLcStP8/ZsqaEAdlNc91Eyz5mLUiN19etsIYaTPNugsVEWYJOiulXXSIwwitoyxQ1t2jr6VS0mXcB"
}`)
				assert.Equal(t, expected, test.StringToJSONMap(responseString))
			})
		})

		Convey("ManageOffer", func() {
			data := test.StringToJSONMap(`{
  "source": "GBWJES3WOKK7PRLJKZVGIPVFGQSSGCRMY7H3GCZ7BEG6ZTDB4FZXTPJ5",
  "sequence_number": "123",
  "operations": [
    {
        "type": "manage_offer",
        "body": {
        	"selling": {
        		"code": "EUR",
        		"issuer": "GDOJMKTDLGGLROSSM5BV5MXIAQ3JZHASQFUV55WBJ45AFOUXSVVFGPTJ"
        	},
        	"buying": {
        		"code": "USD",
        		"issuer": "GACETOPHMOLSZLG5IQ3D6KQDKCAAYUYTTQHIEY6IGZE4VOBDD2YY6YAO"
        	},
        	"amount": "123456",
        	"price": "2.93850088",
        	"offer_id": "100"
        }
    }
  ],
  "signers": ["SABY7FRMMJWPBTKQQ2ZN43AUJQ3Z2ZAK36VYSG2SPE2ABNQXA66H5E5G"]
}`)

			Convey("it should return correct XDR", func() {
				statusCode, response := net.JSONGetResponse(testServer, data)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				expected := test.StringToJSONMap(`{
  "transaction_envelope": "AAAAAGySS3ZylffFaVZqZD6lNCUjCizHz7MLPwkN7Mxh4XN5AAAAZAAAAAAAAAB7AAAAAAAAAAAAAAABAAAAAAAAAAMAAAABRVVSAAAAAADclipjWYy4ulJnQ16y6AQ2nJwSgWle9sFPOgK6l5VqUwAAAAFVU0QAAAAAAARJuedjlyys3UQ2PyoDUIAMUxOcDoJjyDZJyrgjHrGPAAABH3GCoAACMHl9AL68IAAAAAAAAABkAAAAAAAAAAGfjbT8AAAAQEpMML2mghfM2Dzkpw6eT1N00rrIC7v3xe8zy7yc8rcGzFxIw/4/E69uq+rst+xDoeMTn0b3iBtjr2DEV52o/wE="
}`)
				assert.Equal(t, expected, test.StringToJSONMap(responseString))
			})
		})

		Convey("CreatePassiveOffer", func() {
			data := test.StringToJSONMap(`{
  "source": "GBWJES3WOKK7PRLJKZVGIPVFGQSSGCRMY7H3GCZ7BEG6ZTDB4FZXTPJ5",
  "sequence_number": "123",
  "operations": [
    {
        "type": "create_passive_offer",
        "body": {
        	"selling": {
        		"code": "EUR",
        		"issuer": "GDOJMKTDLGGLROSSM5BV5MXIAQ3JZHASQFUV55WBJ45AFOUXSVVFGPTJ"
        	},
        	"buying": {
        		"code": "USD",
        		"issuer": "GACETOPHMOLSZLG5IQ3D6KQDKCAAYUYTTQHIEY6IGZE4VOBDD2YY6YAO"
        	},
        	"amount": "123456",
        	"price": "2.93850088"
        }
    }
  ],
  "signers": ["SABY7FRMMJWPBTKQQ2ZN43AUJQ3Z2ZAK36VYSG2SPE2ABNQXA66H5E5G"]
}`)

			Convey("it should return correct XDR", func() {
				statusCode, response := net.JSONGetResponse(testServer, data)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				expected := test.StringToJSONMap(`{
  "transaction_envelope": "AAAAAGySS3ZylffFaVZqZD6lNCUjCizHz7MLPwkN7Mxh4XN5AAAAZAAAAAAAAAB7AAAAAAAAAAAAAAABAAAAAAAAAAQAAAABRVVSAAAAAADclipjWYy4ulJnQ16y6AQ2nJwSgWle9sFPOgK6l5VqUwAAAAFVU0QAAAAAAARJuedjlyys3UQ2PyoDUIAMUxOcDoJjyDZJyrgjHrGPAAABH3GCoAACMHl9AL68IAAAAAAAAAABn420/AAAAEAtK8juIThYp4LXtgpN8gVNRR42iiR6tz8euSKqqqzKGELCHcPrmFUuYqtecrJi8CyPCYTp0nqGY9mtJCHFYpsC"
}`)
				assert.Equal(t, expected, test.StringToJSONMap(responseString))
			})
		})

		Convey("SetOptions", func() {
			data := test.StringToJSONMap(`{
  "source": "GBWJES3WOKK7PRLJKZVGIPVFGQSSGCRMY7H3GCZ7BEG6ZTDB4FZXTPJ5",
  "sequence_number": "123",
  "operations": [
    {
        "type": "set_options",
        "body": {
        	"inflation_dest": "GBMPZVOMJ67WQBTBCVURDKTGL4557272EGQMAJCXPSMLOE63XPLL6SVA",
        	"set_flags": [1, 2],
        	"clear_flags": [4],
        	"master_weight": 100,
        	"low_threshold": 1,
        	"medium_threshold": 2,
        	"high_threshold": 3,
        	"home_domain": "stellar.org",
        	"signer": {
        		"public_key": "GA6VMJJQM2QBPPIXK2UVTAOS4XSSSAKSCOGFQE55IMRBQR65GIVDTTQV",
        		"weight": 5
        	}
        }
    }
  ],
  "signers": ["SABY7FRMMJWPBTKQQ2ZN43AUJQ3Z2ZAK36VYSG2SPE2ABNQXA66H5E5G"]
}`)

			Convey("it should return correct XDR", func() {
				statusCode, response := net.JSONGetResponse(testServer, data)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				expected := test.StringToJSONMap(`{
  "transaction_envelope": "AAAAAGySS3ZylffFaVZqZD6lNCUjCizHz7MLPwkN7Mxh4XN5AAAAZAAAAAAAAAB7AAAAAAAAAAAAAAABAAAAAAAAAAUAAAABAAAAAFj81cxPv2gGYRVpEapmXzvf6/ohoMAkV3yYtxPbu9a/AAAAAQAAAAQAAAABAAAAAwAAAAEAAABkAAAAAQAAAAEAAAABAAAAAgAAAAEAAAADAAAAAQAAAAtzdGVsbGFyLm9yZwAAAAABAAAAAD1WJTBmoBe9F1apWYHS5eUpAVITjFgTvUMiGEfdMio5AAAABQAAAAAAAAABn420/AAAAEAtQAlVOLBR6sb/YHRg7XcSEPSJ07irs6cCSDpK95rYE7Ga5ghiLXHqRJQ2B9cMmf8FYqzeaHdYPiESZqowhb0F"
}`)
				assert.Equal(t, expected, test.StringToJSONMap(responseString))
			})
		})

		Convey("ChangeTrust", func() {
			data := test.StringToJSONMap(`{
  "source": "GBWJES3WOKK7PRLJKZVGIPVFGQSSGCRMY7H3GCZ7BEG6ZTDB4FZXTPJ5",
  "sequence_number": "123",
  "operations": [
    {
        "type": "change_trust",
        "body": {
        	"asset": {
        		"code": "USD",
        		"issuer": "GCHGRVNTXAV3OXNMCSA63BUCD6AZZX6PN2542QB6GIVTXGHQ65XS35DS"
        	}
        }
    }
  ],
  "signers": ["SABY7FRMMJWPBTKQQ2ZN43AUJQ3Z2ZAK36VYSG2SPE2ABNQXA66H5E5G"]
}`)

			Convey("it should return correct XDR", func() {
				statusCode, response := net.JSONGetResponse(testServer, data)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				expected := test.StringToJSONMap(`{
  "transaction_envelope": "AAAAAGySS3ZylffFaVZqZD6lNCUjCizHz7MLPwkN7Mxh4XN5AAAAZAAAAAAAAAB7AAAAAAAAAAAAAAABAAAAAAAAAAYAAAABVVNEAAAAAACOaNWzuCu3XawUge2Ggh+BnN/PbrvNQD4yKzuY8PdvLX//////////AAAAAAAAAAGfjbT8AAAAQFftcSiqTvZOQwDJnoJ7buLgYXyjRacggCZ7yEhnPN4eXxlpQycvLLFa3U8xv0Mcnx5frSNKxu0sDIOm88Iicw8="
}`)
				assert.Equal(t, expected, test.StringToJSONMap(responseString))
			})
		})

		Convey("AllowTrust", func() {
			data := test.StringToJSONMap(`{
  "source": "GBWJES3WOKK7PRLJKZVGIPVFGQSSGCRMY7H3GCZ7BEG6ZTDB4FZXTPJ5",
  "sequence_number": "123",
  "operations": [
    {
        "type": "allow_trust",
        "body": {
        	"asset_code": "USDUSD",
        	"trustor": "GBLH67TQHRNRLERQEIQJDNBV2DSWPHAPP43MBIF7DVKA7X55APUNS4LL",
        	"authorize": true
        }
    }
  ],
  "signers": ["SABY7FRMMJWPBTKQQ2ZN43AUJQ3Z2ZAK36VYSG2SPE2ABNQXA66H5E5G"]
}`)

			Convey("it should return correct XDR", func() {
				statusCode, response := net.JSONGetResponse(testServer, data)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				expected := test.StringToJSONMap(`{
  "transaction_envelope": "AAAAAGySS3ZylffFaVZqZD6lNCUjCizHz7MLPwkN7Mxh4XN5AAAAZAAAAAAAAAB7AAAAAAAAAAAAAAABAAAAAAAAAAcAAAAAVn9+cDxbFZIwIiCRtDXQ5WecD382wKC/HVQP370D6NkAAAACVVNEVVNEAAAAAAAAAAAAAQAAAAAAAAABn420/AAAAEA9Ht9mJaKdYoRg/rAX/cl/Q89Juhmi8f7iGBdCrSVAs+VN7NVJXR+0aZpoZIjcJD/QBPiuzZIK1ea2fN7I0I8J"
}`)
				assert.Equal(t, expected, test.StringToJSONMap(responseString))
			})
		})

		Convey("AccountMerge", func() {
			data := test.StringToJSONMap(`{
  "source": "GBWJES3WOKK7PRLJKZVGIPVFGQSSGCRMY7H3GCZ7BEG6ZTDB4FZXTPJ5",
  "sequence_number": "123",
  "operations": [
    {
        "type": "account_merge",
        "body": {
        	"destination": "GBLH67TQHRNRLERQEIQJDNBV2DSWPHAPP43MBIF7DVKA7X55APUNS4LL"
        }
    }
  ],
  "signers": ["SABY7FRMMJWPBTKQQ2ZN43AUJQ3Z2ZAK36VYSG2SPE2ABNQXA66H5E5G"]
}`)

			Convey("it should return correct XDR", func() {
				statusCode, response := net.JSONGetResponse(testServer, data)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				expected := test.StringToJSONMap(`{
  "transaction_envelope": "AAAAAGySS3ZylffFaVZqZD6lNCUjCizHz7MLPwkN7Mxh4XN5AAAAZAAAAAAAAAB7AAAAAAAAAAAAAAABAAAAAAAAAAgAAAAAVn9+cDxbFZIwIiCRtDXQ5WecD382wKC/HVQP370D6NkAAAAAAAAAAZ+NtPwAAABALCyRn/E/CgLdPWGgP+1pd2Lkf3jWgNANKQ4QeGgUxgROhqkTUXaPA6XzOWS8yUpzZMufl6nkh8UFqa6Hc1emCA=="
}`)
				assert.Equal(t, expected, test.StringToJSONMap(responseString))
			})
		})

		Convey("Inflation", func() {
			data := test.StringToJSONMap(`{
  "source": "GBWJES3WOKK7PRLJKZVGIPVFGQSSGCRMY7H3GCZ7BEG6ZTDB4FZXTPJ5",
  "sequence_number": "123",
  "operations": [
    {
        "type": "inflation",
        "body": {}
    }
  ],
  "signers": ["SABY7FRMMJWPBTKQQ2ZN43AUJQ3Z2ZAK36VYSG2SPE2ABNQXA66H5E5G"]
}`)

			Convey("it should return correct XDR", func() {
				statusCode, response := net.JSONGetResponse(testServer, data)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				expected := test.StringToJSONMap(`{
  "transaction_envelope": "AAAAAGySS3ZylffFaVZqZD6lNCUjCizHz7MLPwkN7Mxh4XN5AAAAZAAAAAAAAAB7AAAAAAAAAAAAAAABAAAAAAAAAAkAAAAAAAAAAZ+NtPwAAABAlBFCwJ3VzBd+CE+n3mA4t71SVrDIjSgRyBnz9zYLN7qkqu8AD6cyvMRj8/alSozSPAZcSe+qBEO7E5biR+YrAA=="
}`)
				assert.Equal(t, expected, test.StringToJSONMap(responseString))
			})
		})

		Convey("ManageData", func() {
			data := test.StringToJSONMap(`{
  "source": "GBWJES3WOKK7PRLJKZVGIPVFGQSSGCRMY7H3GCZ7BEG6ZTDB4FZXTPJ5",
  "sequence_number": "123",
  "operations": [
    {
        "type": "manage_data",
        "body": {
        	"name": "test_data",
        	"data": "AQIDBAUG"
        }
    }
  ],
  "signers": ["SABY7FRMMJWPBTKQQ2ZN43AUJQ3Z2ZAK36VYSG2SPE2ABNQXA66H5E5G"]
}`)

			Convey("it should return correct XDR", func() {
				statusCode, response := net.JSONGetResponse(testServer, data)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				expected := test.StringToJSONMap(`{
  "transaction_envelope": "AAAAAGySS3ZylffFaVZqZD6lNCUjCizHz7MLPwkN7Mxh4XN5AAAAZAAAAAAAAAB7AAAAAAAAAAAAAAABAAAAAAAAAAoAAAAJdGVzdF9kYXRhAAAAAAAAAQAAAAYBAgMEBQYAAAAAAAAAAAABn420/AAAAEBkO27ebDbsn1WzzLH5lUfJH3Y0Pgd1dlRx3Ip1dEZkvRPFFDLZuXi5DlW9uxNgeqThNsqnK7PPHfhyuWBVQpgN"
}`)
				assert.Equal(t, expected, test.StringToJSONMap(responseString))
			})
		})
	})
}
