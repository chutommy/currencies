# Currencies

The Currencies is a microservice, which uses <a href="https://grpc.io/" target="_blank">gRPC technology</a>.
It supports both unary and bidirectional streaming calls, which allows data update every 6 seconds.
The service provides the latest market exchange rates of all currencies from all country. When an error occurs, it handles it in a non-fatal way with an error message.

The whole service is containerized using a Docker engine and the service can be easily launched and deployed with the pre-prepared `make` commands in the Makefile.

The Currencies obtains all necessary data from the <a href="https://markets.businessinsider.com/currencies" target="_blank">Business Insider</a> website. The algorithm does not infringe any copyrights nor the website robots exclusion protocol.

## Installation

### Requirements
- <a href="https://git-scm.com/downloads" target="_blank">Git</a>
- <a href="https://docs.docker.com/get-docker/" target="_blank">Docker Engine</a>

### Linux/Mac
The installation process on the Windows machines would be slightly different.
```bash
$ git clone https://github.com/chutommy/currencies.git     # download repository
$ cd currencies               # move to repository dir
$ make build                  # build docker image
$ make run                    # initialize service
```

## Supported currency codes
<table>
    <tr> <td>BTG</td> <td>KZT</td> <td>GMD</td> <td>NMR</td> <td>PEN</td> <td>ARDR</td> <td>DCR</td> <td>FUN</td> <td>REQ</td> <td>IRR</td> </tr>
    <tr> <td>QSP</td> <td>ANT</td> <td>CVE</td> <td>UAH</td> <td>MCO</td> <td>PYG</td> <td>BTM</td> <td>DOP</td> <td>NGN</td> <td>QTUM</td> </tr>
    <tr> <td>SDG</td> <td>XRP</td> <td>CVC</td> <td>ICX</td> <td>LBC</td> <td>VTC</td> <td>XVG</td> <td>COP</td> <td>CZK</td> <td>SCR</td> </tr>
    <tr> <td>GNO</td> <td>KMF</td> <td>NXT</td> <td>IDR</td> <td>PAB</td> <td>BSV</td> <td>BTS</td> <td>OMG</td> <td>PPT</td> <td>BIF</td> </tr>
    <tr> <td>MKD</td> <td>XEM</td> <td>HRK</td> <td>NET</td> <td>UYU</td> <td>INR</td> <td>MAD</td> <td>RLC</td> <td>HKD</td> <td>LAK</td> </tr>
    <tr> <td>BCC</td> <td>BZD</td> <td>EUR</td> <td>GBP</td> <td>MYR</td> <td>BTC</td> <td>GIP</td> <td>LRC</td> <td>NEO</td> <td>NULS</td> </tr>
    <tr> <td>STX</td> <td>AFN</td> <td>DATA</td> <td>ETH</td> <td>CND</td> <td>GRID</td> <td>LKR</td> <td>TNT</td> <td>ADA</td> <td>DRGN</td> </tr>
    <tr> <td>LINK</td> <td>PKR</td> <td>PLN</td> <td>RWF</td> <td>SYS</td> <td>CNY</td> <td>SNT</td> <td>WTC</td> <td>XMR</td> <td>ETC</td> </tr>
    <tr> <td>KHR</td> <td>LRD</td> <td>LTC</td> <td>LYD</td> <td>SOS</td> <td>THB</td> <td>MMK</td> <td>USD</td> <td>NPR</td> <td>AED</td> </tr>
    <tr> <td>ENJ</td> <td>GTQ</td> <td>AION</td> <td>CNX</td> <td>XAF</td> <td>BGN</td> <td>DTR</td> <td>MONA</td> <td>RCN</td> <td>BAT</td> </tr>
    <tr> <td>BOB</td> <td>TOP</td> <td>UTK</td> <td>BYN</td> <td>GNT</td> <td>TJS</td> <td>WAVES</td> <td>KCS</td> <td>LEND</td> <td>MGA</td> </tr>
    <tr> <td>REP</td> <td>ZRX</td> <td>ARK</td> <td>ISK</td> <td>MANA</td> <td>STORJ</td> <td>GBYTE</td> <td>KMD</td> <td>NIO</td> <td>HTG</td> </tr>
    <tr> <td>SZL</td> <td>AUD</td> <td>EOS</td> <td>LA</td> <td>TWD</td> <td>DZD</td> <td>SVC</td> <td>JPY</td> <td>KNC</td> <td>MAID</td> </tr>
    <tr> <td>PIVX</td> <td>SGD</td> <td>USDT</td> <td>CNH</td> <td>ELF</td> <td>GRS</td> <td>MXN</td> <td>ALL</td> <td>FCT</td> <td>KES</td> </tr>
    <tr> <td>SBD</td> <td>ZEN</td> <td>EDO</td> <td>BNB</td> <td>NZD</td> <td>RDD</td> <td>XZC</td> <td>ENG</td> <td>ILS</td> <td>ITC</td> </tr>
    <tr> <td>TMT</td> <td>BCN</td> <td>MIOTA</td> <td>RUB</td> <td>TZS</td> <td>BRL</td> <td>BSD</td> <td>DNA</td> <td>ZMW</td> <td>BAY</td> </tr>
    <tr> <td>NAS</td> <td>SLS</td> <td>XLM</td> <td>MVR</td> <td>MWK</td> <td>NOK</td> <td>SEK</td> <td>STRAT</td> <td>MTL</td> <td>CLP</td> </tr>
    <tr> <td>DASH</td> <td>HNL</td> <td>TAAS</td> <td>TRX</td> <td>JMD</td> <td>KRW</td> <td>MOP</td> <td>SAR</td> <td>VND</td> <td>ZEC</td> </tr>
    <tr> <td>BNT</td> <td>DKK</td> <td>ETB</td> <td>OMR</td> <td>TTD</td> <td>XWC</td> <td>EGP</td> <td>KIN</td> <td>MLN</td> <td>NAD</td> </tr>
    <tr> <td>PHP</td> <td>TNB</td> <td>ZAR</td> <td>CAD</td> <td>CHF</td> <td>MUR</td> <td>YER</td> <td>ARS</td> <td>BDT</td> <td>DENT</td> </tr>
    <tr> <td>BMD</td> <td>CRC</td> <td>DGB</td> <td>GNF</td> <td>STEEM</td> <td>ETN</td> <td>MOON</td> <td>SC</td> <td>TRY</td> <td>AE</td> </tr>
    <tr> <td>HUF</td> <td>JOD</td> <td>LBP</td> <td>LSK</td> <td>SRD</td> <td>BND</td> <td>CUP</td> <td>GAS</td> <td>UGX</td> <td>AMD</td> </tr>
    <tr> <td>LSL</td> <td>BHD</td> <td>DJF</td> <td>DOGE</td> <td>GYD</td> <td>KWD</td> <td>POWR</td> <td>QAR</td>  </tr>
