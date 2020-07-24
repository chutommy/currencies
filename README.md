# Currencies

The Currencies is a microservice, which using <a href="https://markets.businessinsider.com/currencies" target="_blank">gRPC technology</a>. It supports both unary and bidirectional calls, which allows data updates every 6 seconds.
It also provides the latest exchange rates of all currencies from every country. When an error occurs, it can handle it in a non-faal way with the error messages.

The whole service is containerized using a Docker engine and everything can be easily run and deployed with the pre-prepared `make` commands in the Makefile.

The Currencies obtains all necessary data from the <a href="https://markets.businessinsider.com/currencies" target="_blank">Business Insider</a> website. The algorithm does not infringe any copyrights nor the website robots exclusion protocol.

## Instalation

### Requirements
- <a href="https://git-scm.com/downloads" target="_blank">Git</a>
- <a href="https://docs.docker.com/get-docker/" target="_blank">Docker Engine</a>

### Linux/Mac
This is the exact way to download and run the service. On a Windows machine, the installation process would be slightly different.
```bash
$ git clone https://github.com/chutified/currencies.git     # download repository
$ cd currencies         # move to repository dir
$ make build                  # build docker image
$ make run                    # initialize service
```

## Supported currency codes
<table>
    <tr>
        <td>BTG</td>
        <td>KZT</td>
        <td>GMD</td>
        <td>NMR</td>
        <td>PEN</td>
        <td>ARDR</td>
        <td>DCR</td>
        <td>FUN</td>
        <td>REQ</td>
        <td>IRR</td>
        <td>QSP</td>
        <td>ANT</td>
    </tr>
    <tr>
        <td>CVE</td>
        <td>UAH</td>
        <td>MCO</td>
        <td>PYG</td>
        <td>BTM</td>
        <td>DOP</td>
        <td>NGN</td>
        <td>QTUM</td>
        <td>SDG</td>
        <td>XRP</td>
        <td>CVC</td>
        <td>ICX</td>
    </tr>
    <tr>
        <td>LBC</td>
        <td>VTC</td>
        <td>XVG</td>
        <td>COP</td>
        <td>CZK</td>
        <td>SCR</td>
        <td>GNO</td>
        <td>KMF</td>
        <td>NXT</td>
        <td>IDR</td>
        <td>PAB</td>
        <td>BSV</td>
    </tr>
    <tr>
        <td>BTS</td>
        <td>OMG</td>
        <td>PPT</td>
        <td>BIF</td>
        <td>MKD</td>
        <td>XEM</td>
        <td>HRK</td>
        <td>NET</td>
        <td>UYU</td>
        <td>INR</td>
        <td>MAD</td>
        <td>RLC</td>
    </tr>
    <tr>
        <td>HKD</td>
        <td>LAK</td>
        <td>BCC</td>
        <td>BZD</td>
        <td>EUR</td>
        <td>GBP</td>
        <td>MYR</td>
        <td>BTC</td>
        <td>GIP</td>
        <td>LRC</td>
        <td>NEO</td>
        <td>NULS</td>
    </tr>
    <tr>
        <td>STX</td>
        <td>AFN</td>
        <td>DATA</td>
        <td>ETH</td>
        <td>CND</td>
        <td>GRID</td>
        <td>LKR</td>
        <td>TNT</td>
        <td>ADA</td>
        <td>DRGN</td>
        <td>LINK</td>
        <td>PKR</td>
    </tr>
    <tr>
        <td>PLN</td>
        <td>RWF</td>
        <td>SYS</td>
        <td>CNY</td>
        <td>SNT</td>
        <td>WTC</td>
        <td>XMR</td>
        <td>ETC</td>
        <td>KHR</td>
        <td>LRD</td>
        <td>LTC</td>
        <td>LYD</td>
    </tr>
    <tr>
        <td>SOS</td>
        <td>THB</td>
        <td>MMK</td>
        <td>USD</td>
        <td>NPR</td>
        <td>AED</td>
        <td>ENJ</td>
        <td>GTQ</td>
        <td>AION</td>
        <td>CNX</td>
        <td>XAF</td>
        <td>BGN</td>
    </tr>
    <tr>
        <td>DTR</td>
        <td>MONA</td>
        <td>RCN</td>
        <td>BAT</td>
        <td>BOB</td>
        <td>TOP</td>
        <td>UTK</td>
        <td>BYN</td>
        <td>GNT</td>
        <td>TJS</td>
        <td>WAVES</td>
        <td>KCS</td>
    </tr>
    <tr>
        <td>LEND</td>
        <td>MGA</td>
        <td>REP</td>
        <td>ZRX</td>
        <td>ARK</td>
        <td>ISK</td>
        <td>MANA</td>
        <td>STORJ</td>
        <td>GBYTE</td>
        <td>KMD</td>
        <td>NIO</td>
        <td>HTG</td>
    </tr>
    <tr>
        <td>SZL</td>
        <td>AUD</td>
        <td>EOS</td>
        <td>LA</td>
        <td>TWD</td>
        <td>DZD</td>
        <td>SVC</td>
        <td>JPY</td>
        <td>KNC</td>
        <td>MAID</td>
        <td>PIVX</td>
        <td>SGD</td>
    </tr>
    <tr>
        <td>USDT</td>
        <td>CNH</td>
        <td>ELF</td>
        <td>GRS</td>
        <td>MXN</td>
        <td>ALL</td>
        <td>FCT</td>
        <td>KES</td>
        <td>SBD</td>
        <td>ZEN</td>
        <td>EDO</td>
        <td>BNB</td>
    </tr>
    <tr>
        <td>NZD</td>
        <td>RDD</td>
        <td>XZC</td>
        <td>ENG</td>
        <td>ILS</td>
        <td>ITC</td>
        <td>TMT</td>
        <td>BCN</td>
        <td>MIOTA</td>
        <td>RUB</td>
        <td>TZS</td>
        <td>BRL</td>
    </tr>
    <tr>
        <td>BSD</td>
        <td>DNA</td>
        <td>ZMW</td>
        <td>BAY</td>
        <td>NAS</td>
        <td>SLS</td>
        <td>XLM</td>
        <td>MVR</td>
        <td>MWK</td>
        <td>NOK</td>
        <td>SEK</td>
        <td>STRAT</td>
    </tr>
    <tr>
        <td>MTL</td>
        <td>CLP</td>
        <td>DASH</td>
        <td>HNL</td>
        <td>TAAS</td>
        <td>TRX</td>
        <td>JMD</td>
        <td>KRW</td>
        <td>MOP</td>
        <td>SAR</td>
        <td>VND</td>
        <td>ZEC</td>
    </tr>
    <tr>
        <td>BNT</td>
        <td>DKK</td>
        <td>ETB</td>
        <td>OMR</td>
        <td>TTD</td>
        <td>XWC</td>
        <td>EGP</td>
        <td>KIN</td>
        <td>MLN</td>
        <td>NAD</td>
        <td>PHP</td>
        <td>TNB</td>
    </tr>
    <tr>
        <td>ZAR</td>
        <td>CAD</td>
        <td>CHF</td>
        <td>MUR</td>
        <td>YER</td>
        <td>ARS</td>
        <td>BDT</td>
        <td>DENT</td>
        <td>BMD</td>
        <td>CRC</td>
        <td>DGB</td>
        <td>GNF</td>
    </tr>
    <tr>
        <td>STEEM</td>
        <td>ETN</td>
        <td>MOON</td>
        <td>SC</td>
        <td>TRY</td>
        <td>AE</td>
        <td>HUF</td>
        <td>JOD</td>
        <td>LBP</td>
        <td>LSK</td>
        <td>SRD</td>
        <td>BND</td>
    </tr>
    <tr>
        <td>CUP</td>
        <td>GAS</td>
        <td>UGX</td>
        <td>AMD</td>
        <td>LSL</td>
        <td>BHD</td>
        <td>DJF</td>
        <td>DOGE</td>
        <td>GYD</td>
        <td>KWD</td>
        <td>POWR</td>
        <td>QAR</td>
    </tr>
</table>

**Note:**
*The Currency request holds the key "Name" and its value is **not** case sensitive.*
*Currency names must not be completely lowercase nor uppercase to be found.*


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
