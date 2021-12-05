package main

import (
	"errors"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

//连接池的简单实现
type ConnPoolOpt struct {
	MaxIdle  time.Duration
	Limit    int
	DialFunc func() (net.Conn, error)
}

type Conn struct {
	id         int
	parent     *ConnPool
	conn       net.Conn
	lastAccess atomic.Value
	using      int32
}

func (c *Conn) serve() {
	checkIdle := time.NewTimer(c.parent.maxIdle)

	defer checkIdle.Stop()
	defer c.conn.Close()

	for {
		select {
		case <-c.parent.closeCh:
			return
		case <-checkIdle.C:
			if c.expired() {
				return
			}
			checkIdle.Reset(c.parent.maxIdle)
		}
	}
}

func (c *Conn) expired() bool {
	if lastAccess, ok := c.lastAccess.Load().(time.Time); ok {
		return time.Since(lastAccess) > c.parent.maxIdle
	}

	return false
}

func (c *Conn) isUsing() bool {
	return atomic.LoadInt32(&c.using) == 1
}

func (c *Conn) put() {
	atomic.CompareAndSwapInt32(&c.using, 1, 0)
}

func (c *Conn) get() (net.Conn, error) {
	if !atomic.CompareAndSwapInt32(&c.using, 0, 1) {
		return nil, errors.New("conn using")
	}

	defer c.lastAccess.Store(time.Now())
	return c.conn, nil
}

type ConnPool struct {
	conns []*Conn
	rwmu  sync.RWMutex

	closeCh chan struct{}
	closed  int32

	maxIdle  time.Duration
	limit    int
	dialFunc func() (net.Conn, error)
}

func NewConnPool(opt *ConnPoolOpt) *ConnPool {
	if opt.MaxIdle < time.Second {
		opt.MaxIdle = 10 * time.Second
	}
	if opt.Limit < 5 {
		opt.Limit = 10
	}

	cp := &ConnPool{
		closeCh: make(chan struct{}),
		closed:  0,

		maxIdle:  opt.MaxIdle,
		limit:    opt.Limit,
		dialFunc: opt.DialFunc,
	}

	return cp
}

func (cp *ConnPool) Close() error {
	if atomic.CompareAndSwapInt32(&cp.closed, 0, 1) {
		close(cp.closeCh)
	}
	return nil
}

func (cp *ConnPool) Put(conn net.Conn) {
	cp.rwmu.RLock()
	for _, c := range cp.conns {
		if c.conn == conn {
			c.put()
			cp.rwmu.RUnlock()
			return
		}
	}
	cp.rwmu.RUnlock()

	conn.Close()
}

func (cp *ConnPool) Get() (net.Conn, error) {
	cp.rwmu.RLock()
	connLen := len(cp.conns)
	cp.rwmu.RUnlock()

	if connLen <= 0 {
		cp.rwmu.Lock()
		connLen := len(cp.conns)

		if connLen <= 0 {
			conn, err := cp.dialFunc()
			if err != nil {
				cp.rwmu.Unlock()
				return nil, err
			}

			newConn := &Conn{
				id:     connLen,
				parent: cp,
				conn:   conn,
			}
			newConn.lastAccess.Store(time.Now())
			cp.conns = append(cp.conns, newConn)

			go newConn.serve()

			cp.rwmu.Unlock()
			return newConn.get()
		}

		for _, idleConn := range cp.conns {
			if !idleConn.isUsing() {
				cp.rwmu.Unlock()
				return idleConn.get()
			}
		}
		cp.rwmu.Unlock()
	}

	cp.rwmu.Lock()
	for _, idleConn := range cp.conns {
		if !idleConn.isUsing() {
			cp.rwmu.Unlock()
			return idleConn.get()
		}
	}
	cp.rwmu.Unlock()

	return nil, errors.New("reach max limit")
}
