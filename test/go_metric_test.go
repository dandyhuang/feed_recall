package main

import (
	"fmt"
	"github.com/rcrowley/go-metrics"
	"log"
	"os"
	"sync/atomic"
	"time"
)


func Gauge() {
	g := metrics.NewGauge()
	metrics.Register("bar", g)
	g.Update(1)


	go metrics.Log(metrics.DefaultRegistry,
		1 * time.Second,
		log.New(os.Stdout, "metrics: ", log.Lmicroseconds))


	var j int64
	j = 1
	for true {
		time.Sleep(time.Second * 1)
		g.Update(j)
		j++
	}
	g.Value()
}
type StatMetric struct {
	m metrics.StandardCounter
	count int64
}

func (c *StatMetric) Swap() int64 {
	old:=atomic.SwapInt64(&c.count, 0)
	return old
}

func main() {
	var c StatMetric
    go func() {
    	for {
			atomic.AddInt64(&c.count, 1)
			fmt.Println("add count", c.count)
			time.Sleep(time.Second/2)
		}
	}()
	for _ = range time.Tick(time.Second*2) {
		old:=atomic.SwapInt64(&c.count, 0)
		fmt.Println("add old", old)
	}
	//m := metrics.NewMeter()
	//metrics.Register("quux", m)
	//m.Mark(1)
	//go metrics.Log(metrics.DefaultRegistry,
	//	1 * time.Second,
	//	log.New(os.Stdout, "metrics: ", log.Lmicroseconds))
	//
	//
	//var j int64
	//j = 1
	//for true {
	//	time.Sleep(time.Second * 1)
	//	j++
	//	m.Mark(j)
	//}

	//t := metrics.NewTimer()
	//metrics.Register("bang", t)
	//
	//t.Time(func() {
	//	//do some thing
	//})
	//t.Update(47)

	//c := metrics.NewCounter()
	//metrics.Register("foo", c)
	// c.Inc(45)
	//
	//go metrics.Log(metrics.DefaultRegistry,
	//	1*time.Second,
	//	log.New(os.Stdout, "metrics: ", log.Lmicroseconds))
	//
	//var j int64
	//j = 1
	//for true {
	//	time.Sleep(time.Second * 1)
	//	c.Inc(1)
	//	j++
	//}
}
