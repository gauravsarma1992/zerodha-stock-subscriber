package stocksubscriber

const (
	SubscriberConfigFile = "./config/subscriber_config.json"
)

type (
	Subscriber struct {
		WorkerMgr  *WorkerMgr
		KiteClient *KiteClient
		Config     *SubscriberConfig `json:"config"`
	}
	SubscriberConfig struct {
		ApiKey      string   `json:"api_key"`
		AccessToken string   `json:"access_token"`
		Stocks      []string `json:"stocks"`
	}
)

func NewSubscriber(config *SubscriberConfig) (subscriber *Subscriber, err error) {
	subscriber = &Subscriber{
		Config: config,
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
		subscriber.WorkerMgr,
	); err != nil {
		return
	}
	return
}

func (subscriber *Subscriber) Run() (err error) {
	go subscriber.KiteClient.Run()
	go subscriber.WorkerMgr.Run()
	return
}
