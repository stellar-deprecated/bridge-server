package compliance

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	complianceProtocol "github.com/stellar/go/protocols/compliance"
	"github.com/stellar/go/support/errors"
)

// DefaultPendingTime is a value (in seconds) that will be returned in
// `pending` response if user callback haven't returned a valid response.
const DefaultPendingTime = 600

type pendingResponse struct {
	Pending int `json:"pending"`
}

// SanctionsCheck performs AML sanctions check of the sender.
func (s *CallbackStrategy) SanctionsCheck(data complianceProtocol.AuthData, response *complianceProtocol.AuthResponse) (err error) {
	if s.SanctionsCheckURL == "" {
		response.TxStatus = complianceProtocol.AuthStatusOk
		return
	}

	resp, err := http.PostForm(s.SanctionsCheckURL, url.Values{"attachment": {data.AttachmentJSON}})
	if err != nil {
		err = errors.Wrap(err, "Error connecting sanctions server")
		return
	}

	switch resp.StatusCode {
	case http.StatusOK: // AuthStatusOk
		response.TxStatus = complianceProtocol.AuthStatusOk
	case http.StatusAccepted: // AuthStatusPending
		response.TxStatus = complianceProtocol.AuthStatusPending

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrap(err, "Error reading sanctions server response")
		}

		pendingResponse := pendingResponse{}
		err = json.Unmarshal(body, &pendingResponse)
		if err != nil {
			// Set default value
			response.Pending = DefaultPendingTime
		} else {
			response.Pending = pendingResponse.Pending
		}
	case http.StatusForbidden: // AuthStatusDenied
		response.TxStatus = complianceProtocol.AuthStatusDenied
	default:
		err = fmt.Errorf("Invalid status code from sanctions server: %d", resp.StatusCode)
		return
	}

	return
}

// GetUserData check if user data is required and if so decides
// whether to allow access to customer data or not.
func (s *CallbackStrategy) GetUserData(data complianceProtocol.AuthData, response *complianceProtocol.AuthResponse) (err error) {
	// If sender doesn't need info, return AuthStatusOk
	if !data.NeedInfo {
		response.InfoStatus = complianceProtocol.AuthStatusOk
		return
	}

	// If there is no way to fetch data, return AuthStatusDenied
	if s.GetUserDataURL == "" {
		response.InfoStatus = complianceProtocol.AuthStatusDenied
		return
	}

	resp, err := http.PostForm(s.GetUserDataURL, url.Values{"attachment": {data.AttachmentJSON}})
	if err != nil {
		err = errors.Wrap(err, "Error connecting fetch info server")
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.Wrap(err, "Error reading fetch info server response")
		return
	}

	switch resp.StatusCode {
	case http.StatusOK: // AuthStatusOk
		response.TxStatus = complianceProtocol.AuthStatusOk
		response.DestInfo = string(body)
	case http.StatusAccepted: // AuthStatusPending
		response.TxStatus = complianceProtocol.AuthStatusPending

		var pending int
		pendingResponse := pendingResponse{}
		err := json.Unmarshal(body, &pendingResponse)
		if err != nil {
			// Set default value
			pending = 600
		} else {
			pending = pendingResponse.Pending
		}

		// Check if SanctionsCheck pending time is smaller
		if pending > response.Pending {
			response.Pending = pending
		}
	case http.StatusForbidden: // AuthStatusDenied
		response.TxStatus = complianceProtocol.AuthStatusDenied
	default:
		err = fmt.Errorf("Invalid status code from fetch info server: %d", resp.StatusCode)
		return
	}

	return
}
