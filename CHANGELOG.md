# Changelog

As this project is pre 1.0, breaking changes may happen for minor version bumps. A breaking change will get clearly notified in this log.

## 0.0.10

* Send only relevant data to compliance callbacks (#17).
* `hooks` are now called `callbacks` in `bridge` server.

## 0.0.9

* Transaction builder (#14)

## 0.0.8

* [Compliance protocol](https://www.stellar.org/developers/learn/integration-guides/compliance-protocol.html) support.
* Saving and reading memo preimage.
* This repo will now contain two apps: `bridge` (for building, submitting and monitoring transactions) and `compliance` (for Compliance protocol). Both are built in a single build process. Each app has it's own README file.
* Dependency injection is now done using [facebookgo/inject](https://godoc.org/github.com/facebookgo/inject).
* Handling and validation of requests and responses is now done in `protocols` package. This package contains methods for transforming `url.Values` from/to request structs and for marshalling responses. It also contains common errors (missing/invalid fields, internal server error, etc.) and all protocol-specific error responses. It also includes stellar.toml and federation resolving.
* New `net` and `server` packages that contain some helper network connected functions and structs.
* Improvements to `db` package.

## 0.0.7

* Add path payments,
* Change `config.toml` file structure,
* Partial implementation of Compliance Protocol.

## 0.0.6

* When there are no `ReceivePayment`s in database, payment listener will start with cursor `now`.
* Fix a bug in `db.Repository.GetLastCursorValue()`.

## 0.0.5

* Add `MEMO_HASH` support.

## 0.0.4

* Fixed bugs connected with running server using `postgres` DB (full refactoring of `db` package),
* Fixed starting a minimum server with a single endpoint: `/payment`.

## 0.0.3

* Send `create_account` operation in `/payment` if account does not exist.
* Fixed major bug in `PaymentListener`.
* Sending to Stellar address with memo in `/send`.
* Standardized responses.
* Updated README file.

## 0.0.2

* Added `/payment` endpoint.
* Now it's possible to start a server with parameter that are not required. Minimum version starts a server with a single endpoint: `/payment`.
* Added config parameters validation.
* Added `network_passphrase` config parameter.
* `postgres` migration files.
* Fixed sending to Stellar address.
* Fixed `horizon.AccountResponse.SequenceNumber` bug.
* Fixed minor bugs.
* Code refactoring.
* Added example config file to the release package
* Updated README file.

## 0.0.1

* Initial release.
