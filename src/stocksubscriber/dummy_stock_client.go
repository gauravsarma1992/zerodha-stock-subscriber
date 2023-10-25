package stocksubscriber

import "log"

type (
	DummyStockClient struct{}
)

func NewDummyClient() (ds *DummyStockClient, err error) {
	log.Println("Setting up Dummy Client")
	return
}

func (ds *DummyStockClient) Run() (err error) {
	return
}
