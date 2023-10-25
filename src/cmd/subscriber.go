package main

import (
	"log"

	"github.com/gauravsarma1992/stocksubscriber/stocksubscriber"
)

func main() {
	var (
		subsObj    *stocksubscriber.Subscriber
		subsConfig *stocksubscriber.SubscriberConfig
		err        error
	)
	log.Println("Starting Zerodha Stock Subscriber")
	subsConfig = &stocksubscriber.SubscriberConfig{
		ApiKey:      "temp",
		AccessToken: "temp",
		Stocks:      []uint32{43423},
	}
	if subsObj, err = stocksubscriber.NewSubscriber(subsConfig); err != nil {
		log.Fatal(err)
	}
	if err = subsObj.Run(); err != nil {
		log.Fatal(err)
	}

}
