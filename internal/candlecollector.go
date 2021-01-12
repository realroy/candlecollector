package internal

import (
	"log"

	binance "candlecollector/internal/infrastructure/datasource/binance"
)

// CollectorStrategy ...
type CollectorStrategy string

const (
	// BinanceCollector collector strategy
	BinanceCollector CollectorStrategy = "Binance"
	// BitKubCollector collector strategy
	BitKubCollector = "BitKub"
)

// CandleCollectorable ...
type CandleCollectorable interface {
	Collect(options map[string]interface{}) error
}

// Persistable ...
type Persistable interface {
	Save(path string, data interface{}) error
	Load(path string) interface{}
}

// CandleCollector ...
type CandleCollector struct {
	strategy    CandleCollectorable
	options     map[string]interface{}
	persistence Persistable
}

// NewCandleCollector ...
func NewCandleCollector(c CollectorStrategy, p Persistable, options map[string]interface{}) *CandleCollector {
	var s CandleCollectorable

	if c == BinanceCollector {
		s = binance.NewCandleCollectorStrategy(p)
	}

	return &CandleCollector{
		strategy:    s,
		options:     options,
		persistence: p,
	}
}

// Start ...
func (c *CandleCollector) Start() {
	log.Println("Candle Collector is starting 🍾")

	c.strategy.Collect(c.options)
}
