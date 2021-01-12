build:
	go build -o candlecollector ./cmd/candlecollector/main.go

dev:
	go run ./cmd/candlecollector/main.go