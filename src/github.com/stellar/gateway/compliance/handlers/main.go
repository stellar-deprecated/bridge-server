package handlers

import (
	"github.com/stellar/gateway/compliance/config"
	"github.com/stellar/gateway/crypto"
	"github.com/stellar/gateway/db"
	"github.com/stellar/gateway/net"
	"github.com/stellar/go/clients/federation"
	"github.com/stellar/go/clients/stellartoml"
)

// RequestHandler implements compliance server request handlers
type RequestHandler struct {
	Config                  *config.Config                 `inject:""`
	Client                  net.HTTPClientInterface        `inject:""`
	EntityManager           db.EntityManagerInterface      `inject:""`
	Repository              db.RepositoryInterface         `inject:""`
	SignatureSignerVerifier crypto.SignerVerifierInterface `inject:""`
	StellarTomlClient       stellartoml.Client             `inject:""`
	FederationClient        federation.Client              `inject:""`
}
