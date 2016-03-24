package build

import (
	"errors"

	"github.com/stellar/go-stellar-base/keypair"
	"github.com/stellar/go-stellar-base/xdr"
)

func setAccountId(addressOrSeed string, aid *xdr.AccountId) error {
	kp, err := keypair.Parse(addressOrSeed)
	if err != nil {
		return err
	}

	return aid.SetAddress(kp.Address())
}

func createAlphaNumAsset(code, issuerAccountId string) (xdr.Asset, error) {
	var issuer xdr.AccountId
	err := setAccountId(issuerAccountId, &issuer)
	if err != nil {
		return xdr.Asset{}, err
	}

	length := len(code)
	switch {
	case length >= 1 && length <= 4:
		var codeArray [4]byte
		byteArray := []byte(code)
		copy(codeArray[:], byteArray[0:length])
		asset := xdr.AssetAlphaNum4{codeArray, issuer}
		return xdr.NewAsset(xdr.AssetTypeAssetTypeCreditAlphanum4, asset)
	case length >= 5 && length <= 12:
		var codeArray [12]byte
		byteArray := []byte(code)
		copy(codeArray[:], byteArray[0:length])
		asset := xdr.AssetAlphaNum12{codeArray, issuer}
		return xdr.NewAsset(xdr.AssetTypeAssetTypeCreditAlphanum12, asset)
	default:
		return xdr.Asset{}, errors.New("Asset code length is invalid")
	}
}
