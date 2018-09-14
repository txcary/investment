# investment
A web site help to invest in China.

## Build Server
```bash
go build -o bin/server github.com/txcary/investment/cli/server
```

## Make config file
```bash
cd $GOPATH
vim investment.ini
```

Below is an example of the investment.ini
```ini
[lixinger]
token=YOUR TOKEN from https://www.lixinger.com/open/api/toke

[stock]
db=stock.db

[restful]
root=/workspace/go/src/github.com/txcary/investment/web
```
- [lixinger] token is to let server fetches the market data from lixinger.com.
- [stock] db is to setup the level db which to store the market data.
- [restful] root is to setup the file path where located the statics html/js files

## Launch Server
```
bin/server
```

## restful URIs
- portfolio application: http://server/static/portfolio.html
- stock data URI (json): http://server/stock/<stockId>
