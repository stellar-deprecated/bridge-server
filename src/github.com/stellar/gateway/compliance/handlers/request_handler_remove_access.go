package handlers

import (
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/stellar/gateway/protocols"
	"github.com/stellar/gateway/server"
	"github.com/zenazn/goji/web"
)

// HandlerRemoveAccess implements /remove_access endpoint
func (rh *RequestHandler) HandlerRemoveAccess(c web.C, w http.ResponseWriter, r *http.Request) {
	domain := r.PostFormValue("domain")
	userID := r.PostFormValue("user_id")

	// TODO check params

	var entityManagerErr error

	if userID != "" {
		allowedUser, err := rh.Repository.GetAllowedUserByDomainAndUserID(domain, userID)
		if err != nil {
			log.WithFields(log.Fields{"err": err}).Warn("Error getting allowed user")
			server.Write(w, protocols.InternalServerError)
			return
		}

		if allowedUser == nil {
			log.WithFields(log.Fields{"err": err}).Warn("User does not exist")
			server.Write(w, protocols.InternalServerError)
			return
		}

		entityManagerErr = rh.EntityManager.Delete(allowedUser)
	} else {
		allowedFi, err := rh.Repository.GetAllowedFiByDomain(domain)
		if err != nil {
			log.WithFields(log.Fields{"err": err}).Warn("Error getting allowed FI")
			server.Write(w, protocols.InternalServerError)
			return
		}

		if allowedFi == nil {
			log.WithFields(log.Fields{"err": err}).Warn("FI does not exist")
			server.Write(w, protocols.InternalServerError)
			return
		}

		entityManagerErr = rh.EntityManager.Delete(allowedFi)
	}

	if entityManagerErr != nil {
		log.WithFields(log.Fields{"err": entityManagerErr}).Warn("Error deleting /allow entity")
		server.Write(w, protocols.InternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
