package metrics

import (
	"fmt"
	"sync/atomic"

	"github.com/alwayGo/metrics"
)

var _ metrics.Metric = &counter{}

//CounterOpts ..
type CounterOpts metrics.Opts
type counter struct {
	val int64
}

//NewCounter ..
func NewCounter(opt CounterOpts) metrics.Metric {
	return &counter{}
}

func (c *counter) Add(val int64) {
	if val < 0 {
		panic(fmt.Errorf("stat/metric can not decrease in negative val :%d!", val))
	}
	atomic.AddInt64(&c.val, val)
}

func (c *counter) Value() int64 {
	return atomic.LoadInt64(&c.val)
}