</table>

**Note:**
*The Currency request holds the key "Name" and its value is **not** case sensitive.*

## Usage
### Currency.GetRate
GetCurrency provides current data of one certain currency. The data holds the currency code, country of origin, the description, the last currency value change in percentages, the exchange rate to USD and the time of the last update.

__GetRateRequest__ defines the request message for the GetRate call. It needs the base currency and the destination currency. Supported currencies are <a href="https://github.com/chutommy/currencies#supported-currency-codes">here</a>.

*Base represents the base currency for the exchange rate.*<br>
*Destination represents the destination currency for the exchange rate.*
```proto
message GetRateRequest {
    string Base = 1;
    string Destination = 2;
}
```
```json
{
    "Base":"EUR",
    "Destination":"USD"
}
```

__GetRateResponse__ defines the response message for the GetRate call. It holds only the exchange rate of the request's base and destination.

*Rate is the result exchange rate.*
```proto
message GetRateResponse {
    float Rate = 1;
}
```
```json
    "Rate":1.1655
```

### Currency.GetCurrency
GetRate calculates the exchange rate between the base and the destination. The service takes the latest data from the source.

__GetCurrencyRequest__ defines the request message for the GetCurrency and the SubscribeCurrency calls.

*Name stands for the currency code for the currency. The Name value is not case sensitive.*
```proto
message GetCurrencyRequest {
    string Name = 1;
```
```json
{
    "Name":"CAD"
}
```

