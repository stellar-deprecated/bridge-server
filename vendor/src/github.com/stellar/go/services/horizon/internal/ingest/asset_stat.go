package ingest

import (
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/stellar/go/services/horizon/internal/db2/core"
	"github.com/stellar/go/services/horizon/internal/db2/history"
	"github.com/stellar/go/support/db"
	"github.com/stellar/go/xdr"
)

func (assetsModified AssetsModified) handlePaymentOp(paymentOp *xdr.PaymentOp, sourceAccount *xdr.AccountId) error {
	err := assetsModified.updateIfAssetIssuerInvolved(paymentOp.Asset, *sourceAccount)
	if err != nil {
		return err
	}
	return assetsModified.updateIfAssetIssuerInvolved(paymentOp.Asset, paymentOp.Destination)
}

func defaultSourceAccount(sourceAccount *xdr.AccountId, defaultAccount *xdr.AccountId) *xdr.AccountId {
	if sourceAccount != nil {
		return sourceAccount
	}
	return defaultAccount
}

func (assetsModified AssetsModified) add(asset xdr.Asset) {
	assetsModified[asset.String()] = asset
}

// IngestOperation updates the assetsModified using the passed in operation
func (assetsModified AssetsModified) IngestOperation(err error, op *xdr.Operation, source *xdr.AccountId, coreQ *core.Q) error {
	if err != nil {
		return err
	}

	body := op.Body
	sourceAccount := defaultSourceAccount(op.SourceAccount, source)
	switch body.Type {
	// TODO NNS 2 need to fix GetCreateAssetID call when adding assets from account
	// case xdr.OperationTypeSetOptions:
	// 	assetsModified.addAssetsFromAccount(coreQ, sourceAccount)
	case xdr.OperationTypePayment:
		// payments is the only operation where we currently perform the optimization of checking against the issuer
		return assetsModified.handlePaymentOp(body.PaymentOp, sourceAccount)
	case xdr.OperationTypePathPayment:
		// if this gets expensive then we can limit it to only include those assets that includes the issuer
		assetsModified.add(body.PathPaymentOp.DestAsset)
		assetsModified.add(body.PathPaymentOp.SendAsset)
		for _, asset := range body.PathPaymentOp.Path {
			assetsModified.add(asset)
		}
	case xdr.OperationTypeManageOffer:
		// if this gets expensive then we can limit it to only include those assets that includes the issuer
		assetsModified.add(body.ManageOfferOp.Buying)
		assetsModified.add(body.ManageOfferOp.Selling)
	case xdr.OperationTypeCreatePassiveOffer:
		// if this gets expensive then we can limit it to only include those assets that includes the issuer
		assetsModified.add(body.CreatePassiveOfferOp.Buying)
		assetsModified.add(body.CreatePassiveOfferOp.Selling)
	case xdr.OperationTypeChangeTrust:
		assetsModified.add(body.ChangeTrustOp.Line)
	case xdr.OperationTypeAllowTrust:
		asset := body.AllowTrustOp.Asset.ToAsset(*sourceAccount)
		assetsModified.add(asset)
	}

	return nil
}

// UpdateAssetStats updates the db with the latest asset stats for the assets that were modified
func (assetsModified AssetsModified) UpdateAssetStats(is *Session) {
	if is.Err != nil {
		return
	}

	hasValue := false
	for _, asset := range assetsModified {
		assetStat := computeAssetStat(is, &asset)
		if is.Err != nil {
			return
		}

		if assetStat != nil {
			hasValue = true
			is.Ingestion.assetStats = is.Ingestion.assetStats.Values(
				assetStat.ID,
				assetStat.Amount,
				assetStat.NumAccounts,
				assetStat.Flags,
				assetStat.Toml,
			)
		}
	}

	if hasValue {
		// perform a delete first since upsert is not supported if postgres < 9.5
		is.Err = assetsModified.deleteRows(is.Ingestion.DB)
		if is.Err != nil {
			return
		}

		// can perform a direct upsert if postgres > 9.4
		// is.Ingestion.assetStats = is.Ingestion.assetStats.
		// 	Suffix("ON CONFLICT (id) DO UPDATE SET (amount, num_accounts, flags, toml) = (excluded.amount, excluded.num_accounts, excluded.flags, excluded.toml)")
		_, is.Err = is.Ingestion.DB.Exec(is.Ingestion.assetStats)
	}
}

