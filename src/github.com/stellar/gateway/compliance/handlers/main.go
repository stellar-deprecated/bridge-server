package handlers

import (
	"github.com/stellar/gateway/compliance/config"
	"github.com/stellar/gateway/crypto"
	"github.com/stellar/gateway/db"
	"github.com/stellar/gateway/net"
	"github.com/stellar/gateway/protocols/federation"
	"github.com/stellar/gateway/protocols/stellartoml"
)

// RequestHandler implements compliance server request handlers
type RequestHandler struct {
	Config                  *config.Config                 `inject:""`
	Client                  net.HTTPClientInterface        `inject:""`
	EntityManager           db.EntityManagerInterface      `inject:""`
	Repository              db.RepositoryInterface         `inject:""`
	SignatureSignerVerifier crypto.SignerVerifierInterface `inject:""`
	StellarTomlResolver     stellartoml.ResolverInterface  `inject:""`
	FederationResolver      federation.ResolverInterface   `inject:""`
}
