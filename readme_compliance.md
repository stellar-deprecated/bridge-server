# compliance-server
This is a stand alone server written in go. It is designed to make [Compliance protocol](https://www.stellar.org/developers/learn/integration-guides/compliance-protocol.html) requests to other organizations. You can connect to it from the `bridge` server or any other server that can talk to it (check API section).

## Downloading the server
[Prebuilt binaries](https://github.com/stellar/bridge-server/releases) of the bridge-server server are available on the [releases page](https://github.com/stellar/bridge-server/releases).

| Platform       | Binary file name                                                                         |
|----------------|------------------------------------------------------------------------------------------|
| Mac OSX 64 bit | [bridge-darwin-amd64](https://github.com/stellar/bridge-server/releases)      |
| Linux 64 bit   | [bridge-linux-amd64](https://github.com/stellar/bridge-server/releases)       |
| Windows 64 bit | [bridge-windows-amd64.exe](https://github.com/stellar/bridge-server/releases) |

Alternatively, you can [build](#building) the binary yourself.

## Config

The `config.toml` file must be present in a working directory. Config file should contain following values:

* `external_port` - external server listening port (should be accessible from public)
* `internal_port` - internal server listening port (should be accessible from your internal network only!)
* `needs_auth` - set to `true` if you need to do sanctions check for payment receiver
* `network_passphrase` - passphrase of the network that will be used with this bridge server:
   * test network: `Test SDF Network ; September 2015`
   * public network: `Public Global Stellar Network ; September 2015`
* `database`
  * `type` - database type (mysql, postgres)
  * `url` - url to database connection
* `keys`
  * `signing_seed` - The secret seed that will be used to sign messages. Public key derived from this secret key should be in your `stellar.toml` file.
  * `encryption_key` - The secret key used to decrypt messages. _Not working yet._
* `callbacks`
  * `sanctions` - Callback that performs sanctions check. Read [Callbacks](#callbacks) section.
  * `ask_user` - Callback that asks user for permission for reading their data. Read [Callbacks](#callbacks) section.
  * `fetch_info` - Callback that returns user data.

Check [`config-example.toml`](./config-example.toml).

## Getting started

After creating `config.toml` file, you need to run DB migrations:
```
./compliance --migrate-db
```

Then you can start the server:
```
./compliance
```

## API

`Content-Type` of requests data should be `application/x-www-form-urlencoded`.

### POST :external_port/ (Auth endpoint)

Process auth request from external organization sent before sending a payment. Check [Compliance protocol](https://www.stellar.org/developers/learn/integration-guides/compliance-protocol.html) for more info. It also saves memo preimage to the database.

#### Request Parameters

name |  | description
--- | --- | ---
`data` | required | Auth data.
`sig` | required | Signature.

Read more in [Compliance protocol](https://www.stellar.org/developers/learn/integration-guides/compliance-protocol.html#auth_server) doc.

#### Response

Returns [Auth response](https://www.stellar.org/developers/learn/integration-guides/compliance-protocol.html#reply).

### POST :internal_port/send

Sends Auth request to another organization.

#### Request Parameters

name |  | description
--- | --- | ---
`source` | required | Account ID of transaction source account.
`sender` | required | Stellar address (ex. `bob*stellar.org`) of payment sender account.
`destination` | required | Account ID or Stellar address (ex. `bob*stellar.org`) of payment destination account
`amount` | required | Amount that destination will receive
`extra_memo` | optional | Additional information attached to memo preimage.
`asset_code` | optional | Asset code (XLM when empty) destination will receive
`asset_issuer` | optional | Account ID of asset issuer (XLM when empty) destination will receive
`send_max` | optional | [path_payment] Maximum amount of send_asset to send
`send_asset_code` | optional | [path_payment] Sending asset code (XLM when empty)
`send_asset_issuer` | optional | [path_payment] Account ID of sending asset issuer (XLM when empty)
`path[n][asset_code]` | optional | [path_payment] If the path isn't specified the bridge server will find the path for you. Asset code of `n`th asset on the path (XLM when empty, but empty parameter must be sent!)
`path[n][asset_issuer]` | optional | [path_payment] Account ID of `n`th asset issuer (XLM when empty, but empty parameter must be sent!)
`path[n+1][asset_code]` | optional | [path_payment] Asset code of `n+1`th asset on the path (XLM when empty, but empty parameter must be sent!)
`path[n+1][asset_issuer]` | optional | [path_payment] Account ID of `n+1`th asset issuer (XLM when empty, but empty parameter must be sent!)
... | ... | _Up to 5 assets in the path..._

#### Response

Returns [`SendResponse`]().

### POST :internal_port/receive

Returns memo preimage.

#### Request Parameters

name |  | description
--- | --- | ---
`memo` | required | Memo hash.

#### Response

Returns [`ReceiveResponse`]().

### POST :internal_port/allow_access

Allows access to users data for external user or FI.

#### Request Parameters

name |  | description
--- | --- | ---
`name` | required | Name of the external FI.
`domain` | required | Domain of the external FI.
`public_key` | required | Public key of the external FI.
`user_id` | optional | If set, only this user will be allowed.

#### Response

Will response with `200 OK` if saved. Any other status is an error.

### POST :internal_port/remove_access

Allows access to users data for external user or FI.

#### Request Parameters

name |  | description
--- | --- | ---
`domain` | required | Domain of the external FI.
`user_id` | optional | If set, only this user entry will be removed.

#### Response

Will response with `200 OK` if removed. Any other status is an error.

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
