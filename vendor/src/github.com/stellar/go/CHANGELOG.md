# Changelog

All notable changes to this project will be documented in this
file.  This project adheres to [Semantic Versioning](http://semver.org/).

As this project is pre 1.0, breaking changes may happen for minor version
bumps.  A breaking change will get clearly notified in this log.

NOTE:  this changelog represents the changes that are associated with the library code in this repo (rather than the tools or services in this repo).  

## [Unreleased]

### Added

- xdr: added support for new signer types
- build: `Signer` learned support for new signer types
- strkey: added support for new signer types
- network:  Added the `HashTransaction` helper func to get the hash of a transaction targetted to a specific stellar network.

### Changed:

- build: _BREAKING CHANGE_:  A transaction built and signed using the `build` package no longer default to the test network.

[Unreleased]: https://github.com/stellar/go/commits/master