package handlers

import (
	"net/http"
)

func (rh *RequestHandler) Send(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	r.PostForm.Set("source", *rh.Config.Accounts.IssuingSeed)
	rh.Payment(w, r)
}
