package stocksubscriber

import (
	"log"

	kitemodels "github.com/zerodha/gokiteconnect/v4/models"
	kiteticker "github.com/zerodha/gokiteconnect/v4/ticker"
)

type (
	KiteClient struct {
		ApiKey      string
		AccessToken string
		Ticker      *kiteticker.Ticker
	}
)

func NewKiteClient(apiKey, accessToken string) (kc *KiteClient, err error) {
	kc = &KiteClient{
		ApiKey:      apiKey,
		AccessToken: accessToken,
		Ticker:      kiteticker.New(apiKey, accessToken),
	}
	return
}

func (kc *KiteClient) OnConnect() {
	log.Println("KiteClient Connect")
	return
}

func (kc *KiteClient) OnClose(code int, reason string) {
	log.Println("KiteClient Closed", code, reason)
	return
}

func (kc *KiteClient) OnTick(tick kitemodels.Tick) {
	log.Println("Tick received ->", tick)
}

func (kc *KiteClient) Run() (err error) {

	kc.Ticker.OnConnect(kc.OnConnect)
	kc.Ticker.OnTick(kc.OnTick)
	kc.Ticker.OnClose(kc.OnClose)

	kc.Ticker.Serve()

	return
}