// func (assetsModified AssetsModified) addAssetsFromAccount(coreQ *core.Q, account *xdr.AccountId) {
// 	if account == nil {
// 		return
// 	}

// 	var assets []xdr.Asset
// 	coreQ.AssetsForAddress(&assets, account.Address())

// 	for _, asset := range assets {
// 		if asset.Type != xdr.AssetTypeAssetTypeNative {
// 			assetsModified.add(asset)
// 		}
// 	}
// }

func (assetsModified AssetsModified) deleteRows(session *db.Session) error {
	if len(assetsModified) == 0 {
		return nil
	}

	assets := make([]xdr.Asset, 0, len(assetsModified))
	for _, asset := range assetsModified {
		assets = append(assets, asset)
	}
	historyQ := history.Q{Session: session}
	ids, err := historyQ.GetAssetIDs(assets)
	if err != nil {
		return err
	}
	if len(ids) == 0 {
		return nil
	}

	deleteStmt := sq.Delete("asset_stats").Where(sq.Eq{"id": ids})
	_, err = session.Exec(deleteStmt)
	return err
}

func (assetsModified AssetsModified) updateIfAssetIssuerInvolved(asset xdr.Asset, account xdr.AccountId) error {
	var assetType, assetCode, assetIssuer string
	err := asset.Extract(&assetType, &assetCode, &assetIssuer)
	if err != nil {
		return err
	}

	if assetIssuer == account.Address() {
		assetsModified.add(asset)
	}
	return nil
}

func computeAssetStat(is *Session, asset *xdr.Asset) *history.AssetStat {
	if asset.Type == xdr.AssetTypeAssetTypeNative {
		return nil
	}

	historyQ := history.Q{Session: is.Ingestion.DB}
	assetID, err := historyQ.GetCreateAssetID(*asset)
	if err != nil {
		is.Err = err
		return nil
	}

	var assetType xdr.AssetType
	var assetCode, assetIssuer string
	err = asset.Extract(&assetType, &assetCode, &assetIssuer)
	if err != nil {
		is.Err = err
		return nil
	}

	coreQ := &core.Q{Session: is.Cursor.DB}

	numAccounts, amount, err := statTrustlinesInfo(coreQ, assetType, assetCode, assetIssuer)
	if err != nil {
		is.Err = err
		return nil
	}

	flags, toml, err := statAccountInfo(coreQ, assetIssuer)
	if err != nil {
		is.Err = err
		return nil
	}

	return &history.AssetStat{
		ID:          assetID,
		Amount:      amount,
		NumAccounts: numAccounts,
		Flags:       flags,
		Toml:        toml,
	}
}

// statTrustlinesInfo fetches all the stats from the trustlines table
func statTrustlinesInfo(coreQ *core.Q, assetType xdr.AssetType, assetCode string, assetIssuer string) (int32, int64, error) {
	return coreQ.BalancesForAsset(int32(assetType), assetCode, assetIssuer)
}

// statAccountInfo fetches all the stats from the accounts table
func statAccountInfo(coreQ *core.Q, accountID string) (int8, string, error) {
	var account core.Account
	err := coreQ.AccountByAddress(&account, accountID)
	if err != nil {
		return -1, "", err
	}

	var toml string
	if !account.HomeDomain.Valid {
		toml = ""
	} else {
		trimmed := strings.TrimSpace(account.HomeDomain.String)
		if trimmed != "" {
			toml = "https://" + account.HomeDomain.String + "/.well-known/stellar.toml"
		} else {
			toml = ""
		}
	}

	return int8(account.Flags), toml, nil
}
