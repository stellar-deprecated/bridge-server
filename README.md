# gateway-server

## Downloading the server
[Prebuilt binaries](https://github.com/stellar/gateway-server/releases) of the gateway-server server are available on the [releases page](https://github.com/stellar/gateway-server/releases).

| Platform       | Binary file name                                                                         |
|----------------|------------------------------------------------------------------------------------------|
| Mac OSX 64 bit | [gateway-darwin-amd64](https://github.com/stellar/gateway-server/releases)      |
| Linux 64 bit   | [gateway-linux-amd64](https://github.com/stellar/gateway-server/releases)       |
| Windows 64 bit | [gateway-windows-amd64.exe](https://github.com/stellar/gateway-server/releases) |

Alternatively, you can [build](#building) the binary yourself.

## Config

The `config.toml` file must be present in a working directory. Config file should contain following values:

* `port` - server listening port
* `api_key` - when set, all requests to gateway server must contain `api_key` parameter with a correct value, otherwise the server will respond with `503 Forbidden`
* `network_passphrase` - passphrase of the network that will be used with this gateway server, default: `Test SDF Network ; September 2015`
* `horizon` - URL to [horizon](https://github.com/stellar/horizon) server instance
* `assets` - array of approved assets codes that this server can authorize and send 
* `database`
  * `type` - database type (sqlite3, mysql, postgres)
  * `url` - url to database connection
* `accounts`
  * `authorizing_seed` - secret seed of the account to send `allow_trust` operations
  * `issuing_seed` - secret seed of the account to send `payment` operations
  * `receiving_account_id` - ID of the account to track incoming payments
* `hooks`
  * `receive` - URL of the webhook where requests will be sent when a new payment appears in receiving account. **WARNING** Gateway server can send multiple requests to this webhook for a single payment! You need to be prepared for it. See: [Security](#security).
  * `error` - URL of the webhook where requests will be sent when there is an error with incoming payment

Check [`config-example.toml`](./config-example.toml).

## Getting started

After creating `config.toml` file, you need to run DB migrations:
```
./gateway --migrate-db
```

Then you can start the server:
```
./gateway
```

## API

`Content-Type` of requests data should be `application/x-www-form-urlencoded`.

### POST /payment

Builds and submits a transaction with a single payment operation built from following parameters.

#### Request Parameters

name |  | description
--- | --- | ---
`source` | required | Secret seed of transaction source account
`destination` | required | Account ID of payment destination account
`amount` | required | Amount to send
`asset_code` | optional | Asset code (XLM when empty)
`asset_issuer` | optional | Account ID of asset issuer (XLM when empty)
`memo_type` | optional | Memo type, one of: `id`, `text`
`memo` | optional | Memo value, when `memo_type` is `id` it must be uint64

#### Response

Check [`SubmitTransactionResponse`](./src/github.com/stellar/gateway/horizon/submit_transaction_response.go) struct.

#### Example

```sh
curl -X POST -d \
"source=SBNDIK4N7ZM3ZJKDJJDWDSPSRPHNI2RFL36WNNNEGQEW3G3AH6VJ2QB7&\
amount=1&\
destination=GBIUXI4S27PSL6TTJCJMPYDCF3K6AW2MYORFRTC7QBFE6NNEGVOQK46H&\
asset_code=USD&\
asset_issuer=GASZUHRFAFIZX5LR4WNHBWUXJBZNBEWCHFTR4XZHPF5TMVM5XUZBP5DT&\
memo_type=id&\
memo=125" \
http://localhost:8001/payment
```

## Hooks

//

## Security

//

## Building

[gb](http://getgb.io) is used for building and testing.

Given you have a running golang installation, you can build the server with:

```
gb build
```

After successful completion, you should find `bin/gateway` is present in the project directory.

## Running tests

```
gb test
```
