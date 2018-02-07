package stellar

import (
	"sync"

	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/support/log"
)

// AccountConfigurator is responsible for configuring new Stellar accounts that
// participate in ICO.
type AccountConfigurator struct {
	Horizon           horizon.ClientInterface `inject:""`
	NetworkPassphrase string
	IssuerPublicKey   string
	SignerSecretKey   string
	NeedsAuthorize    bool
	TokenAssetCode    string
	OnAccountCreated  func(destination string)
	OnAccountCredited func(destination string, assetCode string, amount string)

	signerPublicKey      string
	sequence             uint64
	sequenceMutex        sync.Mutex
	processingCount      int
	processingCountMutex sync.Mutex
	log                  *log.Entry
}