__GetCurrencyResponse__ defines the response message for the GetCurrency call and the StreamingSubscribeResponse message.

*Name stands for the currency code for the currency. Every Name value is capitalized.*<br>
*Country holds the name of the country where the currency came from.*<br>
*Description is the full name of the currency.*<br>
*Change represents the latest currency change in the percentages.*<br>
*RateUSD is the exchange rates between the currency and the USD. Both currency values are taken from the lastest source update.*<br>
*UpdatedAt is the time of the last update of the currency in the source.*
```proto
message GetCurrencyResponse {
   string Name = 1;
   string Country = 2;
   string Description = 3;
   float Change = 4;
   float RateUSD = 5;
   string UpdatedAt = 6;
}
```
```json
{
    "Name": "CAD",
    "Country": "Canada",
    "Description": "Canadian Dollar",
    "Change": -0.02,
    "RateUSD": 1.3416,
    "UpdatedAt": "2020-07-25 04:04:00 +0000 UTC"
}
```

### Currency.SubscribeCurrency
SubscribeCurrency works as the GetCurrency call, except that it does not send a response instantly but wait until the database changes some of its value, then it sends all subscribed currency data to each client.

__GetCurrencyRequest__ defines the request message for the GetCurrency and the SubscribeCurrency calls.
```json
{"Name":"GBP"}
{"Name":"VND"}
```

__StreamingSubscribeResponse__ defines the response message for the SubscribeCurrency call. It holds either GetCurrencyResponse or the Status error.

*Get_currency_response defines the response message with the data about the currency.*<br>
*Error defines the error status of the problem which occurred.*
```proto
message StreamingSubscribeResponse {
    oneof message{
        GetCurrencyResponse GetCurrencyResponse = 1;
        google.rpc.Status Error = 2;
    }
}
```
```
{
    "GetCurrencyResponse": {
        "Name": "GBP",
        "Country": "England",
        "Description": "British Pound",
        "Change": -0.03,
        "RateUSD": 0.7815,
        "UpdatedAt": "2020-07-25 04:04:00 +0000 UTC"
    }
}
{
    "GetCurrencyResponse": {
        "Name": "VND",
        "Country": "Vietnam",
        "Description": "Vietnamese Dong",
        "Change": 0.02,
        "RateUSD": 23185,
        "UpdatedAt": "2020-07-25 04:04:00 +0000 UTC"
    }
}
```
Server logs:
```bash
[CURRENCY SERVICE]2020/07/25 10:25:14 [start] listening on 127.0.0.1:10502
[CURRENCY SERVICE]2020/07/25 10:25:23 [success] client successfully subscribed to: GBP
[CURRENCY SERVICE]2020/07/25 10:25:31 [success] client successfully subscribed to: VND
[CURRENCY SERVICE]2020/07/25 10:25:32 [update] currency data updated
```

## Examples
For these examples, I am using the tool called <a href="https://github.com/fullstorydev/grpcurl" target="_blank">gRPCurl</a> to generate binary calls to gRPC servers.

### GetRate
#### Currency.GetRate: `{"Base":"RUB", "Destination":"USD"}`
```bash
[chutommy@localhost currencies]$ grpcurl --plaintext -d '{"Base":"RUB", "Destination":"USD"}' 127.0.0.1:10502 Currency.GetRate
{
    "Rate": 0.0139
}
```

#### Currency.GetRate: `{"Base":"GBP", "Destination":"EUR"}`
```bash
[chutommy@localhost currencies]$ grpcurl --plaintext -d '{"Base":"GBP", "Destination":"EUR"}' 127.0.0.1:10502 Currency.GetRate
{
    "Rate": 1.0979
}
```

