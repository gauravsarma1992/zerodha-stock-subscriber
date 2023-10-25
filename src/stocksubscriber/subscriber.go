package stocksubscriber

const (
	SubscriberConfigFile = "./config/subscriber_config.json"
)

type (
	Subscriber struct {
		WorkerMgr       *WorkerMgr
		KiteClient      *KiteClient
		Config          *SubscriberConfig
		CallbackHandler CallbackHandler

		ExitCh chan bool
	}
	SubscriberConfig struct {
		ApiKey      string   `json:"api_key"`
		AccessToken string   `json:"access_token"`
		Stocks      []uint32 `json:"stocks"`
	}
	CallbackHandler interface {
		Process(*Stock) error
	}
)

func NewSubscriber(config *SubscriberConfig, callbackHandler CallbackHandler) (subscriber *Subscriber, err error) {
	subscriber = &Subscriber{
		Config:          config,
		CallbackHandler: callbackHandler,
		ExitCh:          make(chan bool),
	}
	if err = subscriber.Setup(); err != nil {
		return
	}
	return
}

func (subscriber *Subscriber) Setup() (err error) {
	if subscriber.WorkerMgr, err = NewWorkerMgr(); err != nil {
		return
	}
	if subscriber.KiteClient, err = NewKiteClient(
		subscriber.Config.ApiKey,
		subscriber.Config.AccessToken,
		subscriber.Config.Stocks,
		subscriber.WorkerMgr,
	); err != nil {
		return
	}
	return
}

func (subscriber *Subscriber) Run() (err error) {
	go subscriber.KiteClient.Run()
	go subscriber.WorkerMgr.Run()
	<-subscriber.ExitCh
	return
}
