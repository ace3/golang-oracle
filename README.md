# ace3 Pyth Pricing Engine

## Installation
1. Git clone the project
2. Install `golang` if not already installed, i'm using the `v1.21.4`
3. Run `go mod tidy` to install the dependencies
4. Run `go run main.go` to try it on your local machine
5. Run `go build main.go` to build the binary based on your architecture

## Feature
1. Fetch price from **Orca** & **Binance** 
2. Cache the price for **3** seconds


## Adding new tickers
You can add the ticker by adding the slice of tickers on `tickers` variables in `main.go` file

## TO-DO
- [x] Support for Binance
- [ ] Refactor the code to make it more easier to read
- [ ] Extract the tickers into the database or file