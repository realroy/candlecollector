package domain

// CandleStick ...
type CandleStick struct {
	OpenTime                 float64 `json:"open_time"`
	Open                     string  `json:"open"`
	High                     string  `json:"high"`
	Low                      string  `json:"low"`
	Close                    string  `json:"close"`
	Volume                   string  `json:"volume"`
	CloseTime                int64   `json:"close_time"`
	QuoteAssetVolume         string  `json:"quote_asset_volume"`
	NumberOfTrades           float64 `json:"number_of_trades"`
	TakerBuyBaseAssetVolume  string  `json:"taker_buy_base_asset_volume"`
	TakerBuyQuoteAssetVolume string  `json:"taker_buy_quote_asset_volume"`
}
