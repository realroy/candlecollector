package binance

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// CandleCollectorStrategy ...
type CandleCollectorStrategy struct {
	baseUrl     string
	persistence Persistable
}

// Persitable ...
type Persistable interface {
	Save(path string, data interface{}) error
}

// NewCandleCollectorStrategy ...
func NewCandleCollectorStrategy(p Persistable) *CandleCollectorStrategy {
	return &CandleCollectorStrategy{
		baseUrl:     "https://api.binance.com/api/v3",
		persistence: p,
	}
}

// Collect ...
func (s *CandleCollectorStrategy) Collect(options map[string]interface{}) error {
	s.getCandleStick("BTCUSDT", "1m")
	s.getCandleStick("BTCUSDT", "1h")
	return nil
}

func (s *CandleCollectorStrategy) getCandleStick(symbol, interval string) {

	serverTime, _ := s.getServerTime()
	now := *serverTime
	stopAt := now.AddDate(0, -1, 0)
	endTime := now

	filePath := fmt.Sprintf("data/%s/%s/%d-%d.json", symbol, interval, now.Year(), int(now.Month()))

	klineHistory := make([]interface{}, 0)

	for endTime.Month() != stopAt.Month() {
		log.Println("now: ", endTime.Format(time.RFC3339), "stopAt: ", stopAt.Format(time.RFC3339))
		log.Println("get klines ...")

		klines, _ := s.GetKlines(symbol, interval, endTime.Unix()*1000)
		oldestKline := klines[0].([]interface{})
		log.Println("oldest", oldestKline[0])

		for i := len(klines) - 1; i >= 0; i-- {
			klineHistory = append(klineHistory, klines[i])
		}

		endTime = s.float64ToTimeUnix(oldestKline[0].(float64))
		s.persistence.Save(filePath, klineHistory)

		time.Sleep(2 * time.Second)
	}

	s.persistence.Save(filePath, klineHistory)
}

func (s *CandleCollectorStrategy) getServerTime() (*time.Time, error) {
	url := fmt.Sprintf("%s/exchangeInfo", s.baseUrl)
	resp, err := http.Get(url)

	log.Println(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var res map[string]interface{}
	err = json.Unmarshal(body, &res)

	serverTime := s.float64ToTimeUnix(res["serverTime"].(float64))

	return &serverTime, nil
}

func (s *CandleCollectorStrategy) float64ToTimeUnix(t float64) time.Time {
	str := strconv.Itoa(int(t))
	sec, _ := strconv.Atoi(str[:10])
	nsec, _ := strconv.Atoi(str[10:])

	return time.Unix(int64(sec), int64(nsec)).UTC()
}

func (s *CandleCollectorStrategy) ping() error {
	resp, err := http.Get("https://api.binance.com/api/v3/ping")

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	log.Println("Ping ...", resp.Status, string(body))

	return nil
}

// GetKlines Kline/candlestick bars for a symbol. Klines are uniquely identified by their open time.
func (s *CandleCollectorStrategy) GetKlines(symbol string, interval string, endTime int64) ([]interface{}, error) {
	query := fmt.Sprintf("?symbol=%s&interval=%s&endTime=%d&limit=1000", symbol, interval, endTime)
	url := fmt.Sprintf("%s/klines%s", s.baseUrl, query)

	fmt.Println(url)

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	m := make([]interface{}, 1000)

	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
