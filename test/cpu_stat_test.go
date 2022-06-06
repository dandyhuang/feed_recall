package main

import (
	"fmt"
	"github.com/go-kratos/aegis/pkg/window"
	"github.com/shirou/gopsutil/v3/cpu"
	"sync/atomic"
	"time"
)
type BBR struct {
	passStat        window.RollingCounter
	rtStat          window.RollingCounter
	inFlight        int64
	bucketPerSecond int64
	bucketDuration  time.Duration

	// prevDropTime defines previous start drop since initTime
	prevDropTime atomic.Value
	maxPASSCache atomic.Value
	minRtCache   atomic.Value

}

func (l *BBR) timespan(lastTime time.Time) int {
	v := int(time.Since(lastTime) / l.bucketDuration)
	if v > -1 {
		return v
	}
	return 100// l.opts.Bucket
}
func main() {
	l:=&BBR{}
	var prevDropTime atomic.Value
	s, _:=prevDropTime.Load().(time.Duration)
	fmt.Println(s)
	l.bucketDuration=time.Second * 10 / time.Duration(100)
	l.bucketPerSecond = int64(time.Second/l.bucketDuration)

	fmt.Println(l, l.timespan(time.Now()), l.timespan(time.Time{}))
	go func() {
		ticker := time.NewTicker(500)
		defer ticker.Stop()
		for {
			<-ticker.C
			var percents []float64
			var u uint64
			percents, err := cpu.Percent(500*time.Millisecond, false)
			if err == nil {
				u = uint64(percents[0] * 10)
			}
			fmt.Println("stat u:", u, percents)
		}
	}()
	select {

	}
}