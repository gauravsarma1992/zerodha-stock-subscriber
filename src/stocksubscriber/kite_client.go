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
		TickHandler TickHandler
	}
	TickHandler interface {
		Process(stock *Stock) (err error)
	}
	Stock struct {
		Tick *kitemodels.Tick
	}
)

func NewStock(tick *kitemodels.Tick) (stock *Stock, err error) {
	stock = &Stock{
		Tick: tick,
	}
	return
}

func (stock *Stock) GetID() (stockID uint16) {
	stockID = uint16(stock.Tick.InstrumentToken)
	return
}

func NewKiteClient(apiKey, accessToken string, tickHandler TickHandler) (kc *KiteClient, err error) {
	kc = &KiteClient{
		ApiKey:      apiKey,
		AccessToken: accessToken,
		Ticker:      kiteticker.New(apiKey, accessToken),
		TickHandler: tickHandler,
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
	stock, _ := NewStock(&tick)
	kc.TickHandler.Process(stock)
}

func (kc *KiteClient) Run() (err error) {

	kc.Ticker.OnConnect(kc.OnConnect)
	kc.Ticker.OnTick(kc.OnTick)
	kc.Ticker.OnClose(kc.OnClose)

	kc.Ticker.Serve()

	return
}
