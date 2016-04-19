# bridge-server

This suite consists of the following apps:

* [`bridge`](./readme_bridge.md) - builds, submits and monitors transaction in Stellar network,
* [`compliance`](./readme_compliance.md) - helper server for using [Compliance protocol](https://www.stellar.org/developers/learn/integration-guides/compliance-protocol.html).

More information about each server can be found in corresponding README file.

## Downloading the server
[Prebuilt binaries](https://github.com/stellar/bridge-server/releases) of the bridge-server server are available on the 
[releases page](https://github.com/stellar/bridge-server/releases).

| Platform       | Binary file name                                                                         |
|----------------|------------------------------------------------------------------------------------------|
| Mac OSX 64 bit | [name-darwin-amd64](https://github.com/stellar/bridge-server/releases)      |
| Linux 64 bit   | [name-linux-amd64](https://github.com/stellar/bridge-server/releases)       |
| Windows 64 bit | [name-windows-amd64.exe](https://github.com/stellar/bridge-server/releases) |

Alternatively, you can [build](#building) the binary yourself.

## Building

[gb](http://getgb.io) is used for building and testing.

Given you have a running golang installation, you can build the server with:

```
gb build
```

After a successful build, you should find `bin/bridge` in the project directory.

## Running tests

```
gb test
```

## Documentation

```
godoc -goroot=. -http=:6060
```

Then simply open:
```
http://localhost:6060/pkg/github.com/stellar/gateway/
```
in a browser.
