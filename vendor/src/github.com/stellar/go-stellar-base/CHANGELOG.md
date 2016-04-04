# Changelog

All notable changes to this project will be documented in this
file.  This project adheres to [Semantic Versioning](http://semver.org/).

As this project is pre 1.0, breaking changes may happen for minor version
bumps.  A breaking change will get clearly notified in this log.

## [Unreleased]

### Added
- Added `EntryType()` to `LedgerEntryChange`
- Many xdr types learned `Equals()`, a helper to work around deficiencies when comparing xdr unions.
- The `amount` package learned `MustParse(string)`.
- `*xdr.Price` learned `String()`.
- `*xdr.PathPaymentResult` learned `SendAmount()`, a helper to extract how much of the source asset was spent when making a multi-asset payment.

- `*xdr.AccountId` learned `Address()` to make getting the strkey form of an account id simpler.
- `build` package learned `ClearData()` and `SetData()` to configure ManageData operations.


[Unreleased]: https://github.com/stellar/go-stellar-base/compare/df92a863a...master
