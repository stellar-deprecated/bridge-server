package handlers

import (
	log "github.com/Sirupsen/logrus"
	"net/http"

	"github.com/stellar/gateway/protocols/compliance"
	"github.com/stellar/gateway/server"
	"github.com/zenazn/goji/web"
)

func (rh *RequestHandler) HandlerRemoveAccess(c web.C, w http.ResponseWriter, r *http.Request) {
	domain := r.PostFormValue("domain")
	userId := r.PostFormValue("user_id")

	// TODO check params

	var entityManagerErr error

	if userId != "" {
		allowedUser, err := rh.Repository.GetAllowedUserByDomainAndUserId(domain, userId)
		if err != nil {
			log.WithFields(log.Fields{"err": err}).Warn("Error getting allowed user")
			server.Write(w, compliance.InternalServerError)
			return
		}

		if allowedUser == nil {
			log.WithFields(log.Fields{"err": err}).Warn("User does not exist")
			server.Write(w, compliance.InternalServerError)
			return
		}

		entityManagerErr = rh.EntityManager.Delete(allowedUser)
	} else {
		allowedFi, err := rh.Repository.GetAllowedFiByDomain(domain)
		if err != nil {
			log.WithFields(log.Fields{"err": err}).Warn("Error getting allowed FI")
			server.Write(w, compliance.InternalServerError)
			return
		}

		if allowedFi == nil {
			log.WithFields(log.Fields{"err": err}).Warn("FI does not exist")
			server.Write(w, compliance.InternalServerError)
			return
		}

		entityManagerErr = rh.EntityManager.Delete(allowedFi)
	}

	if entityManagerErr != nil {
		log.WithFields(log.Fields{"err": entityManagerErr}).Warn("Error deleting /allow entity")
		server.Write(w, compliance.InternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
