package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/avast/retry-go"
	"github.com/gofiber/fiber/v2"
)

type Ticker struct {
	Symbol  string `json:"symbol"`
	Ticker  string `json:"ticker"`
	Address string `json:"address"`
	Source  string `json:"source"`
}

type JupiterResponse struct {
	Data map[string]JupiterPrice `json:"data"`
}

type JupiterPrice struct {
	Price float64 `json:"price"`
}

type BinanceResponse []BinancePrice

type BinancePrice struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	LastPrice          string `json:"lastPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstId            int64  `json:"firstId"`
	LastId             int64  `json:"lastId"`
	Count              int    `json:"count"`
}

func main() {
	app := fiber.New()

	app.Get("/prices", func(c *fiber.Ctx) error {
		tickers := []Ticker{
			{"Crypto.LST/USD", "LST", "LSTxxxnJzKDFSLr4dUkPcmCf5VyryEqzPLz5j4bpxFp", "jupiter"},
			{"Crypto.JTO/USD", "JTO", "jtojtomepa8beP8AuQc6eXt5FriJwfFMwQx2v2f9mCL", "jupiter"},
			{"Crypto.RENDER/USD", "RENDER", "rndrizKT3MK1iimdxRdWabcF7Zg7AR5T4nud4EkHBof", "jupiter"},
			{"Crypto.BTC/USD", "BTC", "BTCUSDT", "binance"},
		}

		jupiterTickers, binanceTickers := filterTickers(tickers)

		// Use channels to receive data and errors from goroutines
		jupiterChan := make(chan map[string]JupiterPrice)
		binanceChan := make(chan BinanceResponse)
		errChan := make(chan error)

		// Goroutine for fetching Jupiter prices
		go func() {
			jupiterPrices, err := fetchJupiterPrices(jupiterTickers)
			if err != nil {
				errChan <- err
				return
			}
			jupiterChan <- jupiterPrices
		}()

		// Goroutine for fetching Binance prices
		go func() {
			binancePrices, err := fetchBinancePrices(binanceTickers)
			if err != nil {
				errChan <- err
				return
			}
			binanceChan <- binancePrices
		}()

		// Initialize variables for the results
		var jupiterPrices map[string]JupiterPrice
		var binancePrices BinanceResponse

		// Wait for both goroutines to finish
		for i := 0; i < 2; i++ {
			select {
			case jPrices := <-jupiterChan:
				jupiterPrices = jPrices
			case bPrices := <-binanceChan:
				binancePrices = bPrices
			case err := <-errChan:
				return err
			}
		}

		prices := compilePrices(tickers, jupiterPrices, binancePrices)

		return c.JSON(prices)
	})

	app.Listen(":3000")
}

func filterTickers(tickers []Ticker) (string, string) {
	var jupiterTickers, binanceTickers string

	for _, t := range tickers {
		if t.Source == "jupiter" {
			jupiterTickers += t.Ticker + ", "
		} else if t.Source == "binance" {
			binanceTickers += `"` + t.Address + `",`
		}
	}

	// Trim the trailing comma from the binanceTickers string
	binanceTickers = strings.TrimSuffix(binanceTickers, ",")

	return jupiterTickers, "[" + binanceTickers + "]"
}

func fetchJupiterPrices(tickers string) (map[string]JupiterPrice, error) {
	var jupiterResp JupiterResponse

	err := retry.Do(
		func() error {
			url := fmt.Sprintf("https://price.jup.ag/v4/price?ids=%s", tickers)
			resp, reqErr := http.Get(url)
			if reqErr != nil {
				return reqErr
			}
			defer resp.Body.Close()

			body, readErr := ioutil.ReadAll(resp.Body)
			if readErr != nil {
				return readErr
			}

			return json.Unmarshal(body, &jupiterResp)
		},
		retry.Attempts(3),        // Number of retry attempts
		retry.Delay(time.Second), // Delay between retries
		retry.OnRetry(func(n uint, err error) {
			fmt.Printf("Retry #%d due to error: %v\n", n+1, err)
		}),
	)

	if err != nil {
		return nil, err
	}

	return jupiterResp.Data, nil
}

func fetchBinancePrices(tickers string) (BinanceResponse, error) {
	url := fmt.Sprintf("https://api.binance.com/api/v3/ticker?symbols=%s", tickers)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var binanceResp BinanceResponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &binanceResp)
	if err != nil {
		return nil, err
	}

	return binanceResp, nil
}

func compilePrices(tickers []Ticker, jupiterPrices map[string]JupiterPrice, binancePrices BinanceResponse) []interface{} {
	var prices []interface{}

	for _, t := range tickers {
		if t.Source == "jupiter" {
			price := jupiterPrices[t.Ticker].Price
			prices = append(prices, map[string]interface{}{
				"symbol":  t.Symbol,
				"ticker":  t.Ticker,
				"address": t.Address,
				"price":   price,
				"source":  t.Source,
			})
		} else if t.Source == "binance" {
			for _, bPrice := range binancePrices {
				if bPrice.Symbol == t.Address {

					lastPrice, err := strconv.ParseFloat(bPrice.LastPrice, 64)
					if err != nil {
						lastPrice = 0 // handle or log the error as appropriate
					}

					prices = append(prices, map[string]interface{}{
						"symbol":  t.Symbol,
						"ticker":  t.Ticker,
						"address": t.Address,
						"price":   lastPrice,
						"source":  t.Source,
					})

				}
			}
		}
	}

	return prices
}
