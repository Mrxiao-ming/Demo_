package ants_pool

import "github.com/panjf2000/ants/v2"

type Pool struct {
	pool *ants.Pool
}

// ants also supports customizing the capacity of pool.
// You can invoke the NewPool method to instantiate a pool with a given capacity, as following:
func NewPool(size int, options ...ants.Option) (*ants.Pool, error) {
	return ants.NewPool(size, options...)
}

// NewPoolWithFunc generates an instance of ants pool with a specific function.
func NewPoolWithFunc(size int, pf func(interface{}), options ...ants.Option) (*ants.PoolWithFunc, error) {
	return ants.NewPoolWithFunc(size, pf, options...)
}

type PoolInterface interface {
	Release()
}

