package stocksubscriber

import (
	"log"
	"runtime"
)

const (
	ConcurrencyFactor = 2
)

type (
	StatusT uint8

	WorkerMgr struct {
		Workers         map[uint16]*Worker
		CallbackHandler CallbackHandler
	}
	Worker struct {
		ID uint16

		InpCh  chan *Stock
		ExitCh chan bool

		CallbackHandler CallbackHandler

		Status StatusT
	}
)

func NewWorkerMgr(callbackHandler CallbackHandler) (workerMgr *WorkerMgr, err error) {
	workerMgr = &WorkerMgr{
		Workers:         make(map[uint16]*Worker),
		CallbackHandler: callbackHandler,
	}
	if err = workerMgr.Setup(); err != nil {
		return
	}
	return
}

func (workerMgr *WorkerMgr) Setup() (err error) {
	for idx := 0; idx < runtime.NumCPU()*ConcurrencyFactor; idx++ {
		workerId := uint16(idx)
		worker, _ := NewWorker(workerId, workerMgr.CallbackHandler)
		workerMgr.Workers[workerId] = worker
	}
	return
}

func (workerMgr *WorkerMgr) Process(stock *Stock) (err error) {
	workerMgr.Workers[stock.GetID()%ConcurrencyFactor].InpCh <- stock
	return
}

func (workerMgr *WorkerMgr) Run() (err error) {
	for _, worker := range workerMgr.Workers {
		go worker.Run()
	}
	return
}

func NewWorker(id uint16, callbackHandler CallbackHandler) (worker *Worker, err error) {
	worker = &Worker{
		ID:     id,
		InpCh:  make(chan *Stock, 1024),
		ExitCh: make(chan bool),
	}
	return
}

func (worker *Worker) Run() (err error) {
	for {
		select {
		case stock := <-worker.InpCh:
			log.Println("Stock received in worker", stock)
			if err = worker.CallbackHandler.Process(stock); err != nil {
				log.Println(err)
			}
		case <-worker.ExitCh:
			return
		}
	}
	return
}
