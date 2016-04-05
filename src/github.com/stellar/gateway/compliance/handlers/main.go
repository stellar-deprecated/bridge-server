package handlers

import (
	"github.com/stellar/gateway/compliance/config"
	"github.com/stellar/gateway/db"
	"github.com/stellar/gateway/net"
	"github.com/stellar/gateway/protocols/federation"
	"github.com/stellar/gateway/protocols/stellartoml"
)

type RequestHandler struct {
	Config              *config.Config                `inject:""`
	Client              net.HttpClientInterface       `inject:""`
	EntityManager       db.EntityManagerInterface     `inject:""`
	Repository          db.RepositoryInterface        `inject:""`
	StellarTomlResolver stellartoml.ResolverInterface `inject:""`
	FederationResolver  federation.ResolverInterface  `inject:""`
}