#### Currency.GetRate: `{"Base":"CZK", "Destination":"CAD"}`
```bash
[chutommy@localhost currencies]$ grpcurl --plaintext -d '{"Base":"CZK", "Destination":"CAD"}' 127.0.0.1:10502 Currency.GetRate
{
    "Rate": 0.0596
}
```

#### Server logs
```bash
[CURRENCY SERVICE]2020/07/25 10:38:01 [start] listening on 127.0.0.1:10502
[CURRENCY SERVICE]2020/07/25 10:38:32 [handle] GetRate call, base: RUB, destination: USD
[CURRENCY SERVICE]2020/07/25 10:38:51 [handle] GetRate call, base: GBP, destination: EUR
[CURRENCY SERVICE]2020/07/25 10:39:11 [handle] GetRate call, base: CZK, destination: CAD
```

### GetCurrency
#### Currency.GetCurrency: `{"Name":"USD"}`
```bash
[chutommy@localhost currencies]$ grpcurl --plaintext -d '{"Name":"USD"}' 127.0.0.1:10502 Currency.GetCurrency
{
    "Name": "USD",
    "Country": "United States of America",
    "Description": "United States Dollar",
    "RateUSD": 1,
    "UpdatedAt": "2020-07-25 10:42:26.385028702 +0200 CEST m=+0.180421634"
}
```

#### Currency.GetCurrency: `{"Name":"GBP"}`
```bash
[chutommy@localhost currencies]$ grpcurl --plaintext -d '{"Name":"GBP"}' 127.0.0.1:10502 Currency.GetCurrency
{
    "Name": "GBP",
    "Country": "England",
    "Description": "British Pound",
    "Change": -0.03,
    "RateUSD": 0.7815,
    "UpdatedAt": "2020-07-25 04:04:00 +0000 UTC"
}
```
#### Currency.GetCurrency: `{"Name":"CAD"}`
```bash
[chutommy@localhost currencies]$ grpcurl --plaintext -d '{"Name":"CAD"}' 127.0.0.1:10502 Currency.GetCurrency
{
    "Name": "CAD",
    "Country": "Canada",
    "Description": "Canadian Dollar",
    "Change": -0.02,
    "RateUSD": 1.3416,
    "UpdatedAt": "2020-07-25 04:04:00 +0000 UTC"
}
```
#### Server logs
```bash
[CURRENCY SERVICE]2020/07/25 10:42:10 [start] listening on 127.0.0.1:10502
[CURRENCY SERVICE]2020/07/25 10:42:27 [handle] GetCurrency call: USD
[CURRENCY SERVICE]2020/07/25 10:42:32 [handle] GetCurrency call: GBP
[CURRENCY SERVICE]2020/07/25 10:42:38 [handle] GetCurrency call: CAD
```

### SubscribeCurrency
#### Currency.SubscribeCurrency: `{"Name":"CAD"}{"Name":"CZK"}{"Name":"GBP"}`
```bash
[chutommy@localhost currencies]$ grpcurl --plaintext -d @ 127.0.0.1:10502 Currency.SubscribeCurrency
{"Name":"CAD"}
{"Name":"CZK"}
{"Name":"GBP"}
```

#### UPDATE
```bash
{
    "GetCurrencyResponse": {
        "Name": "CAD",
        "Country": "Canada",
        "Description": "Canadian Dollar",
        "Change": -0.02,
        "RateUSD": 1.3416,
        "UpdatedAt": "2020-07-25 04:04:00 +0000 UTC"
    }
}
{
    "GetCurrencyResponse": {
        "Name": "CZK",
        "Country": "Czech Republic",
        "Description": "Czech Koruna",
        "RateUSD": 22.526,
        "UpdatedAt": "2020-07-25 04:04:00 +0000 UTC"
    }
}
{
    "GetCurrencyResponse": {
        "Name": "GBP",
        "Country": "England",
        "Description": "British Pound",
        "Change": -0.03,
        "RateUSD": 0.7815,
        "UpdatedAt": "2020-07-25 04:04:00 +0000 UTC"
    }
}
```

