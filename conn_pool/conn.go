package conn_pool

import (
	"context"
	"database/sql/driver"
	"errors"
	"sync"
	"time"
)

/*
* @author: Heng ChenChi
* @date: 2022/12/12 0012 17:02
* @version: 1.0
* @description:
**/

type conn struct {
	// connPool 连接池
	connPool *ConnPool
	// createdAt 当前连接被创建的时间点
	createdAt time.Time

	mu sync.Mutex
	// raw 真实的数据库连接
	raw driver.Conn
	// needReset 是否重置当前连接的会话
	needReset bool
	// closed 标记 raw 是否已经被关闭
	closed bool

	inUse bool
	// returnedAt 连接被创建或返回时的时间点
	returnedAt time.Time
}

// releaseConn 释放当前连接
// 将当前连接放回连接池
func (c *conn) releaseConn(err error) {
	c.connPool.putConn(c, err, true)
}

func (c *conn) expired(timeout time.Duration) bool {
	if timeout < 0 {
		return false
	}
	return c.createdAt.Add(timeout).Before(time.Now())
}

func (c *conn) resetSession(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.needReset {
		return nil
	}
	if sr, ok := c.raw.(driver.SessionResetter); ok {
		return sr.ResetSession(ctx)
	}
	return nil
}

// validateConn 检查当前连接是否可用，能否继续使用
func (c *conn) validateConn(needReset bool) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if needReset {
		c.needReset = true
	}
	if v, ok := c.raw.(driver.Validator); ok {
		return v.IsValid()
	}
	return true
}

// Close 关闭此连接
func (c *conn) Close() error {
	c.mu.Lock()
	if c.closed {
		c.mu.Unlock()
		return errors.New("conn_pool: duplicate conn close")
	}
	c.closed = true
	c.mu.Unlock()
	return nil
}
