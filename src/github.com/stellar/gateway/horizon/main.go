package horizon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
    "net/http"
    "net/url"
)

type Horizon struct {
	ServerUrl string
}

func (h *Horizon) LoadAccount(accountId string) (response AccountResponse, err error) {
	resp, err := http.Get(h.ServerUrl+"/accounts/"+accountId)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if (resp.StatusCode != 200) {
		err = fmt.Errorf("StatusCode indicates error: %s", body)
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	return
}

func (h *Horizon) SubmitTransaction(txeBase64 string) (response SubmitTransactionResponse, err error) {
    v := url.Values{}
    v.Set("tx", txeBase64)
    // TODO add request timeout
	resp, err := http.PostForm(h.ServerUrl+"/transactions", v)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	log.Print("Response from horizon ", string(body))

	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	return
}
