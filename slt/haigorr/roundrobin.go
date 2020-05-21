package haigorr

import (
	"github.com/oceanho/haigo/slt"
	"sync"
)

type roundRobin struct {
	cur    int64
	max    int64
	locker sync.Mutex
}

func (rr *roundRobin) Next() int64 {
	rr.locker.Lock()
	defer rr.locker.Unlock()
	cur := rr.cur
	if cur < rr.max {
		rr.cur++
	} else {
		cur = 0
		rr.cur = 0
	}
	return cur
}

func New(max int64) slt.Selector {
	return &roundRobin{
		cur: 0,
		max: max,
	}
}
