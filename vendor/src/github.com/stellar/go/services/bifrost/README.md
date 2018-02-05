# bifrost

> In Norse mythology, Bifröst (/ˈbɪvrɒst/ or sometimes Bilröst or Bivrost) is a burning rainbow bridge that reaches between Midgard (Earth) and Asgard, the realm of the gods. [**Wikipedia**](https://en.wikipedia.org/wiki/Bifr%C3%B6st)

Bifrost is highly available and secure Bitcoin/Ethereum → Stellar bridge. It allows users to move BTC/ETH to Stellar network and then trade them to other tokens or participate in ICOs (initial coin offering).

It solves many problems connected to moving tokens to Stellar network:

* Security:
  * Developers don’t have access to users’ Stellar keys.
  * No Bitcoin/Ethereum private keys for receiving accounts uploaded to any application machine.
* High availability:
  * Can be deployed to multiple availability zones (one or more machines can fail).
  * Handles high load.
* Easy to use:
  * Simple process for users.
  * Easy installation for developers.

We are releasing the **alpha version** of this software. We encourage our community of developers to test and improve it.

Download the binary from [the release page](https://github.com/stellar/go/releases/tag/bifrost-v0.0.1) and use it with it's [Bifrost JS SDK](https://github.com/stellar/bifrost-js-sdk).

## How it works

1. User opens the web app implemented using [Bifrost JS SDK](https://github.com/stellar/bifrost-js-sdk).
1. User is presented with her public and private Stellar keys where Bitcoin/Ethereum will be sent.
1. User selects what cryptocurrency she wants to move to Stellar network.
1. A receiving Bitcoin/Etherem address is generated.
1. User sends funds in Bitcoin/Ethereum network.
1. Bifrost listens to Bitcoin and Ethereum network events. When payment arrives it creates a Stellar [account](https://www.stellar.org/developers/guides/concepts/accounts.html) for the user.
1. User creates a [trust line](https://www.stellar.org/developers/guides/concepts/assets.html) to BTC/ETH issued by Bifrost account.
1. Bifrost sends corresponding amount of BTC/ETH to user's Stellar account.
1. As an optional step, web application can also exchange BTC/ETH to other token in Stellar network at a given rate.

## Demo

We use Ethereum Ropsten test network for this demo because it's faster that Bitcoin testnet.

https://bifrost.stellar.org/

1. First you need some ETH on Ethereum testnet.
1. Create an account at https://www.myetherwallet.com/, then switch Network (top-right dropdown) to "Ropsten (MyEtherWallet)". Write down/copy your Ethereum address somewhere.
1. Use http://faucet.ropsten.be:3001/ to send 3 ETH to your testnet account.
1. Now open Bifrost demo: https://bifrost.stellar.org/ It will display address where to send your test ETH.
1. Go back to MyEtherWallet and send ETH to the address displayed by Bifrost. Sometimes Ropsten network is overloaded so monitor the Etherscan to check if your tx was included in a block. If not, increase the gas price (this can be done in "Send Offline" tab).
1. Switch back to Bifrost demo and check the progress.

## Config

* `port` - bifrost server listening port
* `using_proxy` (default `false`) - set to `true` if bifrost lives behind a proxy or load balancer
* `bitcoin`
  * `master_public_key` - master public key for bitcoin keys derivation (read more in [BIP-0032](https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki))
  * `rpc_server` - URL of [bitcoin-core](https://github.com/bitcoin/bitcoin) >= 0.15.0 RPC server
  * `rpc_user` (default empty) - username for RPC server (if any)
  * `rpc_pass` (default empty) - password for RPC server (if any)
  * `testnet` (default `false`) - set to `true` if you're testing bifrost in ethereum
  * `minimum_value_btc` - minimum transaction value in BTC that will be accepted by Bifrost, everything below will be ignored.
* `ethereum`
  * `master_public_key` - master public key for bitcoin keys derivation (read more in [BIP-0032](https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki))
  * `rpc_server` - URL of [geth](https://github.com/ethereum/go-ethereum) >= 1.7.1 RPC server
  * `network_id` - network ID (`3` - Ropsten testnet, `1` - live Ethereum network)
  * `minimum_value_eth` - minimum transaction value in ETH that will be accepted by Bifrost, everything below will be ignored.
* `stellar`
  * `token_asset_code` - asset code for the token that will be distributed
  * `issuer_public_key` - public key of the assets issuer or hot wallet,
  * `signer_secret_key` - issuer's secret key if only one instance of Bifrost is deployed OR [channel](https://www.stellar.org/developers/guides/channels.html)'s secret key if more than one instance of Bifrost is deployed. Signer's sequence number will be consumed in transaction's sequence number.
  * `horizon` - URL to [horizon](https://github.com/stellar/go/tree/master/services/horizon) server
  * `network_passphrase` - Stellar network passphrase (`Public Global Stellar Network ; September 2015` for production network, `Test SDF Network ; September 2015` for test network)
* `database`
  * `type` - currently the only supported database type is: `postgres`
  * `dsn` - data source name for postgres connection (`postgres://user:password@host/dbname?sslmode=sslmode` - [more info](https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters))

## Deployment

There are two ways you can deploy Bifrost:
* **Simple/single-instance deployment** - Bifrost will live on a single instance along with geth, bitcoin-core and database. This is the easiest way to deploy Bifrost but if this instance goes down, so does Bifrost and all the other components.
* **High-availability deployment** - Each of the components is deployed to at least 2 instances so even if random instances go down, Bifrost will still be online accepting payments.

We recommend the latter.

### High-availability deployment

Here's the proposed architecture diagram of high-availability deployment:

![Architecture](./images/architecture.png)

## Going to production

* Remember that everyone with master public key and **any** child private key can recover your **master** private key. Do not share your master public key and obviously any private keys. Treat your master public key as if it was a private key. Read more in BIP-0032 [Security](https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki#security) section.
* Make sure "Sell [your token] for BTC" and/or "Sell [your token] for ETH" exist in Stellar production network. If not, you can create an offer by sending a transaction with `manage_offer` operation.
* Make sure you don't use account from `stellar.issuer_secret_key` anywhere else than bifrost. Otherwise, sequence numbers will go out of sync and bifrost will stop working. It's good idea to create a new signer on issuing account.
* Check if public master key correct. Use CLI tool (`bifrost check-keys`) to generate a few addresses and ensure you have corresponding private keys! You should probably send test transactions to some of these addresses and check if you can withdraw funds.
* Make sure `using_proxy` variable is set to correct value. Otherwise you will see your proxy IP instead of users' IPs in logs.
* Make sure you're not connecting to testnets.
* Deploy at least 2 bifrost, bitcoin-core, geth, stellar-core and horizon servers. Use multi-AZ database.
* Do not use SDF's horizon servers. There is no SLA and we cannot guarantee it will handle your load.
* Make sure bifrost <-> bitcoin-core and bifrost <-> geth connections are not public or are encrypted (mitm attacks).
* It's a good idea to set "Authorization required" and "Authorization revocable" [flags](https://www.stellar.org/developers/guides/concepts/accounts.html#flags) on issuing account during ICO stage to remove trustlines to accounts with lost keys. You need to set `stellar.needs_authorize` in config to `true` if you want to use this feature.
* Monitor bifrost logs and react to all WARN and ERROR entries.
* Make sure you are using geth >= 1.7.1 and bitcoin-core >= 0.15.0.
* Increase horizon rate limiting to handle expected load.
* Make sure you configured minimum accepted value for Bitcoin and Ethereum transactions to the value you really want.
* Make sure you start from a fresh Bifrost DB in production. If Bifrost was running, you stopped bitcoin-core or geth and then started it again then all the Bitcoin and Ethereum blocks mined during that period will be processed which can take a lot of time.
