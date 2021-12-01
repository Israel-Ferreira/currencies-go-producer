package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Israel-Ferreira/currencies-go-producer/config"
	"github.com/Israel-Ferreira/currencies-go-producer/domain"
	"github.com/Israel-Ferreira/currencies-go-producer/producers"
)

const BINANCE_URL = "https://api.binance.com/api/v3/ticker/24hr"
const DEFAULT_FIAT_CURRENCY = "BRL"

func GetBinanceBtcCurrency() {

	pair := fmt.Sprintf("BTC%s", DEFAULT_FIAT_CURRENCY)
	ticker, err := getCurrency(pair)

	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(ticker)

}

func GetBinanceEthCurrency() {
	pair := fmt.Sprintf("ETH%s", DEFAULT_FIAT_CURRENCY)

	body, err := getCurrency(pair)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(body)
}

func getCurrency(pair string) (domain.Ticker, error) {
	binanceUrl := fmt.Sprintf("%s?symbol=%s", BINANCE_URL, pair)
	fmt.Printf("Obtendo a ultima cotação: %s \n", pair)

	producer := producers.CurrenciesProducer{Producer: config.Producer}

	var ticker domain.Ticker

	resp, err := http.Get(binanceUrl)

	if err != nil {
		return domain.Ticker{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return domain.Ticker{}, err
	}

	if err = json.Unmarshal(body, &ticker); err != nil {
		return domain.Ticker{}, err
	}

	fmt.Printf("%v \n", ticker)

	if err = producer.SendKafkaTicker("BINANCE_BTC_TOPIC", ticker); err != nil {
		return ticker, err
	}

	return ticker, nil

}
