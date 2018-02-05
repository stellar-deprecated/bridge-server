package handlers

import (
	"strconv"
	"time"

	"github.com/stellar/gateway/compliance/config"
	"github.com/stellar/gateway/crypto"
	"github.com/stellar/gateway/db"
	"github.com/stellar/gateway/external"
	"github.com/stellar/gateway/net"
	"github.com/stellar/go/clients/federation"
)

// RequestHandler implements compliance server request handlers
type RequestHandler struct {
	Config                  *config.Config                      `inject:""`
	Client                  net.HTTPClientInterface             `inject:""`
	EntityManager           db.EntityManagerInterface           `inject:""`
	Repository              db.RepositoryInterface              `inject:""`
	SignatureSignerVerifier crypto.SignerVerifierInterface      `inject:""`
	StellarTomlResolver     external.StellarTomlClientInterface `inject:""`
	FederationResolver      federation.ClientInterface          `inject:""`
	NonceGenerator          NonceGeneratorInterface             `inject:""`
}

type NonceGeneratorInterface interface {
	Generate() string
}

type NonceGenerator struct{}

func (n *NonceGenerator) Generate() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

type TestNonceGenerator struct{}

func (n *TestNonceGenerator) Generate() string {
	return "nonce"
}
