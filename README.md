# gateway-server
This is a stand alone server written in go. It is designed to makes connecting to the Stellar network as easy as possible.
It assumes you are using only one issuing account.  It requires you to implement a recieve web hook that is notified when there is an incoming payment. 

Handles:

- Creating the Stellar transactions.
- Monitoring the receiving Stellar account.


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
* `network_passphrase` - passphrase of the network that will be used with this gateway server:
   * test network: `Test SDF Network ; September 2015`
   * public network: `Public Global Stellar Network ; September 2015`
* `horizon` - URL to [horizon](https://github.com/stellar/horizon) server instance
* `assets` - array of approved assets codes that this server can authorize or receive. These are currency code/issuer pairs. 
* `database`
  * `type` - database type (mysql, postgres)
  * `url` - url to database connection
* `accounts`
  * `authorizing_seed` - The secret seed of the public key that is able to submit `allow_trust` operations on the issuing account.
  * `receiving_account_id` - The account ID that receives incoming payments. The `receive hook` will be called when a payment is received by this account.
* `hooks`
  * `receive` - URL of the webhook where requests will be sent when a new payment appears in receiving account. **WARNING** Gateway server can send multiple requests to this webhook for a single payment! You need to be prepared for it. See: [Security](#security).
  * `error` - URL of the webhook where requests will be sent when there is an error with incoming payment

Check [`config-example.toml`](./config-example.toml).

The minimal set of config values contains:
* `port`
* `network_passphrase`
* `horizon`

It will start a server with a single endpoint: `/payment`.

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

Builds and submits a transaction with a single [`payment`](https://www.stellar.org/developers/learn/concepts/list-of-operations.html#payment), [`path_payment`](https://www.stellar.org/developers/learn/concepts/list-of-operations.html#path-payment) or [`create_account`](https://www.stellar.org/developers/learn/concepts/list-of-operations.html#create-account) (when sending native asset to account that does not exist) operation built from following parameters.

#### Request Parameters

Every request must contain required parameters from the following list. Additionally, depending on `type` parameter, every request must contain required parameters for equivalent operation type (check tables below).

name |  | description
--- | --- | ---
`source` | required | Secret seed of transaction source account
`destination` | required | Account ID or Stellar address (ex. `bob*stellar.org`) of payment destination account
`type` | optional | Operation type: `payment` (default) or `path_payment`.
`memo_type` | optional | Memo type, one of: `id`, `text`, `hash`
`memo` | optional | Memo value, when `memo_type` is `id` it must be uint64, when `hash` it must be 32 bytes hex value

##### CreateAccount / Payment operation parameters

name |  | description
--- | --- | ---
`amount` | required | Amount to send
`asset_code` | optional | Asset code (XLM when empty)
`asset_issuer` | optional | Account ID of asset issuer (XLM when empty)

##### PathPayment operation parameters

name |  | description
--- | --- | ---
`send_max` | required | Maximum amount of send_asset to send
`send_asset_code` | optional | Sending asset code (XLM when empty)
`send_asset_issuer` | optional | Account ID of sending asset issuer (XLM when empty)
`destination_amount` | required | Amount of destination_asset to receiving account will get
`destination_asset_code` | optional | Destination asset code (XLM when empty)
`destination_asset_issuer` | optional | Account ID of destination asset issuer (XLM when empty)
`path[n][asset_code]` | optional | Asset code of `n`th asset on the path (XLM when empty, but empty parameter must be sent!)
`path[n][asset_issuer]` | optional | Account ID of `n`th asset issuer (XLM when empty, but empty parameter must be sent!)
`path[n+1][asset_code]` | optional | Asset code of `n+1`th asset on the path (XLM when empty, but empty parameter must be sent!)
`path[n+1][asset_issuer]` | optional | Account ID of `n+1`th asset issuer (XLM when empty, but empty parameter must be sent!)
... | ... | _Up to 5 assets in the path..._

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

### POST /authorize

Builds and submits a transaction with a [`allow_trust`](https://www.stellar.org/developers/learn/concepts/list-of-operations.html#allow-trust) operation. The source of this transaction will be the account specified by `accounts.authorizing_seed` config parameter.

#### Request Parameters

name |  | description
--- | --- | ---
`account_id` | required | Account ID of the account to authorize
`asset_code` | required | Asset code of the asset to authorize. Must be present in `assets` config array.

#### Response

Check [`SubmitTransactionResponse`](./src/github.com/stellar/gateway/horizon/submit_transaction_response.go) struct.

### POST /send

Builds and submits a transaction with a [`payment`](https://www.stellar.org/developers/learn/concepts/list-of-operations.html#payment) operation. The source of this transaction will be the account specified by `accounts.issuing_seed` config parameter.

#### Request Parameters

name |  | description
--- | --- | ---
`destination` | required | Account ID or Stellar address (ex. `bob*stellar.org`) of the destination account
`asset_code` | required | Asset code of the asset to send. Must be present in `assets` config array.
`amount` | required | Amount to send.
`memo_type` | optional | Memo type, one of: `id`, `text`, `hash`
`memo` | optional | Memo value, when `memo_type` is `id` it must be uint64, when `hash` it must be 32 bytes hex value

#### Response

Check [`SubmitTransactionResponse`](./src/github.com/stellar/gateway/horizon/submit_transaction_response.go) struct.

## Hooks

Gateway server listens for payment operations to the account specified by `accounts.receiving_account_id`. Every time a payment arrives it will send a HTTP POST request to `hooks.receive`.

`Content-Type` of requests data will be `application/x-www-form-urlencoded`.

### `hooks.receive`

The POST request with following parameters will be sent to this hook when a payment arrives.

> **Warning!** This hook can be called multiple times. Please check `id` parameter and respond with `200 OK` in case of duplicate payment.

#### Request

name | description
--- | ---
`id` | Operation ID
`from` | Account ID of the sender
`amount` | Amount that was sent
`asset_code` | Code of the asset sent (ex. `USD`)
`memo_type` | Type of the memo attached to the transaction. This field will be empty when no memo was attached.
`memo` | Value of the memo attached. This field will be empty when no memo was attached.

#### Response

Response with `200 OK` when processing succeeded. Any other status code will be considered an error.

## Security

* This server must be set up in an isolated environment (ex. AWS VPC). Please make sure your firewall is properly configured and accepts connections from a trusted IPs only. You can also set `api_key` config parameter but it's not recommended. If you will not set this properly, an unauthorized person will be able to submit transactions from your accounts!
* Make sure `hooks` accepts connections from the gateway server IP only.
* Remember that `hooks.receive` may be called multiple times with the same payment. Check `id` parameter and ignore requests with the same value (just send `200 OK` response).

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
