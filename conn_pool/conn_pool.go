package conn_pool

import (
	"context"
	"database/sql/driver"
	"sync"
	"time"
)

/*
* @author: Heng ChenChi
* @date: 2022/12/12 0012 16:40
* @version: 1.0
* @description: Golang标准库(database/sql)下的数据库连接池，并非与标准库完全一致。
**/

type ConnPool struct {
	waitDuration int64
	connector    driver.Connector
	numClosed    uint64

	mu           sync.Mutex
	freeConn     []*conn
	connRequests map[uint64]chan connRequest
	nextRequest  uint64
	numOpen      int
	openerCh     chan struct{}
	closed       bool

	maxIdleCount      int
	maxOpen           int
	maxLifeTime       time.Duration
	maxIdleTime       time.Duration
	cleanerCh         chan struct{}
	waitCount         int64
	maxIdleClosed     int64
	maxIdleTimeClosed int64
	maxLifetimeClosed int64

	stop func()
}

// Stats 返回连接池的状态统计信息
func (pool *ConnPool) Stats() {

}

// 在有申请连接的请求，并且连接未达到限制时；通知
// connectionOpener 创建新的连接
func (pool *ConnPool) maybeOpenNewConnection() {

}

// 在单个goroutine中运行，创建一个新的连接，在此goroutine
// 的外部有相应的控制
func (pool *ConnPool) connectionOpener(ctx context.Context) {
	// 监听控制消息
	for {
		select {
		case <-ctx.Done():
			return
		case <-pool.openerCh:
			pool.openNewConnection(ctx)
		}
	}
}

// openNewConnection 新建立一个连接
func (pool *ConnPool) openNewConnection(ctx context.Context) {
	// 是否需要创建新的连接由 maybeOpenNewConnection 方法控制
	// numOpen 在 maybeOpenNewConnection 中已经加1
	raw, err := pool.connector.Connect(ctx)
	pool.mu.Lock()
	defer pool.mu.Unlock()
	// 若连接池已经处于关闭状态，此时不应该新建连接
	if pool.closed {
		if err == nil {
			raw.Close()
		}
		pool.numOpen--
		return
	}
	// 新建连接出错
	if err != nil {
		pool.numOpen--
		pool.maybeOpenNewConnection()
		return
	}
	// 创建 conn_pool.conn
	conn := &conn{
		pool:       pool,
		createdAt:  time.Now(),
		returnedAt: time.Now(),
		raw:        raw,
	}
	pool.putConnLocked(conn, err)
}

// 将是一个 connRequest 得到满足，或者将连接放入空闲池
func (pool *ConnPool) putConnLocked(conn *conn, err error) bool {
	if pool.closed {
		return false
	}
	// 池中当前连接数已经超过设置的最大限制数
	if pool.maxOpen > 0 && pool.numOpen > pool.maxOpen {
		return false
	}
	// 有正在等待创建连接的goroutine
	if n := len(pool.connRequests); n > 0 {
		var (
			reqCh  chan connRequest
			reqKey uint64
		)
		// 取出一个请求
		for reqKey, reqCh = range pool.connRequests {
			break
		}
		// 将可用连接传递给该请求，并删除这个请求
		delete(pool.connRequests, reqKey)
		if err == nil {
			conn.inUse = true
		}
		reqCh <- connRequest{
			conn: conn,
			err:  err,
		}
		return true
		// 无连接请求，则放入空闲队列中
	} else if err == nil && !pool.closed {
		// 还能放入新连接
		if pool.maxIdleConns() > len(pool.freeConn) {
			pool.freeConn = append(pool.freeConn, conn)
			pool.startCleaner()
			return true
		}
		pool.maxIdleClosed++
	}
	return false
}

// 默认最大空闲连接数
const defaultMaxIdleConns = 2

// 获取连接池当前设置的最大空闲连接数
func (pool *ConnPool) maxIdleConns() int {
	n := pool.maxIdleCount
	switch {
	case n == 0:
		return defaultMaxIdleConns
	case n < 0:
		return 0
	default:
		return n
	}
}

func (pool *ConnPool) putConn(conn *conn, err error, resetSession bool) {

}

func (pool *ConnPool) startCleaner() {

}

// New 创建连接池实例
func New(d driver.Driver, dataSourceName string) (*ConnPool, error) {
	// 这里需要根据驱动具体实现的方法来新建连接
	// 检查驱动是否实现了 driver.DriverContext 接口
	if dCtx, ok := d.(driver.DriverContext); ok {
		connector, err := dCtx.OpenConnector(dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewByConnector(connector), nil
	}
	// 未实现 driver.DriverContext 接口
	return NewByConnector(dsnConnector{dsn: dataSourceName, driver: d}), nil
}

// NewByConnector 使用 driver.Connector 对象创建连接池
func NewByConnector(connector driver.Connector) *ConnPool {
	ctx, cancel := context.WithCancel(context.Background())
	// 创建连接池
	pool := &ConnPool{
		connector: connector,
		stop:      cancel,
	}
	// 将通过一个单独的goroutine来创建一个新连接
	go pool.connectionOpener(ctx)
	return pool
}
