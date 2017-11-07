package handlers

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"

	"github.com/stellar/gateway/db/entities"
	"github.com/stellar/gateway/protocols"
	"github.com/stellar/gateway/server"
	"github.com/zenazn/goji/web"
)

// HandlerAllowAccess implements /allow_access endpoint
func (rh *RequestHandler) HandlerAllowAccess(c web.C, w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	domain := r.PostFormValue("domain")
	publicKey := r.PostFormValue("public_key")
	userID := r.PostFormValue("user_id")

	// TODO check params

	var err error

	if userID != "" {
		entity := &entities.AllowedUser{
			FiName:      name,
			FiDomain:    domain,
			FiPublicKey: publicKey,
			UserID:      userID,
			AllowedAt:   time.Now(),
		}
		err = rh.EntityManager.Persist(entity)
	} else {
		entity := &entities.AllowedFi{
			Name:      name,
			Domain:    domain,
			PublicKey: publicKey,
			AllowedAt: time.Now(),
		}
		err = rh.EntityManager.Persist(entity)
	}

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Warn("Error persisting /allow entity")
		server.Write(w, protocols.InternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
