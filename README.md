# Zerodha Stock Subscriber

## Examples

```go

type (
	LoggingCallbackHandler struct {
		StartTime time.Time
		Count     uint32
	}
)

func (handler *LoggingCallbackHandler) Process(stock *stocksubscriber.Stock) (err error) {

	handler.Count += 1
	if handler.Count < 100000 {
		return
	}

	log.Println("Total time taken to execute", time.Now().Sub(handler.StartTime).Seconds())

	return
}

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

	if subsObj, err = stocksubscriber.NewSubscriber(
		subsConfig,
		&LoggingCallbackHandler{
			StartTime: time.Now(),
		},
	); err != nil {
		log.Fatal(err)
	}

	if err = subsObj.Run(); err != nil {
		log.Fatal(err)
	}

}
```
