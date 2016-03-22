package handlers

import (
	"github.com/stellar/gateway/compliance/config"
	"github.com/stellar/gateway/db"
)

type RequestHandler struct {
	Config        *config.Config
	EntityManager db.EntityManagerInterface
	Repository    db.RepositoryInterface
}
