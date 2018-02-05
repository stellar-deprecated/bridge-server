---
title: Trade
---

A trade represents a fulfilled offer.  For example, let's say that there exists an offer to sell 9 `foo_bank/EUR` for 3 `baz_exchange/BTC` and you make an offer to buy 3 `foo_bank/EUR` for 1 `baz_exchange/BTC`.  Since your offer and the existing one cross, a trade happens.  After the trade completes:

- you are 3 `foo_bank/EUR` richer and 1 `baz_exchange/BTC` poorer
- the maker of the other offer is 1 `baz_exchange/BTC` richer and 3 `foo_bank/EUR` poorer
- your offer is completely fulfilled and no longer exists
- the other offer is partially fulfilled and becomes an offer to sell 6 `foo_bank/EUR` for 2 `baz_exchange/BTC`.  The price of that offer doesn't change, but the amount does.

Trades can also be caused by successful [path payments](https://www.stellar.org/developers/learn/concepts/exchange.html), because path payments involve fulfilling offers.

Payments are one-way in that afterwards, the source account has a smaller balance and the destination account of the payment has a bigger one.  Trades are two-way; both accounts increase and decrease their balances.

A trade occurs between two parties - `base` and `counter`. Which is which is either arbitrary or determined by the calling query.

## Attributes
| Attribute    | Type             |                                                                                                                        |
|--------------|------------------|------------------------------------------------------------------------------------------------------------------------|
| id | string | The ID of this trade. |
| paging_token | string | A [paging token](./page.md) suitable for use as a `cursor` parameter.|
| ledger_close_time | string | An ISO 8601 formatted string of when the ledger with this trade was closed.|
| base_account | string | base party of this trade|
| base_amount | string | amount of base asset that was moved from `base_account` to `counter_account`|
| base_asset_type | string | type of base asset|
| base_asset_code | string | code of base asset|
| base_asset_issuer | string | issuer of base asset|
| counter_account | string | counter party of this trade|
| counter_amount | string | amount of counter asset that was moved from `counter_account` to `base_account`|
| counter_asset_type | string | type of counter asset|
| counter_asset_code | string | code of counter asset|
| counter_asset_issuer | string | issuer of counter asset|
| base_is_seller | boolean | indicates which party of the trade made the sell offer|

## Links

| rel          | Example                                                                                           | Description                                                | `templated` |
|--------------|---------------------------------------------------------------------------------------------------|------------------------------------------------------------|-------------|
| base      | `/accounts/{base_account}`      | Link to details about the base account| true        |
| counter | `/accounts/{counter_account}`      | Link to details about the counter account | true        |
| operation | `/operation/{operation_id}` | Link to the operation of the assets bought and sold. | true |

## Endpoints

| Resource                 | Type       | Resource URI Template                |
|--------------------------|------------|--------------------------------------|
| [Trades](../endpoints/trades.md)       | Collection | `/trades`       |
