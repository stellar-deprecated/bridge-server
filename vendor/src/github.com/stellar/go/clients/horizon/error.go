package horizon

import (
	"encoding/json"

	"github.com/stellar/go/support/errors"
	"github.com/stellar/go/xdr"
)

func (herr *Error) Error() string {
	// TODO: use the attached problem to provide a better error message
	return "Horizon error"
}

// Envelope extracts the transaction envelope that triggered this error from the
// extra fields.
func (herr *Error) Envelope() (*xdr.TransactionEnvelope, error) {
	raw, ok := herr.Problem.Extras["envelope_xdr"]
	if !ok {
		return nil, ErrEnvelopeNotPopulated
	}

	var b64 string
	var result xdr.TransactionEnvelope

	err := json.Unmarshal(raw, &b64)
	if err != nil {
		return nil, errors.Wrap(err, "json decode failed")
	}

	err = xdr.SafeUnmarshalBase64(b64, &result)
	if err != nil {
		return nil, errors.Wrap(err, "xdr decode failed")
	}

	return &result, nil
}

// ResultCodes extracts a result code summary from the error, if possible.
func (herr *Error) ResultCodes() (*TransactionResultCodes, error) {
	if herr.Problem.Type != "transaction_failed" {
		return nil, ErrTransactionNotFailed
	}

	raw, ok := herr.Problem.Extras["result_codes"]
	if !ok {
		return nil, ErrResultCodesNotPopulated
	}

	var result TransactionResultCodes
	err := json.Unmarshal(raw, &result)
	if err != nil {
		return nil, errors.Wrap(err, "json decode failed")
	}

	return &result, nil
}
