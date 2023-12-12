package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ace3/golang-oracle/domain"
	"github.com/avast/retry-go"
)

func FetchJupiterPrices(tickers string) (map[string]domain.JupiterPrice, error) {
	var jupiterResp domain.JupiterResponse

	err := retry.Do(
		func() error {
			url := fmt.Sprintf("https://price.jup.ag/v4/price?ids=%s", tickers)
			resp, reqErr := http.Get(url)
			if reqErr != nil {
				return reqErr
			}
			defer resp.Body.Close()

			body, readErr := io.ReadAll(resp.Body)
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
