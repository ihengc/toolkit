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

	mu          sync.Mutex
	freeConn    []*conn
	nextRequest uint64
	numOpen     int
	openerCh    chan struct{}
	closed      bool

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

// openNewConnection 新建立一个连接
func (pool *ConnPool) openNewConnection(ctx context.Context) {

}

func (pool *ConnPool) putConn(conn *conn, err error, resetSession bool) {

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

// New 创建连接池实例
func New(dbDriver driver.Driver, dataSourceName string) (*ConnPool, error) {
	// 这里需要根据驱动具体实现的方法来新建连接
	// 检查驱动是否实现了 driver.DriverContext 接口
	if dCtx, ok := dbDriver.(driver.DriverContext); ok {
		connector, err := dCtx.OpenConnector(dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewByConnector(connector), nil
	}
	// 未实现 driver.DriverContext 接口
	return nil, nil
}

func NewByConnector(connector driver.Connector) *ConnPool {
	ctx, cancel := context.WithCancel(context.Background())
	pool := &ConnPool{
		stop: cancel,
	}
	go pool.connectionOpener(ctx)
	return pool
}
