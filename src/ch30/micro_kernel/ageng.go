package microkernel

import (
	"bytes"
	"context"
	"errors"
	"sync"
)

const (
	Waiting = iota
	b
	c
)

type Event struct{}

type EventReceiver interface {
	OnEvent(event Event)
}

type Collector interface {
	Init(eventReceiver EventReceiver) error
	Start(agtCxt context.Context) error
	Srop() error
	Destory() error
}

type CollectorsError struct {
	CollectorErrors []error
}

func (coerr CollectorsError) Error() string {
	var buf bytes.Buffer
	for _, v := range coerr.CollectorErrors {
		buf.WriteString(v.Error() + ",")
	}
	return buf.String()
}

type Agent struct {
	collectors map[string]Collector
	eventBuf   chan Event
	cancel     context.CancelFunc
	ctx        context.Context
	state      int
}

func (agt *Agent) OnEvent(event Event) {
	agt.eventBuf <- event
}

func (agt *Agent) RegisterCollector(name string, collector Collector) error {
	if agt.state != Waiting {
		return errors.New("error state")
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
				errs.CollectorErrors = append(errs.CollectorErrors, errors.New(name+":"+err.Error()))
			}
		}(name, collector, agt.ctx)
	}
	return errs
}

func(agt *Agent) stopCollectors() error{
	var err error
	var errs CollectorsError
	for name,collector:=range agt.collectors{
		if err = collector.Srop(); err != nil {
			errs.CollectorErrors = append(errs.CollectorErrors, errors.New(name+":"+err.Error()))
		}
	}
	return errs
}
