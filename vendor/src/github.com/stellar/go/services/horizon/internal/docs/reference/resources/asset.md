---
title: Asset
---

**Assets** are the units that are traded on the Stellar Network.

An asset consists of an type, code, and issuer.

To learn more about the concept of assets in the Stellar network, take a look at the [Stellar assets concept guide](https://www.stellar.org/developers/guides/concepts/assets.html).

## Attributes

|    Attribute     |  Type  |                                                                                                                                |
| ---------------- | ------ | ------------------------------------------------------------------------------------------------------------------------------ |
| asset_type               | string | The type of this asset: "credit_alphanum4", or "credit_alphanum12". |
| asset_code               | string | The code of this asset.   |
| asset_issuer             | string | The issuer of this asset. |
| amount                   | number | The number of units of credit issued. |
| num_accounts             | number | The number of accounts that: 1) trust this asset and 2) where if the asset has the auth_required flag then the account is authorized to hold the asset. |
| flags                    | object | The flags on this asset of types: auth_required and auth_revocable. |
| paging_token             | string | A [paging token](./page.md) suitable for use as the `cursor` parameter to transaction collection resources.                   |

## Links

|  rel  |    Example                                        |    Description    |
| toml  | `https://www.stellar.org/.well-known/stellar.toml`| Link to the TOML file for this issuer |

## Example

```json
{
  "_links": {
    "toml": {
      "href": "https://www.stellar.org/.well-known/stellar.toml"
    }
  },
  "asset_type": "credit_alphanum4",
  "asset_code": "USD",
  "asset_issuer": "GBAUUA74H4XOQYRSOW2RZUA4QL5PB37U3JS5NE3RTB2ELJVMIF5RLMAG",
  "paging_token": "USD_GBAUUA74H4XOQYRSOW2RZUA4QL5PB37U3JS5NE3RTB2ELJVMIF5RLMAG_credit_alphanum4",
  "amount": "100.0000000",
  "num_accounts": 91547871,
  "flags": {
    "auth_required": false,
    "auth_revocable": false
  }
}
```

## Endpoints

|  Resource                      |    Type    |    Resource URI Template     |
| ------------------------------ | ---------- | ---------------------------- |
| [All Assets](../assets-all.md) | Collection | `/assets` (`GET`)            |