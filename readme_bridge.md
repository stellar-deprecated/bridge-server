# bridge-server
This is a stand alone server written in go. It is designed to make connecting to the Stellar network as easy as possible. 
It allows you to be notified when a payment is received by a particular account. It also allows you to send a payment via a HTTP request.
It can optionally be connected to a `compliance` server if you want to carry out the compliance protocol.
It can be used by any project that needs to accept or send payments such as anchors or merchants accepting payments.

Handles:

- Creating Stellar transactions.
- Monitoring a receiving Stellar account.


## Downloading the server
[Prebuilt binaries](https://github.com/stellar/bridge-server/releases) of the bridge-server server are available on the [releases page](https://github.com/stellar/bridge-server/releases).

| Platform       | Binary file name                                                                         |
|----------------|------------------------------------------------------------------------------------------|
| Mac OSX 64 bit | [bridge-darwin-amd64](https://github.com/stellar/bridge-server/releases)      |
| Linux 64 bit   | [bridge-linux-amd64](https://github.com/stellar/bridge-server/releases)       |
| Windows 64 bit | [bridge-windows-amd64.exe](https://github.com/stellar/bridge-server/releases) |

Alternatively, you can [build](#building) the binary yourself.

## Config

The `config_bridge.toml` file must be present in a working directory. Config file should contain following values:

* `port` - server listening port
* `api_key` - when set, all requests to bridge server must contain `api_key` parameter with a correct value, otherwise the server will respond with `503 Forbidden`
* `network_passphrase` - passphrase of the network that will be used with this bridge server:
   * test network: `Test SDF Network ; September 2015`
   * public network: `Public Global Stellar Network ; September 2015`
* `compliance` - URL to compliance server instance if you want to carry out the compliance protocol
* `horizon` - URL to [horizon](https://github.com/stellar/horizon) server instance
* `assets` - array of approved assets codes that this server can authorize or receive. These are currency code/issuer pairs. 
* `database`
  * `type` - database type (mysql, postgres)
  * `url` - url to database connection
* `accounts`
  * `base_seed` - The secret seed of the account used to send payments. If left blank you will need to pass it in calls to `/payment`. 
  * `authorizing_seed` - The secret seed of the public key that is able to submit `allow_trust` operations on the issuing account.
  * `issuing_account_id` - The account ID of the issuing account.
  * `receiving_account_id` - The account ID that receives incoming payments. The `receive hook` will be called when a payment is received by this account.
* `hooks`
  * `receive` - URL of the webhook where requests will be sent when a new payment is sent to the receiving account. **WARNING** The bridge server can send multiple requests to this webhook for a single payment! You need to be prepared for it. See: [Security](#security).
  * `error` - URL of the webhook where requests will be sent when there is an error with an incoming payment
* `log_format` - set to `json` for JSON logs

Check [`config-example.toml`](./config-example.toml).

The minimal set of config values contains:
* `port`
* `network_passphrase`
* `horizon`

It will start a server with a single endpoint: `/payment`.

## Getting started

After creating `config_bridge.toml` file, you need to run DB migrations:
```
./bridge --migrate-db
```

Then you can start the server:
```
./bridge
```

## API

`Content-Type` of requests data should be `application/x-www-form-urlencoded`.

### POST /payment

Builds and submits a transaction with a single [`payment`](https://www.stellar.org/developers/learn/concepts/list-of-operations.html#payment), [`path_payment`](https://www.stellar.org/developers/learn/concepts/list-of-operations.html#path-payment) or [`create_account`](https://www.stellar.org/developers/learn/concepts/list-of-operations.html#create-account) (when sending native asset to account that does not exist) operation built from following parameters.

#### Request Parameters

Every request must contain required parameters from the following list. Additionally, depending on a type of payment, every request must contain required parameters for equivalent operation type.

name |  | description
--- | --- | ---
`source` | optional | Secret seed of transaction source account. If ommitted it will use the `base_seed` specified in the config file.
`sender` | optional | Stellar address (ex. `bob*stellar.org`) of payment sender account. Required for when sending using Compliance protocol.
`destination` | required | Account ID or Stellar address (ex. `bob*stellar.org`) of payment destination account
`amount` | required | Amount that destination will receive
`memo_type` | optional | Memo type, one of: `id`, `text`, `hash`, `extra`
`memo` | optional | Memo value, `id` it must be uint64, when `hash` it must be 32 bytes hex value.
`extra_memo` | optional | You can include any info here and it will be included in the pre-image of the transaction's memo hash. See the [Stellar Memo Convention](https://github.com/stellar/stellar-protocol/issues/28). When set and compliance server is connected, `memo` and `memo_type` values will be ignored.
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

It will return [`SubmitTransactionResponse`](./blob/master/src/github.com/stellar/gateway/horizon/submit_transaction_response.go) if there were no errors or with one of the following errors:

* [`InternalServerError`](./blob/master/src/github.com/stellar/gateway/protocols/errors.go)
* [`InvalidParameterError`](./blob/master/src/github.com/stellar/gateway/protocols/errors.go)
* [`MissingParameterError`](./blob/master/src/github.com/stellar/gateway/protocols/errors.go)
* [`TransactionBadSequence`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/errors.go)
* [`TransactionBadAuth`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/errors.go)
* [`TransactionInsufficientBalance`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/errors.go)
* [`TransactionNoAccount`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/errors.go)
* [`TransactionInsufficientFee`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/errors.go)
* [`TransactionBadAuthExtra`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/errors.go)
* [`PaymentCannotResolveDestination`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/payment.go)
* [`PaymentCannotUseMemo`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/payment.go)
* [`PaymentSourceNotExist`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/payment.go)
* [`PaymentAssetCodeNotAllowed`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/payment.go)
* [`PaymentPending`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/payment.go)
* [`PaymentDenied`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/payment.go)
* [`PaymentMalformed`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/payment.go)
* [`PaymentUnderfunded`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/payment.go)
* [`PaymentSrcNoTrust`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/payment.go)
* [`PaymentSrcNotAuthorized`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/payment.go)
* [`PaymentNoDestination`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/payment.go)
* [`PaymentNoTrust`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/payment.go)
* [`PaymentNotAuthorized`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/payment.go)
* [`PaymentLineFull`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/payment.go)
* [`PaymentNoIssuer`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/payment.go)
* [`PaymentTooFewOffers`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/payment.go)
* [`PaymentOfferCrossSelf`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/payment.go)
* [`PaymentOverSendmax`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/payment.go)

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
Can be used to authorize other accounts to hold your assets.
It will build and submits a transaction with a [`allow_trust`](https://www.stellar.org/developers/learn/concepts/list-of-operations.html#allow-trust) operation. 
The source of this transaction will be the account specified by `accounts.authorizing_seed` config parameter. 
You should make sure that this account is a low weight signer on the issuing account. See [Multi-sig](https://www.stellar.org/developers/learn/concepts/multi-sig.html) for more information. 

#### Request Parameters

name |  | description
--- | --- | ---
`account_id` | required | Account ID of the account to authorize
`asset_code` | required | Asset code of the asset to authorize. Must be present in `assets` config array.

#### Response

It will return [`SubmitTransactionResponse`](./blob/master/src/github.com/stellar/gateway/horizon/submit_transaction_response.go) if there were no errors or with one of the following errors:

* [`InternalServerError`](./blob/master/src/github.com/stellar/gateway/protocols/errors.go)
* [`InvalidParameterError`](./blob/master/src/github.com/stellar/gateway/protocols/errors.go)
* [`MissingParameterError`](./blob/master/src/github.com/stellar/gateway/protocols/errors.go)
* [`TransactionBadSequence`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/errors.go)
* [`TransactionBadAuth`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/errors.go)
* [`TransactionInsufficientBalance`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/errors.go)
* [`TransactionNoAccount`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/errors.go)
* [`TransactionInsufficientFee`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/errors.go)
* [`TransactionBadAuthExtra`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/errors.go)
* [`AllowTrustMalformed`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/authorize.go)
* [`AllowTrustNoTrustline`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/authorize.go)
* [`AllowTrustTrustNotRequired`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/authorize.go)
* [`AllowTrustCantRevoke`](./blob/master/src/github.com/stellar/gateway/protocols/bridge/authorize.go)

## Hooks

The Bridge server listens for payment operations to the account specified by `accounts.receiving_account_id`. Every time 
a payment arrives it will send a HTTP POST request to `hooks.receive`.

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
`data` | Value of the [AuthData](https://www.stellar.org/developers/learn/integration-guides/compliance-protocol.html). This field will be empty when compliance server is not connected.

#### Response

Respond with `200 OK` when processing succeeded. Any other status code will be considered an error.

## Security

* This server must be set up in an isolated environment (ex. AWS VPC). Please make sure your firewall is properly configured 
and accepts connections from a trusted IPs only. You can also set the `api_key` config parameter but it's not recommended. 
If you don't set this properly, an unauthorized person will be able to submit transactions from your accounts!
* Make sure the `hooks` you provide only accept connections from the bridge server IP.
* Remember that `hooks.receive` may be called multiple times with the same payment. Check `id` parameter and ignore 
requests with the same value (just send `200 OK` response).

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
