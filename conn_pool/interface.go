package conn_pool

import (
	"context"
	"database/sql"
	"time"
)

/*
* @author: Heng ChenChi
* @date: 2022/12/13 0013 14:08
* @version: 1.0
* @description:
**/

// ConnPoolInterface 连接池接口
type ConnPoolInterface interface {
	// Ping 验证一个连接是否存活，在必要时重新建立连接
	Ping() error
	PingContext(ctx context.Context) error
	// Stats 报告连接池的统计信息
	Stats() ConnPoolStatsInterface

	// Conn 获取数据库连接
	Conn(ctx context.Context) (ConnInterface, error)

	// Close 关闭连接池
	// 执行 Close 后，等待所有已经执行的查询操作完成
	Close() error
}

// ConnPoolStatsInterface 连接池统计数据接口
type ConnPoolStatsInterface interface {
	GetMaxOpenConnections() int
	GetOpenConnections() int
	GetInUse() int
	GetIdle() int
	GetWaitCount() int64
	GetWaitDuration() time.Duration
	GetMaxIdleClosed() int64
	GetMaxIdleTimeClosed() int64
	GetMaxLifetimeClosed() int64
}

// ConnInterface 数据库连接接口
type ConnInterface interface {
	// BeginTx 开启事务
	BeginTx(ctx context.Context, options *sql.TxOptions) (*sql.Tx, error)
	// ExecContext 执行SQL
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	// Close 关闭连接
	Close() error
}

type OptionInterface interface {
	apply(option OptionInterface)
}
