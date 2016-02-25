# Changelog

As this project is pre 1.0, breaking changes may happen for minor version bumps. A breaking change will get clearly notified in this log.

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
