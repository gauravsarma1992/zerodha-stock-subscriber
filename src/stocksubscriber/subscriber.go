package stocksubscriber

const (
	SubscriberConfigFile = "./config/subscriber_config.json"
)

type (
	Subscriber struct {
		WorkerMgr *WorkerMgr
		Config    *SubscriberConfig

		StockPollerClient StockPoller
		CallbackHandler   CallbackHandler

		ExitCh chan bool
	}
	SubscriberConfig struct {
		ApiKey      string   `json:"api_key"`
		AccessToken string   `json:"access_token"`
		Stocks      []uint32 `json:"stocks"`
		IsLocal     bool     `json:"is_local"`
	}
	CallbackHandler interface {
		Process(*Stock) error
	}
	StockPoller interface {
		Run() error
	}
)

func NewStockPollerClient(
	isLocal bool,
	apiKey, accessToken string,
	stocks []uint32,
	callback CallbackHandler,
) (poller StockPoller, err error) {

	if isLocal == true {
		poller, _ = NewDummyClient()
		return
	}
	poller, err = NewKiteClient(
		apiKey,
		accessToken,
		stocks,
		callback,
	)
	return
}

func NewSubscriber(config *SubscriberConfig, callbackHandler CallbackHandler) (subscriber *Subscriber, err error) {
	subscriber = &Subscriber{
		Config:          config,
		CallbackHandler: callbackHandler,

		ExitCh: make(chan bool),
	}
	if err = subscriber.Setup(); err != nil {
		return
	}
	return
}

func (subscriber *Subscriber) Setup() (err error) {
	if subscriber.WorkerMgr, err = NewWorkerMgr(subscriber.CallbackHandler); err != nil {
		return
	}
	if subscriber.StockPollerClient, err = NewStockPollerClient(
		subscriber.Config.IsLocal,
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
	go subscriber.StockPollerClient.Run()
	go subscriber.WorkerMgr.Run()
	<-subscriber.ExitCh
	return
}
