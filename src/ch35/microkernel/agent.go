package microkernel

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
)

const (
	// iota 常量生成器
	// 用于生成一组以相似规则初始化的常量
	// 在 const 声明语句中，在第一个生命的常量所在的行
	// iota 将会被置为0，然后再每一个有常量声明的行加一
	Waiting = iota
	Running
)

var WrongStateError = errors.New("can not take the operation in the current state")

type CollectorsError struct {
	CollectorsErrors []error
}

type Event struct {
	Source string
	Content string
}

type EventReceiver interface {
	OnEvent(evt Event)
}

type Collector interface {
	Init(evtReceiver EventReceiver) error
	// 为了 cancel 在不同协程中运行的任务
	Start(agtCtx context.Context) error
	Stop() error
	Destroy() error
}

type Agent struct {
	collectors map[string]Collector
	evtBuf chan Event
	cancel context.CancelFunc
	ctx context.Context
	state int
}


func (ce CollectorsError) Error() string {
	var strs []string
	for _, err := range ce.CollectorsErrors {
		strs = append(strs, err.Error())
	}
	return strings.Join(strs, ";")
}

func (agt *Agent) EventProcessGroutine() {
	var evtSeg [10]Event

	for {
		for i := 0; i < 10; i++ {
			select {
			case evtSeg[i] = <-agt.evtBuf:
			case <-agt.ctx.Done():
				return
			}
		}
		fmt.Println(evtSeg)
	}
}

func NewAgent(sizeEvtBuf int) *Agent {
	agt := Agent {
		collectors : map[string]Collector{},
		evtBuf : make(chan Event, sizeEvtBuf),
		state : Waiting,
	}
	return &agt
}

func (agt *Agent) RegisterCollector(name string, collector Collector) error {
	if agt.state != Waiting {
		return WrongStateError
	}

	agt.collectors[name] = collector
	return collector.Init(agt)
}

func (agt *Agent) startCollectors() error {
	var err error
	var errs CollectorsError
	var mutex sync.Mutex

	for name, collector := range agt.collectors {
		go func(name string, collector Collector, ctx context.Context) {
			defer func() {
				mutex.Unlock()
			}()
			err = collector.Start(ctx)
			mutex.Lock()
			if err != nil {
				errs.CollectorsErrors = append(errs.CollectorsErrors, errors.New(name + ":" + err.Error()))
			}
		}(name, collector, agt.ctx)
	}
	if len(errs.CollectorsErrors) == 0 {
		return nil
	}
	return errs
}

func (agt *Agent) stopCollectors() error {
	var err error
	var errs CollectorsError
	for name, collector := range agt.collectors {
		if err = collector.Stop(); err != nil {
			errs.CollectorsErrors = append(errs.CollectorsErrors, errors.New(name + ":" + err.Error()))
		}
	}
	if len(errs.CollectorsErrors) == 0 {
		return nil
	}
	return errs
}

func (agt *Agent) destroyCollectors() error {
	var err error
	var errs CollectorsError
	for name, collector := range agt.collectors {
		if err = collector.Destroy(); err != nil {
			errs.CollectorsErrors = append(errs.CollectorsErrors, errors.New(name + ":" + err.Error()))
		}
	}
	if len(errs.CollectorsErrors) == 0 {
		return nil
	}
	return errs
}

func (agt *Agent) Start() error {
	if agt.state != Waiting {
		return WrongStateError
	}

	agt.state = Running
	agt.ctx, agt.cancel = context.WithCancel(context.Background())
	go agt.EventProcessGroutine()
	return agt.startCollectors()
}

func (agt *Agent) Stop() error {
	if agt.state != Running {
		return WrongStateError
	}
	agt.state = Waiting
	agt.cancel()
	return agt.stopCollectors()
}

func (agt *Agent) Destroy() error {
	if agt.state != Waiting {
		return WrongStateError
	}
	return agt.destroyCollectors()
}

func (agt *Agent) OnEvent(evt Event) {
	agt.evtBuf <- evt
}
