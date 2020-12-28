package gauge

import (
	"sync/atomic"

	"github.com/alwayGo/metrics"
)

var _ metrics.Metric = &gauge{}

//GaugeOpts  ...
type GaugeOpts metrics.Opts

type gauge struct {
	val int64
}

//NewGauge ..
func NewGauge(opt GaugeOpts) metrics.Metric {
	return &gauge{}
}

func (g *gauge) Add(val int64) {
	atomic.AddInt64(&g.val, val)
}

func (g *gauge) Set(val int64) {
	old := atomic.LoadInt64(&g.val)
	atomic.CompareAndSwapInt64(&g.val, old, val)
}

func (g *gauge) Value() int64 {
	return atomic.LoadInt64(&g.val)
}

//GaugeVecOpts ...
type GaugeVecOpts metrics.Metric
