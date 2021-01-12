package main

import (
	"candlecollector/internal"
	"candlecollector/internal/infrastructure/persistence"
)

func main() {
	options := map[string]interface{}{
		"ApiKey":    "",
		"SecretKey": "",
	}

	p := persistence.NewDiskPersistence()

	candleCollector := internal.NewCandleCollector(internal.BinanceCollector, p, options)
	candleCollector.Start()
}
