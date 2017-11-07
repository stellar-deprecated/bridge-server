# Horizon
[![Build Status](https://travis-ci.org/stellar/horizon.svg?branch=master)](https://travis-ci.org/stellar/horizon)

Horizon is the [client facing API](/docs) server for the Stellar ecosystem.  It acts as the interface between stellar-core and applications that want to access the Stellar network. It allows you to submit transactions to the network, check the status of accounts, subscribe to event streams, etc. See [an overview of the Stellar ecosystem](https://www.stellar.org/developers/guides/get-started/) for more details.

## Downloading the server
[Prebuilt binaries](https://github.com/stellar/go/releases) of horizon are available on the 
[releases page](https://github.com/stellar/go/releases).

See [the old releases page](https://github.com/stellar/horizon/releases) for prior releases

| Platform       | Binary file name                                                                         |
|----------------|------------------------------------------------------------------------------------------|
| Mac OSX 64 bit | [horizon-darwin-amd64](https://github.com/stellar/go/releases/download/v0.11.0/horizon-v0.11.0-darwin-amd64.tar.gz)      |
| Linux 64 bit   | [horizon-linux-amd64](https://github.com/stellar/go/releases/download/v0.11.0/horizon-v0.11.0-linux-amd64.tar.gz)       |
| Windows 64 bit | [horizon-windows-amd64.exe](https://github.com/stellar/go/releases/download/v0.11.0/horizon-v0.11.0-windows-amd64.zip) |

Alternatively, you can [build](#building) the binary yourself.

## Dependencies

Horizon requires go 1.6 or higher to build. See (https://golang.org/doc/install) for installation instructions.

## Building

[glide](https://glide.sh/) is used for building horizon.

Given you have a running golang installation, you can install this with:

```bash
curl https://glide.sh/get | sh
```

Next, you must download the source for packages that horizon depends upon. From within the project directory, run:

```bash
glide install
```

Then, simply run `go install github.com/stellar/go/services/horizon`.  After successful
completion, you should find `horizon` is present in your `$GOPATH/bin` directory.

More detailed intructions and [admin guide](internal/docs/reference/admin.md). 

## Developing Horizon

See [the development guide](internal/docs/developing.md).

## Contributing
Please see the [CONTRIBUTING.md](./CONTRIBUTING.md) for details on how to contribute to this project.
