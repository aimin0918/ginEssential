package utils

import (
	"sync/atomic"
	"time"
)

type GoLimit struct {
	ch    chan int
	count int64
}

func NewGoLimit(max int) *GoLimit {
	if max > 100 {
		return &GoLimit{ch: make(chan int, 100)}
	}
	return &GoLimit{ch: make(chan int, max)}
}

func (g *GoLimit) Add() {
	g.ch <- 1
	atomic.AddInt64(&g.count, 1)
}

func (g *GoLimit) Done() {
	atomic.AddInt64(&g.count, -1)
	<-g.ch
}

func (g *GoLimit) Exit() {
	for {
		count := atomic.LoadInt64(&g.count)
		if count == 0 {
			return
		}

		time.Sleep(time.Second)
	}
}
