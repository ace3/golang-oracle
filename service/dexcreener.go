package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ace3/golang-oracle/domain"
	"github.com/avast/retry-go"
	"github.com/shopspring/decimal"
)

func FetchDexScreener(tickers []domain.JsonTicker) (prices []domain.DexScreenerPrice, err error) {

	dexprices := []domain.DexScreenerPrice{}
	for _, v := range tickers {
		if v.Source == "dexscreener" {

			dexPrice, err := FetchDexScreenerPrices(v.Address)
			if err == nil {
				price, err := decimal.NewFromString(dexPrice.Pair.PriceUsd)
				if err == nil {
					obj := domain.DexScreenerPrice{
						Price: price.InexactFloat64(),
					}

					dexprices = append(dexprices, obj)
				}
			}
		}
	}
	return dexprices, nil
}

func FetchDexScreenerPrices(ticker string) (*domain.DexScreenerResponse, error) {
	var response *domain.DexScreenerResponse

	err := retry.Do(
		func() error {
			url := fmt.Sprintf("https://api.dexscreener.com/latest/dex/pairs/%s", ticker)
			resp, reqErr := http.Get(url)
			if reqErr != nil {
				return reqErr
			}
			defer resp.Body.Close()

			body, readErr := io.ReadAll(resp.Body)
			if readErr != nil {
				return readErr
			}

			return json.Unmarshal(body, &response)
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

	return response, nil
}
