---
title: Trade Aggregations
---

Trade Aggregations are catered specifically for developers of trading clients. They facilitate efficient gathering of historical trade data. This is done by dividing a given time range into segments and aggregate statistics, for a given asset pair (`base`, `counter`) over each of these segments.


## Request

```
GET /trade_aggregations?base_asset_type={base_asset_type}&base_asset_code={base_asset_code}&base_asset_issuer={base_asset_issuer}&counter_asset_type={counter_asset_type}&counter_asset_code={counter_asset_code}&counter_asset_issuer={counter_asset_issuer}
```

### Arguments

| name | notes | description | example |
| ---- | ----- | ----------- | ------- |
| `start_time` | long | lower time boundary represented as millis since epoch| 1512689100000 |
| `end_time` | long | upper time boundary represented as millis since epoch| 1512775500000|
| `resolution` | long | segment duration as millis since epoch. *Supported values are 5 minutes (300000), 15 minutes (900000), 1 hour (3600000), 1 day (86400000) and 1 week (604800000).*| 300000|
| `base_asset_type` | string | Type of base asset | `native` |
| `base_asset_code` | string | Code of base asset, not required if type is `native` | `USD` |
| `base_asset_issuer` | string | Issuer of base asset, not required if type is `native` | 'GA2HGBJIJKI6O4XEM7CZWY5PS6GKSXL6D34ERAJYQSPYA6X6AI7HYW36' |
| `counter_asset_type` | string | Type of counter asset  | `credit_alphanum4` |
| `counter_asset_code` | string | Code of counter asset, not required if type is `native` | `BTC` |
| `counter_asset_issuer` | string | Issuer of counter asset, not required if type is `native` | 'GATEMHCCKCY67ZUCKTROYN24ZYT5GK4EQZ65JJLDHKHRUZI3EUEKMTCH' |
| `?order`  | optional, string, default `asc` | The order, in terms of timeline, in which to return rows, "asc" or "desc". | `asc` |
| `?limit`  | optional, number, default: `10` | Maximum number of records to return. | `200` |

### curl Example Request
```sh
curl "https://horizon.stellar.org/trade_aggregations?base_asset_type=native&start_time=1512689100000&counter_asset_issuer=GATEMHCCKCY67ZUCKTROYN24ZYT5GK4EQZ65JJLDHKHRUZI3EUEKMTCH&limit=200&end_time=1512775500000&counter_asset_type=credit_alphanum4&resolution=300000&order=asc&counter_asset_code=BTC"
```

## Response

A list of collected trade aggregations.

Note
- Segments that fit into the time range but have 0 trades in them, will not be included.
- Partial segments, in the beginning and end of the time range, will not be included.

### Example Response
```json
{
  "_links": {
    "self": {
      "href": "https://horizon.stellar.org/trade_aggregations?base_asset_type=native\u0026start_time=1512689100000\u0026counter_asset_issuer=GATEMHCCKCY67ZUCKTROYN24ZYT5GK4EQZ65JJLDHKHRUZI3EUEKMTCH\u0026limit=200\u0026end_time=1512775500000\u0026counter_asset_type=credit_alphanum4\u0026resolution=300000\u0026order=asc\u0026counter_asset_code=BTC"
    },
    "next": {
      "href": "https://horizon.stellar.org/trade_aggregations?base_asset_type=native\u0026counter_asset_code=BTC\u0026counter_asset_issuer=GATEMHCCKCY67ZUCKTROYN24ZYT5GK4EQZ65JJLDHKHRUZI3EUEKMTCH\u0026counter_asset_type=credit_alphanum4\u0026end_time=1512775500000\u0026limit=200\u0026order=asc\u0026resolution=300000\u0026start_time=1512765000000"
    }
  },
  "_embedded": {
    "records": [
      {
        "timestamp": 1512731100000,
        "trade_count": 2,
        "base_volume": "341.8032786",
        "counter_volume": "0.0041700",
        "avg": "0.0000122",
        "high": "0.0000122",
        "low": "0.0000122",
        "open": "0.0000122",
        "close": "0.0000122"
      },
      {
        "timestamp": 1512732300000,
        "trade_count": 1,
        "base_volume": "233.6065573",
        "counter_volume": "0.0028500",
        "avg": "0.0000122",
        "high": "0.0000122",
        "low": "0.0000122",
        "open": "0.0000122",
        "close": "0.0000122"
      },
      {
        "timestamp": 1512764700000,
        "trade_count": 1,
        "base_volume": "451.0000000",
        "counter_volume": "0.0027962",
        "avg": "0.0000062",
        "high": "0.0000062",
        "low": "0.0000062",
        "open": "0.0000062",
        "close": "0.0000062"
      }
    ]
  }
}
```

## Possible Errors

- The [standard errors](../errors.md#Standard_Errors).
