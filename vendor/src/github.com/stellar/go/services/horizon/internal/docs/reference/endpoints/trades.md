---
title: Trades
---

People on the Stellar network can make [offers](../resources/offer.md) to buy or sell assets. When an offer is fully or partially fulfilled, a [trade](../resources/trade.md) happens.

Trades can be filtered for specific orderbook, defined by an asset pair: `base` and `counter`. 

## Request

```
GET /trades?base_asset_type={base_asset_type}&base_asset_code={base_asset_code}&base_asset_issuer={base_asset_issuer}&counter_asset_type={counter_asset_type}&counter_asset_code={counter_asset_code}&counter_asset_issuer={counter_asset_issuer}&resolution={resolution}&start_time={start_time}&end_time={end_time}
```

### Arguments

| name | notes | description | example |
| ---- | ----- | ----------- | ------- |
| `base_asset_type` | optional, string | Type of base asset | `native` |
| `base_asset_code` | optional, string | Code of base asset, not required if type is `native` | `USD` |
| `base_asset_issuer` | optional, string | Issuer of base asset, not required if type is `native` | 'GA2HGBJIJKI6O4XEM7CZWY5PS6GKSXL6D34ERAJYQSPYA6X6AI7HYW36' |
| `counter_asset_type` | optional, string | Type of counter asset  | `credit_alphanum4` |
| `counter_asset_code` | optional, string | Code of counter asset, not required if type is `native` | `BTC` |
| `counter_asset_issuer` | optional, string | Issuer of counter asset, not required if type is `native` | 'GD6VWBXI6NY3AOOR55RLVQ4MNIDSXE5JSAVXUTF35FRRI72LYPI3WL6Z' |
| `offer_id` | optional, string | filter for by a specific offer id | `283606` |
| `?cursor` | optional, any, default _null_ | A paging token, specifying where to start returning records from. | `12884905984` |
| `?order`  | optional, string, default `asc` | The order, in terms of timeline, in which to return rows, "asc" or "desc". | `asc` |
| `?limit`  | optional, number, default: `10` | Maximum number of records to return. | `200` |

### curl Example Request
```sh 
curl "https://horizon.stellar.org/trades?counter_asset_type=credit_alphanum4&base_asset_type=native&counter_asset_issuer=GAP5LETOV6YIE62YAM56STDANPRDO7ZFDBGSNHJQIYGGKSMOZAHOOS2S&base_asset_issuer=undefined&counter_asset_code=EURT&base_asset_code=XLM&order=desc&limit=200"
```

## Response

The list of trades. `base` and `counter` in the records will match the asset pair filter order. If an asset pair is not specified, the order is arbitrary.

### Example Response
```json
{
  "_links": {
    "self": {
      "href": "https://horizon.stellar.org/trades?order=desc\u0026limit=2\u0026cursor="
    },
    "next": {
      "href": "https://horizon.stellar.org/trades?order=desc\u0026limit=2\u0026cursor=64255919088738305-0"
    },
    "prev": {
      "href": "https://horizon.stellar.org/trades?order=asc\u0026limit=2\u0026cursor=64283226490810369-0"
    }
  },
  "_embedded": {
    "records": [
      {
        "_links": {
          "base": {
            "href": "https://horizon.stellar.org/accounts/GDZYXBXG4PIQYLHY7BXDMMP3CM3QP2MC65W44M2TP2OLIR6XHGHG3OHG"
          },
          "counter": {
            "href": "https://horizon.stellar.org/accounts/GDAGT3NCVD4VCLN4TBRPHPJURX2KKCCZPA3WCROTJDUQI73XXJ4LCIMF"
          },
          "operation": {
            "href": "https://horizon.stellar.org/operations/64283226490810369"
          }
        },
        "id": "64283226490810369-0",
        "paging_token": "64283226490810369-0",
        "ledger_close_time": "2017-12-08T20:27:12Z",
        "offer_id": "286304",
        "base_account": "GDZYXBXG4PIQYLHY7BXDMMP3CM3QP2MC65W44M2TP2OLIR6XHGHG3OHG",
        "base_amount": "451.0000000",
        "base_asset_type": "native",
        "counter_account": "GDAGT3NCVD4VCLN4TBRPHPJURX2KKCCZPA3WCROTJDUQI73XXJ4LCIMF",
        "counter_amount": "0.0027962",
        "counter_asset_type": "credit_alphanum4",
        "counter_asset_code": "BTC",
        "counter_asset_issuer": "GATEMHCCKCY67ZUCKTROYN24ZYT5GK4EQZ65JJLDHKHRUZI3EUEKMTCH",
        "base_is_seller": false
      },
      {
        "_links": {
          "base": {
            "href": "https://horizon.stellar.org/accounts/GDAGT3NCVD4VCLN4TBRPHPJURX2KKCCZPA3WCROTJDUQI73XXJ4LCIMF"
          },
          "counter": {
            "href": "https://horizon.stellar.org/accounts/GADROEGWPXBXIEY4HF4U5R7JHK32HJW33DWDFLLSTID4KH23QVR6KMNC"
          },
          "operation": {
            "href": "https://horizon.stellar.org/operations/64255919088738305"
          }
        },
        "id": "64255919088738305-0",
        "paging_token": "64255919088738305-0",
        "ledger_close_time": "2017-12-08T11:26:20Z",
        "offer_id": "283606",
        "base_account": "GDAGT3NCVD4VCLN4TBRPHPJURX2KKCCZPA3WCROTJDUQI73XXJ4LCIMF",
        "base_amount": "233.6065573",
        "base_asset_type": "native",
        "counter_account": "GADROEGWPXBXIEY4HF4U5R7JHK32HJW33DWDFLLSTID4KH23QVR6KMNC",
        "counter_amount": "0.0028500",
        "counter_asset_type": "credit_alphanum4",
        "counter_asset_code": "BTC",
        "counter_asset_issuer": "GATEMHCCKCY67ZUCKTROYN24ZYT5GK4EQZ65JJLDHKHRUZI3EUEKMTCH",
        "base_is_seller": true
      }
    ]
  }
}
```

## Possible Errors

- The [standard errors](../errors.md#Standard_Errors).
