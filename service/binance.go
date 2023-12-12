package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ace3/golang-oracle/domain"
)

func FetchBinancePrices(tickers string) (domain.BinanceResponse, error) {
	url := fmt.Sprintf("https://api.binance.com/api/v3/ticker?symbols=%s", tickers)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var binanceResp domain.BinanceResponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &binanceResp)
	if err != nil {
		return nil, err
	}

	return binanceResp, nil
}
