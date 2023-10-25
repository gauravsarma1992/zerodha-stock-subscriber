package stocksubscriber

const (
	SubscriberConfigFile = "./config/subscriber_config.json"
)

type (
	Subscriber struct {
		WorkerMgr *WorkerMgr
		Config    *SubscriberConfig `json:"config"`
	}
	SubscriberConfig struct {
		ApiToken string   `json:"api_token"`
		Stocks   []string `json:"stocks"`
	}
	Stock struct{}
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
	return
}

func (subscriber *Subscriber) Run() (err error) {
	return
}
