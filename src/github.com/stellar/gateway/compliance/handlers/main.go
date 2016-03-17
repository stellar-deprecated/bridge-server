package handlers

import (
	"net/http"

	"github.com/stellar/gateway/compliance/config"
	"github.com/stellar/gateway/db"
	"github.com/zenazn/goji/web"
)

type RequestHandler struct {
	Config        *config.Config
	EntityManager db.EntityManagerInterface
	Repository    db.RepositoryInterface
}

func (rh *RequestHandler) HandlerSendPayment(c web.C, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("send_payment"))
}

func (rh *RequestHandler) HandlerReceivePayment(c web.C, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("receive_payment"))
}
