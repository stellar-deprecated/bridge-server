package build

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBuild(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Package: github.com/stellar/go-stellar-base/build")
}

// ExampleTransactionBuilder creates and signs a simple transaction, and then
// encodes it into a base64 string capable of being submitted to stellar-core.
//
// It uses the transaction builder system
func ExampleTransactionBuilder() {
	seed := "SDOTALIMPAM2IV65IOZA7KZL7XWZI5BODFXTRVLIHLQZQCKK57PH5F3H"
	tx := Transaction(
		SourceAccount{seed},
		Sequence{1},
		Payment(
			Destination{"GAWSI2JO2CF36Z43UGMUJCDQ2IMR5B3P5TMS7XM7NUTU3JHG3YJUDQXA"},
			NativeAmount{"50"},
		),
	)

	txe := tx.Sign(seed)
	txeB64, _ := txe.Base64()

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAADZY/nWY0gx6beMpf4S8Ur0qHsjA8fbFtBzBx1cbQzHwAAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAEAAAAALSRpLtCLv2eboZlEiHDSGR6Hb+zZL92fbSdNpObeE0EAAAAAAAAAAB3NZQAAAAAAAAAAARtDMfAAAABA2oIeQxoJl53RMRWFeLB865zcky39f2gf2PmUubCuJYccEePRSrTC8QQrMOgGwD8a6oe8dgltvezdDsmmXBPyBw==
}

// ExamplePathPayment creates and signs a simple transaction with PathPayment operation, and then
// encodes it into a base64 string capable of being submitted to stellar-core.
func ExamplePathPayment() {
	seed := "SDOTALIMPAM2IV65IOZA7KZL7XWZI5BODFXTRVLIHLQZQCKK57PH5F3H"
	tx := Transaction(
		SourceAccount{seed},
		Sequence{1},
		Payment(
			Destination{"GBDT3K42LOPSHNAEHEJ6AVPADIJ4MAR64QEKKW2LQPBSKLYD22KUEH4P"},
			CreditAmount{"USD", "GAWSI2JO2CF36Z43UGMUJCDQ2IMR5B3P5TMS7XM7NUTU3JHG3YJUDQXA", "50"},
			PayWith(CreditAsset("EUR", "GCPZJ3MJQ3GUGJSBL6R3MLYZS6FKVHG67BPAINMXL3NWNXR5S6XG657P"), "100").
				Through(Asset{Native: true}).
				Through(CreditAsset("BTC", "GAHJZHVKFLATAATJH46C7OK2ZOVRD47GZBGQ7P6OCVF6RJDCEG5JMQBQ")),
		),
	)

	txe := tx.Sign(seed)
	txeB64, _ := txe.Base64()

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAADZY/nWY0gx6beMpf4S8Ur0qHsjA8fbFtBzBx1cbQzHwAAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAIAAAABRVVSAAAAAACflO2Jhs1DJkFfo7YvGZeKqpze+F4ENZde22bePZeubwAAAAA7msoAAAAAAEc9q5pbnyO0BDkT4FXgGhPGAj7kCKVbS4PDJS8D1pVCAAAAAVVTRAAAAAAALSRpLtCLv2eboZlEiHDSGR6Hb+zZL92fbSdNpObeE0EAAAAAHc1lAAAAAAIAAAAAAAAAAUJUQwAAAAAADpyeqirBMAJpPzwvuVrLqxHz5shND7/OFUvopGIhupYAAAAAAAAAARtDMfAAAABA5xuIJu/KGKQRuDrdkzNsR4HjT6wX464SHZ/yvYwVb/AkAyyfeMLDNhgKbBxQMWc3Uo5fTst1UHldC+jYNeAhCQ==
}

