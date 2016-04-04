package handlers

import (
	"github.com/stellar/gateway/compliance/config"
	"github.com/stellar/gateway/db"
	"github.com/stellar/gateway/protocols/federation"
	"github.com/stellar/gateway/protocols/stellartoml"
)

type RequestHandler struct {
	Config              *config.Config
	EntityManager       db.EntityManagerInterface
	Repository          db.RepositoryInterface
	StellarTomlResolver stellartoml.ResolverInterface `inject:""`
	FederationResolver  federation.ResolverInterface  `inject:""`
}
