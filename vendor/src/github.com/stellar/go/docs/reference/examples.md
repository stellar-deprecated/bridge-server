---
title: Basic Examples
---

- [Creating a payment Transaction](#creating-and-submitting-a-payment-transaction)


## Creating and submitting a payment transaction

Crafting transactions and getting the base64-encoded transaction envelope (often referred to as the "transaction blob") is a central aspect of interacting with the stellar network.  The Go SDK uses the `build` package to craft transactions.  The example below builds a payment for testnet and outputs the encoded blob to standard out.  For this example, we have two previously created accounts on the test network.

```go
package main

import (
	"fmt"

	b "github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
)

func main() {
	// address: GB6S3XHQVL6ZBAF6FIK62OCK3XTUI4L5Z5YUVYNBZUXZ4AZMVBQZNSAU
	from := "SCRUYGFG76UPX3EIUWGPIQPQDPD24XPR3RII5BD53DYPKZJGG43FL5HI"

	// seed: SDLJZXOSOMKPWAK4OCWNNVOYUEYEESPGCWK53PT7QMG4J4KGDAUIL5LG
	to := "GA3A7AD7ZR4PIYW6A52SP6IK7UISESICPMMZVJGNUTVIZ5OUYOPBTK6X"

	tx := b.Transaction(
 		b.SourceAccount{from},
 		b.TestNetwork,
 		b.AutoSequence{horizon.DefaultTestNetClient},
 		b.Payment(
 			b.Destination{to},
  			b.NativeAmount{"0.1"},
  		),
  	)
	
	txe := tx.Sign(from)
	txeB64, err := txe.Base64()

	if err != nil {
		panic(err)
	}

	fmt.Printf("tx base64: %s", txeB64)
}
```

The above program will output something similar to `tx base64: AAAAAH0t3PCq/ZCAvioV7ThK3edEcX3PcUrhoc0vngMsqGGWAAAAZAA1fpcAAAABAAAAAAAAAAAAAAABAAAAAAAAAAEAAAAANg+Af8x49GLeB3Un+Qr9ESJJAnsZmqTNpOqM9dTDnhkAAAAAAAAAAAAPQkAAAAAAAAAAASyoYZYAAABA/L7Du9hlYzCtR9F89Mp9/alkCXsq9CWuJ1Mpql+Q16fHE4P2+H62p4cx+b2YUp/fUX73ucW+RPxOgSXmeV6uBQ==`, the transaction blob.  Now we need to submit the transaction to the testnet using a horizon client:


```go

package main

import (
	"fmt"

	"github.com/stellar/go/clients/horizon"
)

func main() {
	blob := "AAAAAH0t3PCq/ZCAvioV7ThK3edEcX3PcUrhoc0vngMsqGGWAAAAZAA1fpcAAAABAAAAAAAAAAAAAAABAAAAAAAAAAEAAAAANg+Af8x49GLeB3Un+Qr9ESJJAnsZmqTNpOqM9dTDnhkAAAAAAAAAAAAPQkAAAAAAAAAAASyoYZYAAABA/L7Du9hlYzCtR9F89Mp9/alkCXsq9CWuJ1Mpql+Q16fHE4P2+H62p4cx+b2YUp/fUX73ucW+RPxOgSXmeV6uBQ=="

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(blob)
	if err != nil {
		panic(err)
	}

	fmt.Println("transaction posted in ledger:", resp.Ledger)
}

```

The script above will return an error if the transaction was not successfully accepted by the testnet.