// ExampleChangeTrust creates and signs a simple transaction with ChangeTrust operation, and then
// encodes it into a base64 string capable of being submitted to stellar-core.
func ExampleChangeTrust() {
	seed := "SDOTALIMPAM2IV65IOZA7KZL7XWZI5BODFXTRVLIHLQZQCKK57PH5F3H"
	tx := Transaction(
		SourceAccount{seed},
		Sequence{1},
		Trust("USD", "GAWSI2JO2CF36Z43UGMUJCDQ2IMR5B3P5TMS7XM7NUTU3JHG3YJUDQXA", Limit("100.25")),
	)

	txe := tx.Sign(seed)
	txeB64, _ := txe.Base64()

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAADZY/nWY0gx6beMpf4S8Ur0qHsjA8fbFtBzBx1cbQzHwAAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAYAAAABVVNEAAAAAAAtJGku0Iu/Z5uhmUSIcNIZHodv7Nkv3Z9tJ02k5t4TQQAAAAA7wO+gAAAAAAAAAAEbQzHwAAAAQOIy19X38Y3jcFzvhDsmXu6iDzrzb4iwfS2NAq9GGAFiRJUGoFX85vKtlNcXzQppF4X8oIMNPEb74fuZE/N+GAE=
}

// ExampleChangeTrustMaxLimit creates and signs a simple transaction with ChangeTrust operation (maximum limit), and then
// encodes it into a base64 string capable of being submitted to stellar-core.
func ExampleChangeTrustMaxLimit() {
	seed := "SDOTALIMPAM2IV65IOZA7KZL7XWZI5BODFXTRVLIHLQZQCKK57PH5F3H"
	tx := Transaction(
		SourceAccount{seed},
		Sequence{1},
		Trust("USD", "GAWSI2JO2CF36Z43UGMUJCDQ2IMR5B3P5TMS7XM7NUTU3JHG3YJUDQXA"),
	)

	txe := tx.Sign(seed)
	txeB64, _ := txe.Base64()

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAADZY/nWY0gx6beMpf4S8Ur0qHsjA8fbFtBzBx1cbQzHwAAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAYAAAABVVNEAAAAAAAtJGku0Iu/Z5uhmUSIcNIZHodv7Nkv3Z9tJ02k5t4TQX//////////AAAAAAAAAAEbQzHwAAAAQJQC6R3RqNaw5rOmaxqpAE0lD5onM/njn9I2RVlhtS2SGi2Z7xm65USYVWXTJFVqTCfTwwu+QXFcOuqgJjVtHAk=
}

// ExampleRemoveTrust creates and signs a simple transaction with ChangeTrust operation (remove trust), and then
// encodes it into a base64 string capable of being submitted to stellar-core.
func ExampleRemoveTrust() {
	seed := "SDOTALIMPAM2IV65IOZA7KZL7XWZI5BODFXTRVLIHLQZQCKK57PH5F3H"
	operationSource := "GCVJCNUHSGKOTBBSXZJ7JJZNOSE2YDNGRLIDPMQDUEQWJQSE6QZSDPNU"
	tx := Transaction(
		SourceAccount{seed},
		Sequence{1},
		RemoveTrust(
			"USD",
			"GAWSI2JO2CF36Z43UGMUJCDQ2IMR5B3P5TMS7XM7NUTU3JHG3YJUDQXA",
			SourceAccount{operationSource},
		),
	)

	txe := tx.Sign(seed)
	txeB64, _ := txe.Base64()

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAADZY/nWY0gx6beMpf4S8Ur0qHsjA8fbFtBzBx1cbQzHwAAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAQAAAACqkTaHkZTphDK+U/SnLXSJrA2mitA3sgOhIWTCRPQzIQAAAAYAAAABVVNEAAAAAAAtJGku0Iu/Z5uhmUSIcNIZHodv7Nkv3Z9tJ02k5t4TQQAAAAAAAAAAAAAAAAAAAAEbQzHwAAAAQD5FeGBEwJyeauK+WKfcxYBeKw62EtCqvC0p9Z+1cY32fKQ+5Jz9uE1LaDsHW5NurtStKcUTiG5j2qNDf1QpYgw=
}