#### Server logs
```bash
[CURRENCY SERVICE]2020/07/25 10:46:01 [start] listening on 127.0.0.1:10502
[CURRENCY SERVICE]2020/07/25 10:46:03 [success] client successfully subscribed to: CAD
[CURRENCY SERVICE]2020/07/25 10:46:10 [success] client successfully subscribed to: CZK
[CURRENCY SERVICE]2020/07/25 10:46:15 [success] client successfully subscribed to: GBP
[CURRENCY SERVICE]2020/07/25 10:46:18 [update] currency data updated
```

### Error handling
#### Currency.GetRate: `{"Base":"USD", "Destination":"invalid"}`
```bash
[chutommy@localhost currencies]$ grpcurl --plaintext -d '{"Base":"USD", "Destination":"invalid"}' 127.0.0.1:10502 Currency.GetRate
ERROR:
    Code: NotFound
    Message: Currency was not found: call GetRate: destination currency 'INVALID' not found.
```

#### Currency.GetCurrency: `{"Name":"invalid"}`
```bash
[chutommy@localhost currencies]$ grpcurl --plaintext -d '{"Name":"invalid"}' 127.0.0.1:10502 Currency.GetCurrency
ERROR:
    Code: NotFound
    Message: Currency "invalid" was not found.
```

#### Currency.SubscribeCurrency: `{"Name":"invalid"}`
```bash
[chutommy@localhost currencies]$ grpcurl --plaintext -d @ 127.0.0.1:10502 Currency.SubscribeCurrency
{"Name":"invalid"}
{
    "Error": {
        "code": 5,
        "message": "Currency \"INVALID\" was not found."
    }
}
{"Name":"USD"}
{
    "GetCurrencyResponse": {
        "Name": "USD",
        "Country": "United States of America",
        "Description": "United States Dollar",
        "RateUSD": 1,
        "UpdatedAt": "2020-07-25 11:02:50.145063563 +0200 CEST m=+45.314325138"
    }
}
```

#### Server logs
```bash
[CURRENCY SERVICE]2020/07/25 11:02:08 [start] listening on 127.0.0.1:10502
[CURRENCY SERVICE]2020/07/25 11:02:10 [error] handle error: call GetRate: destination currency 'INVALID' not found
[CURRENCY SERVICE]2020/07/25 11:02:18 [error] handle error: handling GetCurrency call: currency 'INVALID' not found
[CURRENCY SERVICE]2020/07/25 11:02:35 [error] currency "INVALID" not found
[CURRENCY SERVICE]2020/07/25 11:02:47 [success] client successfully subscribed to: USD
[CURRENCY SERVICE]2020/07/25 11:02:50 [update] currency data updated
```

## Client
All clients can be built with the <a href="https://grpc.io/docs/protoc-installation/" target="_blank">Protocol Buffer Compiler</a> + <a href="https://grpc.io/" target="_blank">gRPC</a> plugin.

*The protobuffer of the services:* <a href="https://github.com/chutommy/currencies/blob/master/protos/currency.proto">commodity.proto</a>

## Directory structure
```bash
_
├── config
│   ├── tests
│   │   ├── config_0.yaml
│   │   └── config_1.yaml
│   ├── config.go
│   └── config_test.go
├── data
│   ├── currencies.go
│   ├── currencies_test.go
│   ├── fetching.go
│   └── fetching_test.go
├── models
│   └── currency.go
├── protos
│   ├── currency
│   │   └── currency.pb.go
│   ├── google
│   │   └── rpc
│   │       └── status.proto
│   └── currency.proto
├── server
│   ├── currencies.go
│   ├── currencies_test.go
│   ├── handler.go
│   └── handler_test.go
├── config.yaml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── Makefile
└── README.md
```
