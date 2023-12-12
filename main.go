package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ace3/golang-oracle/domain"
	"github.com/ace3/golang-oracle/service"

	"github.com/go-co-op/gocron"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	fiberrecover "github.com/gofiber/fiber/v2/middleware/recover"
)

type JsonTicker struct {
	Symbol  string `json:"symbol"`
	Ticker  string `json:"ticker"`
	Address string `json:"address"`
	Source  string `json:"source"`
}

func main() {
	tickers := []JsonTicker{
		{"Crypto.LST/USD", "LST", "LSTxxxnJzKDFSLr4dUkPcmCf5VyryEqzPLz5j4bpxFp", "jupiter"},
		{"Crypto.JTO/USD", "JTO", "JTOUSDT", "binance"},
		{"Crypto.RENDER/USD", "RENDER", "rndrizKT3MK1iimdxRdWabcF7Zg7AR5T4nud4EkHBof", "jupiter"},
		{"Crypto.BTC/USD", "BTC", "BTCUSDT", "binance"},
	}

	app := fiber.New()

	var memoryPrices []interface{}

	initialPrice, err := fetchTickers(tickers)
	if err != nil {
		panic(err)
	}
	memoryPrices = initialPrice

	app.Use(fiberrecover.New())

	app.Use(compress.New())

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	app.Get("/prices", func(c *fiber.Ctx) error {

		return c.JSON(memoryPrices)
	})

	s := gocron.NewScheduler(time.UTC)
	s.CronWithSeconds("* * * * * *").Do(func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
			}
		}()
		newTicker, err := fetchTickers(tickers)
		if err != nil {
			fmt.Println("Failed to fetch price")
		} else {
			memoryPrices = newTicker
		}
	})
	s.StartAsync()

	app.Listen(":3000")

}

func fetchTickers(tickers []JsonTicker) (prices []interface{}, err error) {
	jupiterTickers, binanceTickers := filterTickers(tickers)

	// Use channels to receive data and errors from goroutines
	jupiterChan := make(chan map[string]domain.JupiterPrice)
	binanceChan := make(chan domain.BinanceResponse)
	errChan := make(chan error)

	// Goroutine for fetching Jupiter prices
	go func() {
		jupiterPrices, err := service.FetchJupiterPrices(jupiterTickers)
		if err != nil {
			errChan <- err
			return
		}
		jupiterChan <- jupiterPrices
	}()

	// Goroutine for fetching Binance prices
	go func() {
		binancePrices, err := service.FetchBinancePrices(binanceTickers)
		if err != nil {
			errChan <- err
			return
		}
		binanceChan <- binancePrices
	}()

	// Initialize variables for the results
	var jupiterPrices map[string]domain.JupiterPrice
	var binancePrices domain.BinanceResponse

	// Wait for both goroutines to finish
	for i := 0; i < 2; i++ {
		select {
		case jPrices := <-jupiterChan:
			jupiterPrices = jPrices
		case bPrices := <-binanceChan:
			binancePrices = bPrices
		case err := <-errChan:
			return nil, err
		}
	}

	prices = compilePrices(tickers, jupiterPrices, binancePrices)
	return prices, nil
}

func filterTickers(tickers []JsonTicker) (string, string) {
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

func compilePrices(tickers []JsonTicker, jupiterPrices map[string]domain.JupiterPrice, binancePrices domain.BinanceResponse) []interface{} {
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
