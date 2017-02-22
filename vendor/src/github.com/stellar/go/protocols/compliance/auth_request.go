package compliance

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/stellar/go/clients/stellartoml"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/support/errors"
)

func (r *AuthRequest) Populate(request *http.Request) *AuthRequest {
	r.DataJSON = request.PostFormValue("data")
	r.Signature = request.PostFormValue("sig")
	return r
}

// Validate checks if fields are valid:
//
//   * `Data` and `Signature` not empty
//   * `Data` is JSON and pass AuthData.Validate method
//
// This method only performs data validation. You should also call
// VerifySignature to confirm that signature is valid.
func (r *AuthRequest) Validate() error {
	if r.DataJSON == "" || r.Signature == "" {
		return errors.New("`data` and `signature` fields are required")
	}

	authData := AuthData{}
	err := json.Unmarshal([]byte(r.DataJSON), &authData)
	if err != nil {
		return errors.Wrap(err, "Data is not valid JSON")
	}

	// Validate Signature
	_, err = base64.StdEncoding.DecodeString(r.Signature)
	if err != nil {
		return errors.New("Signature is not base64 encoded")
	}

	// Validate DataJSON
	err = authData.Validate()
	if err != nil {
		return errors.Wrap(err, "Invalid Data")
	}

	return nil
}

// VerifySignature verifies if signature is valid. It makes a network connection
// to sender server in order to obtain stellar.toml file and signing key.
func (r *AuthRequest) VerifySignature(sender string) error {
	senderStellarToml, err := stellartoml.GetStellarTomlByAddress(sender)
	if err != nil {
		return errors.Wrap(err, "Cannot get stellar.toml of sender")
	}

	if senderStellarToml.SigningKey == "" {
		return errors.New("No SIGNING_KEY in stellar.toml of sender")
	}

	kp, err := keypair.Parse(senderStellarToml.SigningKey)
	if err != nil {
		return errors.New("SigningKey is invalid")
	}

	signatureBytes, err := base64.StdEncoding.DecodeString(r.Signature)
	if err != nil {
		return errors.New("Signature is not base64 encoded")
	}

	err = kp.Verify([]byte(r.DataJSON), signatureBytes)
	if err != nil {
		return errors.New("Signature is invalid")
	}

	return nil
}

// Data returns AuthData from the request.
func (r *AuthRequest) Data() (data AuthData, err error) {
	err = json.Unmarshal([]byte(r.DataJSON), &data)
	return
}

// ToURLValues returns AuthData encoded as url.Values.
func (r *AuthRequest) ToURLValues() url.Values {
	return url.Values{
		"data": []string{r.DataJSON},
		"sig":  []string{r.Signature},
	}
}
