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
	Pair          struct {
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
		PriceNative   string `json:"priceNative"`
		PriceUsd      string `json:"priceUsd"`
		Fdv           int    `json:"fdv"`
		PairCreatedAt int64  `json:"pairCreatedAt"`
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
