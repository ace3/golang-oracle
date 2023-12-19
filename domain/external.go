package domain

type JupiterResponse struct {
	Data map[string]JupiterPrice `json:"data"`
}
type JsonTicker struct {
	Symbol  string `json:"symbol"`
	Ticker  string `json:"ticker"`
	Address string `json:"address"`
	Source  string `json:"source"`
}
type DexScreenerResponse struct {
	SchemaVersion string `json:"schemaVersion"`
	Pairs         []struct {
		ChainID     string   `json:"chainId"`
		DexID       string   `json:"dexId"`
		URL         string   `json:"url"`
		PairAddress string   `json:"pairAddress"`
		Labels      []string `json:"labels"`
		BaseToken   struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Symbol  string `json:"symbol"`
		} `json:"baseToken"`
		QuoteToken struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Symbol  string `json:"symbol"`
		} `json:"quoteToken"`
		PriceNative string `json:"priceNative"`
		PriceUsd    string `json:"priceUsd"`
		Txns        struct {
			M5 struct {
				Buys  int `json:"buys"`
				Sells int `json:"sells"`
			} `json:"m5"`
			H1 struct {
				Buys  int `json:"buys"`
				Sells int `json:"sells"`
			} `json:"h1"`
			H6 struct {
				Buys  int `json:"buys"`
				Sells int `json:"sells"`
			} `json:"h6"`
			H24 struct {
				Buys  int `json:"buys"`
				Sells int `json:"sells"`
			} `json:"h24"`
		} `json:"txns"`
		Volume struct {
			H24 float64 `json:"h24"`
			H6  float64 `json:"h6"`
			H1  int     `json:"h1"`
			M5  int     `json:"m5"`
		} `json:"volume"`
		PriceChange struct {
			M5  int     `json:"m5"`
			H1  int     `json:"h1"`
			H6  float64 `json:"h6"`
			H24 float64 `json:"h24"`
		} `json:"priceChange"`
		Liquidity struct {
			Usd   float64 `json:"usd"`
			Base  int     `json:"base"`
			Quote float64 `json:"quote"`
		} `json:"liquidity"`
		Fdv           int   `json:"fdv"`
		PairCreatedAt int64 `json:"pairCreatedAt"`
	} `json:"pairs"`
	Pair struct {
		ChainID     string   `json:"chainId"`
		DexID       string   `json:"dexId"`
		URL         string   `json:"url"`
		PairAddress string   `json:"pairAddress"`
		Labels      []string `json:"labels"`
		BaseToken   struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Symbol  string `json:"symbol"`
		} `json:"baseToken"`
		QuoteToken struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Symbol  string `json:"symbol"`
		} `json:"quoteToken"`
		PriceNative string `json:"priceNative"`
		PriceUsd    string `json:"priceUsd"`
		Txns        struct {
			M5 struct {
				Buys  int `json:"buys"`
				Sells int `json:"sells"`
			} `json:"m5"`
			H1 struct {
				Buys  int `json:"buys"`
				Sells int `json:"sells"`
			} `json:"h1"`
			H6 struct {
				Buys  int `json:"buys"`
				Sells int `json:"sells"`
			} `json:"h6"`
			H24 struct {
				Buys  int `json:"buys"`
				Sells int `json:"sells"`
			} `json:"h24"`
		} `json:"txns"`
		Volume struct {
			H24 float64 `json:"h24"`
			H6  float64 `json:"h6"`
			H1  int     `json:"h1"`
			M5  int     `json:"m5"`
		} `json:"volume"`
		PriceChange struct {
			M5  int     `json:"m5"`
			H1  int     `json:"h1"`
			H6  float64 `json:"h6"`
			H24 float64 `json:"h24"`
		} `json:"priceChange"`
		Liquidity struct {
			Usd   float64 `json:"usd"`
			Base  int     `json:"base"`
			Quote float64 `json:"quote"`
		} `json:"liquidity"`
		Fdv           int   `json:"fdv"`
		PairCreatedAt int64 `json:"pairCreatedAt"`
	} `json:"pair"`
}

type JupiterPrice struct {
	Price float64 `json:"price"`
}

type DexScreenerPrice struct {
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
