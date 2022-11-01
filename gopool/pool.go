package gopool

import (
	"sync"

	"github.com/panjf2000/ants/v2"
)

type GoPool struct {
	MaxPoolSize int64
	Wg          *sync.WaitGroup
	options     []ants.Option
	pool        *ants.PoolWithFunc
}

func NewPool() PoolInterface {
	return &GoPool{}
}

type PoolInterface interface {
}

func (p *GoPool) AddGoroutineTask() {
	p.Wg.Add(1)
}

var _ PoolInterface = (*GoPool)(nil)
