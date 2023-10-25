package stocksubscriber

import (
	"log"

	kitemodels "github.com/zerodha/gokiteconnect/v4/models"
)

type (
	DummyStockClient struct {
		callback CallbackHandler
	}
)

func NewDummyClient(callback CallbackHandler) (ds *DummyStockClient, err error) {
	log.Println("Setting up Dummy Client")
	ds = &DummyStockClient{
		callback: callback,
	}
	return
}

func (ds *DummyStockClient) Run() (err error) {
	for idx := 0; idx < 1000000; idx++ {
		ds.callback.Process(&Stock{
			Tick: &kitemodels.Tick{
				InstrumentToken: uint32(idx),
			},
		})
	}
	return
}
